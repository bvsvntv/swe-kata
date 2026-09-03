import express, { Request, Response } from 'express';
import client from '@prometheus-io/client';
import v1Routes from './routes/v1';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { requestLogger } from './middlewares/requestLogger.middleware';
import { metricRegistry } from './lib/metrics';
import { createMetricsMiddleware } from '@shared/src/middlewares/metrics.middleware';
import { env } from './config/env.config';

const app = express();

client.collectDefaultMetrics({
    register: metricRegistry,
    prefix: 'nodejs_',
    gcDurationBuckets: [
        0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
    ],
});
app.get('/metrics', async (_req: Request, res: Response) => {
    res.set('Content-Type', metricRegistry.contentType);
    res.end(await metricRegistry.metrics());
});

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use(requestLogger);

app.use(
    createMetricsMiddleware(metricRegistry, env.SERVICE_NAME ?? 'user-service'),
);
app.use('/api/v1', v1Routes);

// Unknown route handler
app.use(unknownRouteHandler);

// Global error handler
app.use(errorHandler);

export default app;
