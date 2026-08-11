import { Router } from 'express';
import * as authController from '@/controllers/auth.controller';

const router = Router();

router.get('/register', authController.register);
router.get('/login', authController.login);
router.get('/refresh', authController.refreshTokens);
router.get('/logout', authController.logout);

export default router;
