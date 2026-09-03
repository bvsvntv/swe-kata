import { metricRegistry } from '@/lib/metrics';
import prisma from '@/lib/prisma';
import { StartSessionType, UpdateSessionType } from '@/types/auth.types';
import { createDBMetrics } from '@shared/src/utils/dbMetric.util';
import { Session } from 'generated/prisma/client';

export const db = createDBMetrics(metricRegistry, 'auth');

async function startSession(args: StartSessionType): Promise<Session> {
    const { sessionID, userID, token, userAgent, ipAddress, expiresAt } = args;

    return db.measure('session', 'create', () =>
        prisma.session.create({
            data: {
                id: sessionID,
                userID,
                token,
                userAgent,
                ipAddress,
                expiresAt,
            },
        }),
    );
}

async function endSession(userID: string): Promise<void> {
    await db.measure('session', 'deleteMany', () =>
        prisma.session.deleteMany({
            where: { userID: userID },
        }),
    );
}

async function findSessionByID(id: string) {
    return db.measure('session', 'findUnique', () =>
        prisma.session.findUnique({
            where: { id },
            select: {
                id: true,
                token: true,
                expiresAt: true,
                userID: true,
                isRevoked: true,
                isDeleted: true,
                userAgent: true,
            },
        }),
    );
}

async function updateSession(args: UpdateSessionType) {
    const { sessionID, token, expiresAt } = args;

    return db.measure('session', 'update', () =>
        prisma.session.update({
            where: { id: sessionID },
            data: {
                token,
                expiresAt,
            },
        }),
    );
}

export { startSession, endSession, findSessionByID, updateSession };
