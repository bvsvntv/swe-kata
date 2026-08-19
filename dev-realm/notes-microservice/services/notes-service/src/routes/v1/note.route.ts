import {
    create,
    get,
    patch,
    remove,
    update,
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

router.post('/', validate(createNoteSchema, 'body'), asyncHandler(create));
router.get(
    '/:id',
    validate(updateNoteParamsSchema, 'params'),
    asyncHandler(get),
);
router.put(
    '/:id',
    validate(updateNoteParamsSchema, 'params'),
    validate(updateNoteSchema, 'body'),
    asyncHandler(update),
);
router.patch(
    '/:id',

    validate(updateNoteParamsSchema, 'params'),
    validate(patchNoteSchema, 'body'),
    asyncHandler(patch),
);
router.delete(
    '/:id',
    validate(deleteNoteParamsSchema, 'params'),
    asyncHandler(remove),
);

export default router;
