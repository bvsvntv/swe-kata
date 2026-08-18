import {
    create,
    get,
    patch,
    remove,
    update,
} from '@/controllers/note.controller';
import { asyncHandler } from '@shared/src/utils/asyncHandler.util';
import { Router } from 'express';

const router = Router();

router.get('/', asyncHandler(get));
router.post('/', asyncHandler(create));
router.put('/:id', asyncHandler(update));
router.patch('/:id', asyncHandler(patch));
router.delete('/:id', asyncHandler(remove));

export default router;
