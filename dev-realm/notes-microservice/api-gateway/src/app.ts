import express from 'express';
import { rateLimiter } from './middlewares/rateLimiter.middleware';
import { setupProxy } from './routes/proxy';

const app = express();

app.use(express.json({ limit: '10kb' }));
app.use(express.urlencoded({ extended: true }));
app.use(rateLimiter);

// service routes
setupProxy(app);

export default app;
