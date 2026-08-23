import ms from 'ms';
import { signAccessToken, signRefreshToken } from '@/utils/jwt.util';
import { User } from 'generated/prisma/client';
import { endSession, startSession } from '@/repositories/session.repository';
import { env } from '@/config/env.config';
import { hashValue, generateSessionID } from '@/utils/auth.utils';

async function createSession(user: User, userAgent: string, ipAddress: string) {
    const sessionID = generateSessionID();

    const accessToken = signAccessToken({ sub: user.id, sessionID });
    const refreshToken = signRefreshToken({ sub: user.id, sessionID });

    const tokenHash = hashValue(refreshToken);

    const refreshTokenExpiresIn = ms(
        env.REFRESH_TOKEN_EXPIRES_IN as ms.StringValue,
    );
    if (typeof refreshTokenExpiresIn !== 'number') {
        throw new Error(
            'Invalid configuration for refresh token expiration time.',
        );
    }
    const expiresAt = new Date(Date.now() + refreshTokenExpiresIn);

    await startSession({
        sessionID,
        userID: user.id,
        token: tokenHash,
        userAgent,
        ipAddress,
        expiresAt,
    });

    return {
        accessToken,
        refreshToken,
    };
}

async function deleteSession(userID: string): Promise<void> {
    await endSession(userID);
}

export { createSession, deleteSession };
