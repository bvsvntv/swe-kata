import prisma from '@/lib/prisma';
import { AppError } from '@/types';
import { hashPassword } from '@/utils/password.util';

async function register(email: string, password: string) {
    const existingUser = await prisma.user.findUnique({
        where: { email },
    });

    if (existingUser) {
        throw new AppError('Email already taken.', 400);
    }

    const passwordHash = await hashPassword(password);
    const user = await prisma.user.create({
        data: {
            email,
            password: passwordHash,
        },
    });

    return { user };
}

async function login() {
    console.log('login function @ auth service');
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
