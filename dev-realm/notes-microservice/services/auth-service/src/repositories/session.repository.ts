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

export { startSession, endSession };
