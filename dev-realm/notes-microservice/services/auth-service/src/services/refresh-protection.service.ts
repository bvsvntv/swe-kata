import crypto from 'crypto';
import { logger } from '@/lib/logger';
import redis from '@/lib/redis';
import { AppError } from '@shared/src/types';

const REFRESH_LOCK_PREFIX = 'refresh_lock';
const REFRESH_LOCK_TTL = 5;

async function acquireRefreshLock(sessionID: string): Promise<string> {
    const lockKey = `${REFRESH_LOCK_PREFIX}:${sessionID}`;
    const token = crypto.randomUUID();

    const result = await redis.set(
        lockKey,
        token,
        'EX',
        REFRESH_LOCK_TTL,
        'NX',
    );

    if (!result) {
        throw new AppError('Refresh request already in progress', 409);
    }

    return token;
}

async function releaseRefreshLock(sessionID: string, token: string) {
    const lockKey = `${REFRESH_LOCK_PREFIX}:${sessionID}`;

    const script = `
        if redis.call('get', KEYS[1]) == ARGV[1] then
            return redis.call('del', KEYS[1])
        else
            return 0
        end
    `;

    await redis.eval(script, 1, lockKey, token);
}

function assertSessionIsValid(session: {
    isRevoked: boolean;
    isDeleted: boolean;
    expiresAt: Date;
}) {
    if (session.isDeleted) {
        throw new AppError('Session no longer exists.', 401);
    }

    if (session.isRevoked) {
        throw new AppError('Session has been revoked.', 401);
    }

    if (session.expiresAt < new Date()) {
        throw new AppError('Session has been expired.', 401);
    }
}

function logRefreshReuse(data: { userID: string; sessionID: string }) {
    logger.warn(`
        event: refresh_token_reuse,
        userID: ${data.userID},
        sessionID: ${data.sessionID},
        timestamp: ${new Date().toISOString()}
        `);
}

function logSuspiciousRefresh(data: {
    userID: string;
    sessionID: string;
    previousUserAgent?: string | null;
    currentUserAgent?: string | null;
}) {
    logger.warn(`
        event: suspicious_refresh,
        userID: ${data.userID},
        sessionID: ${data.sessionID},
        previousUserAgent: ${data.previousUserAgent},
        currentUserAgent: ${data.currentUserAgent},
        timestamp: ${new Date().toISOString()}
        `);
}

export {
    acquireRefreshLock,
    releaseRefreshLock,
    assertSessionIsValid,
    logRefreshReuse,
    logSuspiciousRefresh,
};
