import express from 'express';
import v1Routes from './routes/v1';

const app = express();

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use('/api/v1', v1Routes);

export default app;
