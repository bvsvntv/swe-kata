import { env } from '@/config/env.config';
import jwt, { SignOptions } from 'jsonwebtoken';

export function signAccessToken(payload: { id: string }) {
    return jwt.sign(payload, env.ACCESS_TOKEN_SECRET, {
        expiresIn: env.ACCESS_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
    });
}

export function signRefreshToken(payload: { id: string }) {
    return jwt.sign(payload, env.REFRESH_TOKEN_SECRET, {
        expiresIn: env.REFRESH_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
    });
}

export function verifyAccessToken(token: string) {
    return jwt.verify(token, env.ACCESS_TOKEN_SECRET);
}

export function verifyRefreshToken(token: string) {
    return jwt.verify(token, env.REFRESH_TOKEN_SECRET);
}
