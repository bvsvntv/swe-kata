import { Router, Request, Response } from 'express';
import userRoutes from './user.route';

const router = Router();

router.use('/users', userRoutes);

router.get('/heartbeat', (_req: Request, res: Response) => {
    const hrtime = process.hrtime.bigint();
    res.status(200).json({ heartbeat: hrtime.toString() });
});

export default router;
