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
import { logger } from '@/lib/logger';

async function register(args: RegisterUserType) {
    const { email, password, userAgent, ipAddress } = args;

    const existingUser = await findUserByEmail(email);
    if (existingUser) {
        logger.warn('Domain: user registration failed - email taken', {
            event: 'register_failed',
            reason: 'email_taken',
            email,
            ipAddress,
        });
        throw new AppError('Email already taken.', 400);
    }

    const passwordHash = await hashPassword(password);
    const user = await createUser(email, passwordHash);

    const { accessToken, refreshToken } = await createSession(
        user,
        userAgent,
        ipAddress,
    );

    logger.info('Domain: user registered', {
        event: 'user_registered',
        userId: user.id,
        email: user.email,
    });

    return {
        accessToken,
        refreshToken,
    };
}

async function login(args: LoginUserType) {
    const { email, password, userAgent, ipAddress } = args;

    const user = await findUserByEmail(email);
    if (!user) {
        logger.warn('Domain: login failed - invalid credentials', {
            event: 'login_failed',
            reason: 'invalid_email',
            email,
            ipAddress,
        });
        throw new AppError('Invalid credentials.', 401);
    }

    const isCorrectPassword = await checkPassword(password, user.password);
    if (!isCorrectPassword) {
        logger.warn('Domain: login failed - invalid credentials', {
            event: 'login_failed',
            reason: 'invalid_password',
            email,
            ipAddress,
        });
        throw new AppError('Invalid credentials.', 401);
    }

    const { accessToken, refreshToken } = await createSession(
        user,
        userAgent,
        ipAddress,
    );

    logger.info('Domain: user logged in', {
        event: 'user_login_success',
        userId: user.id,
        email: user.email,
    });

    return {
        accessToken,
        refreshToken,
    };
}

async function logout(id: string) {
    const user = await findUserByID(id);
    if (!user) {
        logger.warn('Domain: logout failed - user not found', {
            event: 'logout_failed',
            userId: id,
        });
        throw new AppError('User not found.', 404);
    }

    await deleteSession(id);

    logger.info('Domain: user logged out', {
        event: 'user_logout',
        userId: id,
    });
}

async function refreshSession(refreshToken: string, userAgent: string) {
    const payload = jwtUtils.verifyRefreshToken(refreshToken);
    const { sessionID, sub: userID } = payload;

    const lockToken = await acquireRefreshLock(sessionID);
    try {
        const session = await findSessionByID(sessionID);

        if (!session) {
            logger.warn('Domain: session refresh failed - session not found', {
                event: 'refresh_failed',
                reason: 'session_not_found',
                sessionID,
                userId: userID,
            });
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

        logger.info('Domain: session refreshed', {
            event: 'session_refresh',
            sessionID,
            userId: userID,
        });

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
        logger.warn('Domain: profile fetch failed - user not found', {
            event: 'profile_fetch_failed',
            userId: id,
        });
        throw new AppError('User not found.', 404);
    }

    return user;
}

export { register, login, logout, refreshSession, getProfile };
