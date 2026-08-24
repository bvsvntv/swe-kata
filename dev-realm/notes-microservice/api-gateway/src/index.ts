import app from './app';
import { env } from './config/env.config';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;

const server = app.listen(PORT, () => {
    console.log(`API Gateway listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${ENVIRONMENT}`);
});

async function gracefulShutdown(signal: string) {
    console.log(
        `\n> ${signal} received. Shutting down auth-service gracefully...`,
    );

    server.close(async () => {
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
