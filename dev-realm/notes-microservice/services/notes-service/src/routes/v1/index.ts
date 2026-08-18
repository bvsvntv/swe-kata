import { Router, Request, Response } from 'express';
import noteRouter from './note.route';

const router = Router();

router.get('/heartbeat', (_req: Request, res: Response) => {
    const hrtime = process.hrtime.bigint();
    res.status(200).json({ heartbeat: hrtime.toString() });
});

router.use('/notes', noteRouter);

export default router;
