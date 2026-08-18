import { Router, Request, Response } from 'express';

const router = Router();

router.get('/heartbeat', (_req: Request, res: Response) => {
    const hrtime = process.hrtime.bigint();
    res.status(200).json({ heartbeat: hrtime.toString() });
});

export default router;
