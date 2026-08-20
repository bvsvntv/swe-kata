import { jwtUtils } from '@/lib/jwt';
import { createAuthMiddleware } from '@shared/src/middlewares/auth.middleware';

const authMiddleware = createAuthMiddleware(jwtUtils.verifyAccessToken);

export { authMiddleware };
