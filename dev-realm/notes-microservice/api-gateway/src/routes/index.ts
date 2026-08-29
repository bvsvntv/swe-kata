import { Request, Response, Router } from 'express';
import { setupProxy } from './proxy';

const router = Router();
router.get('/heartbeat', (_req: Request, res: Response) => {
    const hrtime = process.hrtime.bigint();
    res.status(200).json({ heartbeat: hrtime.toString() });
});

setupProxy(router);

export default router;
