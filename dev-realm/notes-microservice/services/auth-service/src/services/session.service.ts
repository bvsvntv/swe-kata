import { signAccessToken, signRefreshToken } from '@/utils/jwt.util';
import { User } from 'generated/prisma/client';
import { endSession, startSession } from '@/repositories/session.repository';
import { env } from '@/config/env.config';
import ms from 'ms';
import { hashValue } from '@/utils/auth.utils';

async function createSession(user: User): Promise<{
    accessToken: string;
    refreshToken: string;
}> {
    const accessToken = signAccessToken({ id: user.id });
    const refreshToken = signRefreshToken({ id: user.id });

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

    await startSession(user.id, tokenHash, expiresAt);

    return {
        accessToken,
        refreshToken,
    };
}

async function deleteSession(userID: string): Promise<void> {
    await endSession(userID);
}

export { createSession, deleteSession };
