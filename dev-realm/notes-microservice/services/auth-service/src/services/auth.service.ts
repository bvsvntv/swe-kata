import { AppError } from '@/types';
import { checkPassword, hashPassword } from '@/utils/password.util';
import {
    createUser,
    findUserByEmail,
    findUserByID,
} from '@/repositories/auth.repository';
import { createSession, deleteSession } from './session.service';

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

async function refreshTokens() {
    console.log('refreshTokens function @ auth service');
}

async function getProfile(id: string) {
    const user = await findUserByID(id);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    return user;
}

export { register, login, logout, refreshTokens, getProfile };
