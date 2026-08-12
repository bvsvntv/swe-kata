import express from 'express';

// Routes
import v1Routes from './routes/v1';
import { errorHandler } from './middlewares/error.middleware';

const app = express();

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(errorHandler);

app.use('/api/v1', v1Routes);

export default app;
