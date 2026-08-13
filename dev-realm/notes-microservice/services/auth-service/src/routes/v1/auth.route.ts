import { Router } from 'express';
import * as authController from '@/controllers/auth.controller';
import { validate } from '@/middlewares/validate.middleware';
import { loginSchema, registerSchema } from '@/schemas/auth.schema';

const router = Router();

router.post(
    '/register',
    validate(registerSchema, 'body'),
    authController.register,
);
router.post('/login', validate(loginSchema, 'body'), authController.login);
router.get('/refresh', authController.refreshTokens);
router.get('/logout', authController.logout);

export default router;
