import express from 'express';
import { rateLimiter } from './middlewares/rateLimiter.middleware';

const app = express();

app.use(express.json({ limit: '10kb' }));
app.use(express.urlencoded({ extended: true }));
app.use(rateLimiter);

export default app;
