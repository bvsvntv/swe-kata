import {
    createNoteController,
    getNoteController,
    updateNoteController,
    patchNoteController,
    deleteNoteController,
} from '@/controllers/note.controller';
import {
    createNoteSchema,
    deleteNoteParamsSchema,
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
        validate(createNoteSchema, 'body'),
        asyncHandler(createNoteController),
    );
router
    .route('/:id')
    .get(
        validate(updateNoteParamsSchema, 'params'),
        asyncHandler(getNoteController),
    );
router
    .route('/:id')
    .put(
        validate(updateNoteParamsSchema, 'params'),
        validate(updateNoteSchema, 'body'),
        asyncHandler(updateNoteController),
    );
router
    .route('/:id')
    .patch(
        validate(updateNoteParamsSchema, 'params'),
        validate(patchNoteSchema, 'body'),
        asyncHandler(patchNoteController),
    );
router
    .route('/:id')
    .delete(
        validate(deleteNoteParamsSchema, 'params'),
        asyncHandler(deleteNoteController),
    );

export default router;
