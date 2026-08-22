import express from 'express';
import cookieParser from 'cookie-parser';
import helmet from 'helmet';
import v1Routes from './routes/v1';
import { errorHandler } from '@shared/src/middlewares/error.middleware';
import { unknownRouteHandler } from '@shared/src/middlewares/unknownRoute.middleware';
import { requestLogger } from './middlewares/requestLogger.middleware';

const app = express();

// Middlewares
app.use(helmet());
app.use(express.json({ limit: '10kb' }));
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());

app.use(requestLogger);

app.use('/api/v1', v1Routes);

// Unknown route handler
app.use(unknownRouteHandler);

// Global error handler
app.use(errorHandler);

export default app;
