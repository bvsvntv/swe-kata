import { Router } from 'express';

const router = Router();

router.route('/').post();
router.route('/:id').get();
router.route('/:id').put();
router.route('/:id').patch();
router.route('/:id').delete();

export default router;
