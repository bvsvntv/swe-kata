import express, { Request, Response } from 'express';
import cors from 'cors';
import helmet from 'helmet';
import client from '@prometheus-io/client';
import { rateLimiter } from './middlewares/rateLimiter.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import gatewayRoutes from './routes';
import { metricRegistry } from './lib/metrics';
import { createMetricsMiddleware } from '@shared/src/middlewares/metrics.middleware';

const app = express();

// Add default Node.js metrics (memory, CPU, event loop, etc.)
client.collectDefaultMetrics({
    register: metricRegistry, // Use our custom registry
    prefix: 'nodejs_',
    gcDurationBuckets: [
        0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
    ],
});

// Metrics endpoint - Prometheus scrapes this URL
app.get('/metrics', async (_req: Request, res: Response) => {
    res.set('Content-Type', metricRegistry.contentType);
    res.end(await metricRegistry.metrics());
});

app.use(express.json({ limit: '10kb' }));
app.use(express.urlencoded({ extended: true }));

app.use(helmet());
app.use(cors());
app.use(rateLimiter);

app.use(createMetricsMiddleware(metricRegistry, 'api-gateway'));
app.use(gatewayRoutes);

app.use(unknownRouteHandler);
app.use(errorHandler);

export default app;
