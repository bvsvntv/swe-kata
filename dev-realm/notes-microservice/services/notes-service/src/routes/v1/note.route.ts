import {
    createNoteController,
    getNoteController,
    updateNoteController,
    patchNoteController,
    deleteNoteController,
} from '@/controllers/note.controller';
import { authMiddleware } from '@/middlewares/auth.middleware';
import {
    createNoteSchema,
    deleteNoteParamsSchema,
    getNoteParamsSchema,
    patchNoteSchema,
    updateNoteParamsSchema,
    updateNoteSchema,
} from '@/schemas/note.schema';
import { validate } from '@shared/src/middlewares/validate.middleware';
import { asyncHandler } from '@shared/src/utils/asyncHandler.util';
import { Router } from 'express';

const router = Router();

router
    .route('/')
    .post(
        authMiddleware,
        validate(createNoteSchema, 'body'),
        asyncHandler(createNoteController),
    );
router
    .route('/:id')
    .get(
        authMiddleware,
        validate(getNoteParamsSchema, 'params'),
        asyncHandler(getNoteController),
    );
router
    .route('/:id')
    .put(
        authMiddleware,
        validate(updateNoteParamsSchema, 'params'),
        validate(updateNoteSchema, 'body'),
        asyncHandler(updateNoteController),
    );
router
    .route('/:id')
    .patch(
        authMiddleware,
        validate(updateNoteParamsSchema, 'params'),
        validate(patchNoteSchema, 'body'),
        asyncHandler(patchNoteController),
    );
router
    .route('/:id')
    .delete(
        authMiddleware,
        validate(deleteNoteParamsSchema, 'params'),
        asyncHandler(deleteNoteController),
    );

export default router;
