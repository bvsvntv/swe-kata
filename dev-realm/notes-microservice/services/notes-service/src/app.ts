import express, { Request, Response } from 'express';
import client from '@prometheus-io/client';
import v1Routes from './routes/v1';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { requestLogger } from './middlewares/requestLogger.middleware';

const app = express();

client.collectDefaultMetrics();
app.get('/metrics', async (_req: Request, res: Response) => {
    res.set('Content-Type', client.register.contentType);
    res.end(await client.register.metrics());
});

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use(requestLogger);

app.use('/api/v1', v1Routes);

// Unknown route handler
app.use(unknownRouteHandler);

// Global error handler
app.use(errorHandler);

export default app;
