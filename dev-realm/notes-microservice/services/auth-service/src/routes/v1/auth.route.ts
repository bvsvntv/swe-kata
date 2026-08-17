import { Router } from 'express';
import {
    register,
    login,
    getProfile,
    refreshTokens,
    logout,
} from '@/controllers/auth.controller';
import { validate } from '@/middlewares/validate.middleware';
import {
    loginSchema,
    refreshTokenSchema,
    registerSchema,
} from '@/schemas/auth.schema';
import { asyncHandler } from '@/utils/asyncHandler.util';
import { authMiddleware } from '@/middlewares/auth.middleware';

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
