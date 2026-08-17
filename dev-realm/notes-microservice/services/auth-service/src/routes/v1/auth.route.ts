import { Router } from 'express';
import {
    register,
    login,
    refreshTokens,
    logout,
} from '@/controllers/auth.controller';
import { validate } from '@/middlewares/validate.middleware';
import { loginSchema, registerSchema } from '@/schemas/auth.schema';
import { asyncHandler } from '@/utils/asyncHandler.util';

const router = Router();

router.post(
    '/register',
    validate(registerSchema, 'body'),
    asyncHandler(register),
);
router.post('/login', validate(loginSchema, 'body'), asyncHandler(login));
router.get('/refresh', asyncHandler(refreshTokens));
router.get('/logout', asyncHandler(logout));

export default router;
