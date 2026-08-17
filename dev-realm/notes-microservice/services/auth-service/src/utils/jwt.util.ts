import { env } from '@/config/env.config';
import { JWTPayloadType } from '@/types/auth.types';
import jwt, { SignOptions } from 'jsonwebtoken';

function signAccessToken(payload: JWTPayloadType) {
    return jwt.sign(payload, env.ACCESS_TOKEN_SECRET, {
        expiresIn: env.ACCESS_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
    });
}

function signRefreshToken(payload: JWTPayloadType) {
    return jwt.sign(payload, env.REFRESH_TOKEN_SECRET, {
        expiresIn: env.REFRESH_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
    });
}

function verifyAccessToken(token: string) {
    return jwt.verify(token, env.ACCESS_TOKEN_SECRET) as JWTPayloadType;
}

function verifyRefreshToken(token: string) {
    return jwt.verify(token, env.REFRESH_TOKEN_SECRET) as JWTPayloadType;
}

export {
    signAccessToken,
    verifyAccessToken,
    signRefreshToken,
    verifyRefreshToken,
};
