import { checkPassword, hashPassword } from '@/utils/password.util';
import {
    createUser,
    findUserByEmail,
    findUserByID,
    revokeUserAllSessions,
} from '@/repositories/auth.repository';
import { createSession, deleteSession } from './session.service';
import {
    findSessionByID,
    updateSession,
} from '@/repositories/session.repository';
import ms from 'ms';
import { hashValue } from '@/utils/auth.utils';
import { AppError } from '@shared/src/types';
import { env } from '@/config/env.config';
import { jwtUtils } from '@/lib/jwt';
import { LoginUserType, RegisterUserType } from '@/types/auth.types';
import {
    acquireRefreshLock,
    assertSessionIsValid,
    logRefreshReuse,
    logSuspiciousRefresh,
    releaseRefreshLock,
} from './refresh-protection.service';

async function register(args: RegisterUserType) {
    const { email, password, userAgent, ipAddress } = args;

    const existingUser = await findUserByEmail(email);
    if (existingUser) {
        throw new AppError('Email already taken.', 400);
    }

    const passwordHash = await hashPassword(password);
    const user = await createUser(email, passwordHash);

    const { accessToken, refreshToken } = await createSession(
        user,
        userAgent,
        ipAddress,
    );

    return {
        accessToken,
        refreshToken,
    };
}

async function login(args: LoginUserType) {
    const { email, password, userAgent, ipAddress } = args;

    const user = await findUserByEmail(email);
    if (!user) {
        throw new AppError('Invalid credentials.', 401);
    }

    const isCorrectPassword = await checkPassword(password, user.password);
    if (!isCorrectPassword) {
        throw new AppError('Invalid credentials.', 401);
    }

    const { accessToken, refreshToken } = await createSession(
        user,
        userAgent,
        ipAddress,
    );

    return {
        accessToken,
        refreshToken,
    };
}

async function logout(id: string) {
    const user = await findUserByID(id);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    await deleteSession(id);
}

async function refreshSession(refreshToken: string, userAgent: string) {
    const payload = jwtUtils.verifyRefreshToken(refreshToken);
    const { sessionID, sub: userID } = payload;

    const lockToken = await acquireRefreshLock(sessionID);
    try {
        const session = await findSessionByID(sessionID);

        if (!session) {
            throw new AppError('Session not found.', 404);
        }

        // Validate current session
        assertSessionIsValid({
            isRevoked: session.isRevoked,
            isDeleted: session.isDeleted,
            expiresAt: session.expiresAt,
        });

        if (userAgent && session.userAgent && userAgent !== session.userAgent) {
            logSuspiciousRefresh({
                userID: session.userID,
                sessionID: session.id,
                previousUserAgent: session.userAgent,
                currentUserAgent: userAgent,
            });
        }

        const incomingRefreshToken = hashValue(refreshToken);
        const isIncomingRefreshTokenValid =
            incomingRefreshToken === session.token;

        if (!isIncomingRefreshTokenValid) {
            logRefreshReuse({
                userID: session.userID,
                sessionID: session.id,
            });

            await revokeUserAllSessions(userID);

            throw new AppError('Refresh token reuse detected.', 401);
        }

        const newAccessToken = jwtUtils.signAccessToken({
            sub: userID,
            sessionID,
        });
        const newRefreshToken = jwtUtils.signRefreshToken({
            sub: userID,
            sessionID,
        });

        const refreshTokenExpiresIn = ms(
            env.REFRESH_TOKEN_EXPIRES_IN as ms.StringValue,
        );
        if (typeof refreshTokenExpiresIn !== 'number') {
            throw new Error(
                'Invalid configuration for refresh refreshToken expiration time.',
            );
        }
        const expiresAt = new Date(Date.now() + refreshTokenExpiresIn);

        const hashedToken = hashValue(newRefreshToken);
        await updateSession({ sessionID, token: hashedToken, expiresAt });

        return {
            accessToken: newAccessToken,
            refreshToken: newRefreshToken,
        };
    } finally {
        await releaseRefreshLock(sessionID, lockToken);
    }
}

async function getProfile(id: string) {
    const user = await findUserByID(id);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    return user;
}

export { register, login, logout, refreshSession, getProfile };
