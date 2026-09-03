import prisma from '@/lib/prisma';
import { User } from 'generated/prisma/client';
import { db } from './session.repository';

async function findUserByEmail(email: string) {
    return db.measure('user', 'findUnique', () =>
        prisma.user.findUnique({
            where: { email },
        }),
    );
}

async function findUserByID(id: string) {
    return db.measure('user', 'findUnique', () =>
        prisma.user.findUnique({
            where: { id },
            select: {
                id: true,
                email: true,
                createdAt: true,
                updatedAt: true,
            },
        }),
    );
}

async function createUser(email: string, password: string): Promise<User> {
    return db.measure('user', 'create', () =>
        prisma.user.create({
            data: {
                email,
                password,
            },
        }),
    );
}

async function revokeUserAllSessions(userID: string) {
    return db.measure('user', 'updateMany', () =>
        prisma.session.updateMany({
            where: {
                userID,
            },
            data: {
                isRevoked: true,
            },
        }),
    );
}

export { findUserByEmail, findUserByID, createUser, revokeUserAllSessions };
