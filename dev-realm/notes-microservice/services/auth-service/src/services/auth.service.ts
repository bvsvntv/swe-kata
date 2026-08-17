import { AppError } from '@/types';
import { checkPassword, hashPassword } from '@/utils/password.util';
import {
    createUser,
    findUserByEmail,
    findUserByID,
} from '@/repositories/auth.repository';
import { createSession, deleteSession } from './session.service';
import {
    findSessionByToken,
    updateSession,
} from '@/repositories/session.repository';
import { signAccessToken, signRefreshToken } from '@/utils/jwt.util';
import ms from 'ms';
import { env } from 'process';
import { hashValue } from '@/utils/auth.utils';

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

async function refreshTokens(token: string) {
    const incomingRefreshToken = hashValue(token);

    const session = await findSessionByToken(incomingRefreshToken);
    if (!session) {
        throw new AppError('Session not found.', 404);
    }
    if (session.expiresAt < new Date()) {
        throw new AppError('Refresh token expired.', 401);
    }

    const isIncomingRefreshTokenValid = incomingRefreshToken === session.token;
    if (!isIncomingRefreshTokenValid) {
        throw new AppError('Invalid refresh token.', 401);
    }

    const accessToken = signAccessToken({ id: session.userID });
    const refreshToken = signRefreshToken({ id: session.userID });
    const hashedToken = hashValue(refreshToken);

    const refreshTokenExpiresIn = ms(
        env.REFRESH_TOKEN_EXPIRES_IN as ms.StringValue,
    );
    if (typeof refreshTokenExpiresIn !== 'number') {
        throw new Error(
            'Invalid configuration for refresh token expiration time.',
        );
    }
    const expiresAt = new Date(Date.now() + refreshTokenExpiresIn);

    await updateSession(session.id, hashedToken, expiresAt);

    return {
        accessToken,
        refreshToken,
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
