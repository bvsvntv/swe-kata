import dotenv from 'dotenv';
import app from './app';

dotenv.config();

const PORT = process.env.SERVER_PORT || 18000;

app.listen(PORT, () => {
    console.log(`Auth service listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${process.env.SERVER_ENV}`);
    console.log(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});
