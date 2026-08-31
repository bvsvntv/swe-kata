import express, { Request, Response } from 'express';
import cors from 'cors';
import helmet from 'helmet';
import client, { Registry } from '@prometheus-io/client';
import { rateLimiter } from './middlewares/rateLimiter.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import gatewayRoutes from './routes';

const app = express();

// Create a Registry to hold all metrics
// The registry is responsible for collecting and exposing metrics
const register = new Registry();

// Add default Node.js metrics (memory, CPU, event loop, etc.)
client.collectDefaultMetrics({
    register, // Use our custom registry
    prefix: 'nodejs_', // Prefix all default metrics
    gcDurationBuckets: [0.001, 0.01, 0.1, 1, 2, 5], // GC duration buckets
});

// Metrics endpoint - Prometheus scrapes this URL
app.get('/metrics', async (_req: Request, res: Response) => {
    res.set('Content-Type', client.register.contentType); // Required content type
    res.end(await client.register.metrics()); // Return all registered metrics
});

app.use(express.json({ limit: '10kb' }));
app.use(express.urlencoded({ extended: true }));

app.use(helmet());
app.use(cors());
app.use(rateLimiter);

app.use(gatewayRoutes);

app.use(unknownRouteHandler);
app.use(errorHandler);

export default app;
