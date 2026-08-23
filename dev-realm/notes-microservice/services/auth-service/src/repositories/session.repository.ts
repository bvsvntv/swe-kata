import prisma from '@/lib/prisma';
import { StartSessionType, UpdateSessionType } from '@/types/auth.types';
import { Session } from 'generated/prisma/client';

async function startSession(args: StartSessionType): Promise<Session> {
    const { sessionID, userID, token, userAgent, ipAddress, expiresAt } = args;

    return await prisma.session.create({
        data: {
            id: sessionID,
            userID,
            token,
            userAgent,
            ipAddress,
            expiresAt,
        },
    });
}

async function endSession(userID: string): Promise<void> {
    await prisma.session.deleteMany({
        where: { userID: userID },
    });
}

async function findSessionByID(id: string) {
    return await prisma.session.findUnique({
        where: { id },
        select: {
            id: true,
            token: true,
            expiresAt: true,
            userID: true,
        },
    });
}

async function updateSession(args: UpdateSessionType) {
    const { sessionID, token, expiresAt } = args;
    return await prisma.session.update({
        where: { id: sessionID },
        data: {
            token,
            expiresAt,
        },
    });
}

export { startSession, endSession, findSessionByID, updateSession };
