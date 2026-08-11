import express from 'express';
import dotenv from 'dotenv';

import v1Routes from './routes/v1';
import { fetchTables } from './lib/prisma';

dotenv.config();

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const PORT = process.env.SERVER_PORT || 18000;

fetchTables();

app.use('/api/v1', v1Routes);

app.listen(PORT, () => {
    console.log(`Auth service listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${process.env.SERVER_ENV}`);
    console.log(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});
