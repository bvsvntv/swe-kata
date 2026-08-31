import express from 'express';
import cors from 'cors';
import helmet from 'helmet';
import client from '@prometheus-io/client';
import { rateLimiter } from './middlewares/rateLimiter.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import gatewayRoutes from './routes';

const app = express();

client.collectDefaultMetrics();

app.get('/metrics', async (_req, res) => {
    res.set('Content-Type', client.register.contentType);
    res.end(await client.register.metrics());
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
