import { Response } from 'express';
import { env } from '@/config/env.config';
import crypto from 'crypto';
import ms from 'ms';

const refreshTokenCookieOptions = {
    httpOnly: true,
    secure: env.SERVER_ENV === 'production',
    sameSite: 'lax' as const,
};

function hashValue(value: string): string {
    return crypto.createHash('sha256').update(value).digest('hex');
}

function generateSessionID(): string {
    return crypto.randomUUID();
}

function setCookies(res: Response, refreshToken: string) {
    const refreshTokenMaxAge = ms(
        env.REFRESH_TOKEN_EXPIRES_IN as ms.StringValue,
    );
    if (typeof refreshTokenMaxAge !== 'number') {
        throw new Error(
            'Invalid configuration for refresh token expiration time.',
        );
    }

    res.cookie('refreshToken', refreshToken, {
        ...refreshTokenCookieOptions,
        maxAge: refreshTokenMaxAge,
    });
}

function clearCookies(res: Response) {
    res.clearCookie('refreshToken', refreshTokenCookieOptions);
}

export { hashValue, generateSessionID, setCookies, clearCookies };
