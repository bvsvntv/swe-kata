import { Router } from 'express';
import {
    register,
    login,
    getProfile,
    refreshTokens,
    logout,
} from '@/controllers/auth.controller';
import {
    loginSchema,
    refreshTokenSchema,
    registerSchema,
} from '@/schemas/auth.schema';
import { authMiddleware } from '@/middlewares/auth.middleware';
import { validate } from '@shared/src/middlewares/validate.middleware';
import { asyncHandler } from '@shared/src/utils/asyncHandler.util';

const router = Router();

router.post(
    '/register',
    validate(registerSchema, 'body'),
    asyncHandler(register),
);
router.post('/login', validate(loginSchema, 'body'), asyncHandler(login));
router.post(
    '/refresh',
    authMiddleware,
    validate(refreshTokenSchema, 'body'),
    asyncHandler(refreshTokens),
);
router.get('/me', authMiddleware, asyncHandler(getProfile));
router.get('/logout', authMiddleware, asyncHandler(logout));

export default router;
