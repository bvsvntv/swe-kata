import { env } from '@/config/env.config';
import { createJwtUtils } from '@shared/src/utils/jwt.util';
import { SignOptions } from 'jsonwebtoken';

const jwtUtils = createJwtUtils({
    accessTokenSecret: env.ACCESS_TOKEN_SECRET,
    refreshTokenSecret: env.REFRESH_TOKEN_SECRET,
    accessTokenExpiresIn:
        env.ACCESS_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
    refreshTokenExpiresIn:
        env.REFRESH_TOKEN_EXPIRES_IN as SignOptions['expiresIn'],
});

export { jwtUtils };
