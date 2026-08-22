import { Router } from 'express';
import {
    registerController,
    loginController,
    getProfileController,
    refreshTokensController,
    logoutController,
} from '@/controllers/auth.controller';
import { loginSchema, registerSchema } from '@/schemas/auth.schema';
import { authMiddleware } from '@/middlewares/auth.middleware';
import { validate } from '@shared/src/middlewares/validate.middleware';
import { asyncHandler } from '@shared/src/utils/asyncHandler.util';

const router = Router();

router
    .route('/register')
    .post(validate(registerSchema, 'body'), asyncHandler(registerController));
router
    .route('/login')
    .post(validate(loginSchema, 'body'), asyncHandler(loginController));
router.route('/refresh').post(asyncHandler(refreshTokensController));
router.route('/me').get(authMiddleware, asyncHandler(getProfileController));
router.route('/logout').get(authMiddleware, asyncHandler(logoutController));

export default router;
