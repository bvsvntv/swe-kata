import jwt, { SignOptions } from 'jsonwebtoken';
import { JWTPayloadType } from '../types/auth.types';

type JwtConfig = {
    accessTokenSecret: string;
    refreshTokenSecret?: string;
    accessTokenExpiresIn?: SignOptions['expiresIn'];
    refreshTokenExpiresIn?: SignOptions['expiresIn'];
};

export function createJwtUtils(config: JwtConfig) {
    function signAccessToken(payload: JWTPayloadType) {
        return jwt.sign(payload, config.accessTokenSecret, {
            expiresIn: config.accessTokenExpiresIn as SignOptions['expiresIn'],
        });
    }

    function signRefreshToken(payload: JWTPayloadType) {
        return jwt.sign(payload, config.refreshTokenSecret as string, {
            expiresIn: config.refreshTokenExpiresIn as SignOptions['expiresIn'],
        });
    }

    function verifyAccessToken(token: string) {
        return jwt.verify(token, config.accessTokenSecret) as JWTPayloadType;
    }

    function verifyRefreshToken(token: string) {
        return jwt.verify(
            token,
            config.refreshTokenSecret as string,
        ) as JWTPayloadType;
    }

    return {
        signAccessToken,
        verifyAccessToken,
        signRefreshToken,
        verifyRefreshToken,
    };
}
