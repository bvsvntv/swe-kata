import { Router } from 'express';
import * as authController from '@/controllers/auth.controller';
import { validate } from '@/middlewares/validate.middleware';
import { registerSchema } from '@/schemas/auth.schema';

const router = Router();

router.post(
    '/register',
    validate(registerSchema, 'body'),
    authController.register,
);
router.get('/login', authController.login);
router.get('/refresh', authController.refreshTokens);
router.get('/logout', authController.logout);

export default router;
