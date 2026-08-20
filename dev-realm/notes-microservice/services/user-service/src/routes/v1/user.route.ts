import {
    createProfileController,
    deleteProfileController,
    getProfileController,
    updateProfileController,
} from '@/controllers/user.controller';
import { authMiddleware } from '@/middlewares/auth.middleware';
import {
    createProfileSchema,
    updateProfileSchema,
} from '@/schema/profile.schema';
import { validate } from '@shared/src/middlewares/validate.middleware';
import { asyncHandler } from '@shared/src/utils/asyncHandler.util';
import { Router } from 'express';

const router = Router();

router
    .route('/profile')
    .post(
        authMiddleware,
        validate(createProfileSchema),
        asyncHandler(createProfileController),
    );
router
    .route('/profile')
    .get(authMiddleware, asyncHandler(getProfileController));
router
    .route('/profile')
    .put(
        authMiddleware,
        validate(updateProfileSchema),
        asyncHandler(updateProfileController),
    );
router
    .route('/profile')
    .delete(authMiddleware, asyncHandler(deleteProfileController));

export default router;
