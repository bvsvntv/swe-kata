import prisma from '@/lib/prisma';
import { Session } from 'generated/prisma/client';

async function startSession(
    userID: string,
    token: string,
    expiresAt: Date,
): Promise<Session> {
    return await prisma.session.create({
        data: {
            userID,
            token,
            expiresAt,
        },
    });
}

async function endSession(userID: string): Promise<void> {
    await prisma.session.deleteMany({
        where: { userID: userID },
    });
}

async function findSessionByToken(token: string) {
    return await prisma.session.findUnique({
        where: { token },
        select: {
            id: true,
            token: true,
            expiresAt: true,
            userID: true,
        },
    });
}

async function updateSession(id: string, token: string, expiresAt: Date) {
    return await prisma.session.update({
        where: { id },
        data: {
            token,
            expiresAt,
        },
    });
}

export { startSession, endSession, findSessionByToken, updateSession };
