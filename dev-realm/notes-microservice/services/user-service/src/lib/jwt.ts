import { env } from '@/config/env.config';
import { createJwtUtils } from '@shared/src/utils/jwt.util';

const jwtUtils = createJwtUtils({
    accessTokenSecret: env.ACCESS_TOKEN_SECRET,
});

export { jwtUtils };
