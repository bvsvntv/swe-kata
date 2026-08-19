import { checkPassword, hashPassword } from '@/utils/password.util';
import {
    createUser,
    findUserByEmail,
    findUserByID,
} from '@/repositories/auth.repository';
import { createSession, deleteSession } from './session.service';
import {
    findSessionByID,
    updateSession,
} from '@/repositories/session.repository';
import {
    signAccessToken,
    signRefreshToken,
    verifyRefreshToken,
} from '@/utils/jwt.util';
import ms from 'ms';
import { hashValue } from '@/utils/auth.utils';
import { AppError } from '@shared/src/types';
import { env } from '@/config/env.config';

async function register(email: string, password: string) {
    const existingUser = await findUserByEmail(email);
    if (existingUser) {
        throw new AppError('Email already taken.', 400);
    }

    const passwordHash = await hashPassword(password);
    const user = await createUser(email, passwordHash);

    const { accessToken, refreshToken } = await createSession(user);

    return {
        accessToken,
        refreshToken,
    };
}

async function login(email: string, password: string) {
    const user = await findUserByEmail(email);
    if (!user) {
        throw new AppError('Invalid credentials.', 401);
    }

    const isCorrectPassword = await checkPassword(password, user.password);
    if (!isCorrectPassword) {
        throw new AppError('Invalid credentials.', 401);
    }

    const { accessToken, refreshToken } = await createSession(user);

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

async function refreshTokens(refreshToken: string) {
    const payload = verifyRefreshToken(refreshToken);
    const { sessionID, sub: userID } = payload;

    const session = await findSessionByID(sessionID);

    if (!session) {
        throw new AppError('Session not found.', 404);
    }
    if (session.expiresAt < new Date()) {
        throw new AppError('Refresh refreshToken expired.', 401);
    }

    const incomingRefreshToken = hashValue(refreshToken);
    const isIncomingRefreshTokenValid = incomingRefreshToken === session.token;
    if (!isIncomingRefreshTokenValid) {
        throw new AppError('Invalid refresh refreshToken.', 401);
    }

    const newAccessToken = signAccessToken({ sub: userID, sessionID });
    const newRefreshToken = signRefreshToken({ sub: userID, sessionID });

    const refreshTokenExpiresIn = ms(
        env.REFRESH_TOKEN_EXPIRES_IN as ms.StringValue,
    );
    if (typeof refreshTokenExpiresIn !== 'number') {
        throw new Error(
            'Invalid configuration for refresh refreshToken expiration time.',
        );
    }
    const expiresAt = new Date(Date.now() + refreshTokenExpiresIn);

    const hashedToken = hashValue(refreshToken);
    await updateSession({ sessionID, token: hashedToken, expiresAt });

    return {
        accessToken: newAccessToken,
        refreshToken: newRefreshToken,
    };
}

async function getProfile(id: string) {
    const user = await findUserByID(id);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    return user;
}

export { register, login, logout, refreshTokens, getProfile };
