import { createJwtUtils } from '@shared/src/utils/jwt.util';
import { env } from '@/config/env.config';

const jwtUtils = createJwtUtils({
    accessTokenSecret: env.ACCESS_TOKEN_SECRET,
});

export { jwtUtils };
