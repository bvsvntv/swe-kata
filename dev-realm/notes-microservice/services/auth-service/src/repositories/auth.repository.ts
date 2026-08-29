import prisma from '@/lib/prisma';
import { User } from 'generated/prisma/client';

async function findUserByEmail(email: string) {
    return await prisma.user.findUnique({
        where: { email },
    });
}

async function findUserByID(id: string) {
    return await prisma.user.findUnique({
        where: { id },
        select: {
            id: true,
            email: true,
            createdAt: true,
            updatedAt: true,
        },
    });
}

async function createUser(email: string, password: string): Promise<User> {
    return await prisma.user.create({
        data: {
            email,
            password,
        },
    });
}

async function revokeUserAllSessions(userID: string) {
    return await prisma.session.updateMany({
        where: {
            userID,
        },
        data: {
            isRevoked: true,
        },
    });
}

export { findUserByEmail, findUserByID, createUser, revokeUserAllSessions };
