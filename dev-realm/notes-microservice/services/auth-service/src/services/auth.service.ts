import { AppError } from '@/types';
import { checkPassword, hashPassword } from '@/utils/password.util';
import { createUser, findUserByEmail } from '@/repositories/auth.repository';
import { createSession } from './session.service';

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

async function logout() {
    console.log('logout function @ auth service');
}

async function generateTokens() {
    console.log('generateTokens function @ auth service');
}

async function refreshTokens() {
    console.log('refreshTokens function @ auth service');
}

export { register, login, logout, generateTokens, refreshTokens };
