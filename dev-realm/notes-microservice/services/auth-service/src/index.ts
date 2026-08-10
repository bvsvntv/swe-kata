import express, { Request, Response } from 'express';
import dotenv from 'dotenv';

dotenv.config();

const app = express();
const PORT = process.env.SERVER_PORT || 18000;

app.get('/heartbeat', (_req: Request, res: Response) => {
    const hrtime = process.hrtime.bigint();
    res.status(200).json({ heartbeat: hrtime.toString() });
});

app.listen(PORT, () => {
    console.log(`Auth service listening at http://localhost:${PORT}`);
    console.log(`Environment: ${process.env.SERVER_ENV}`);
    console.log(`Heartbeat: http://localhost:${PORT}/heartbeat`);
});
