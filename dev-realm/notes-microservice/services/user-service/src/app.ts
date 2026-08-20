import express from 'express';
import v1Routes from './routes/v1';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';

const app = express();

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use('/api/v1', v1Routes);

// Unknown route handler
app.use(unknownRouteHandler);

// Global error handler
app.use(errorHandler);

export default app;
