import app from './app';
import { env } from './config/env.config';
import { logger } from './lib/logger';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;

const server = app.listen(PORT, () => {
    logger.info(`API Gateway listening at http://localhost:${PORT}/api/v1`);
    logger.info(`Environment: ${ENVIRONMENT}`);
});

async function gracefulShutdown(signal: string) {
    logger.info(
        `\n> ${signal} received. Shutting down auth-service gracefully...`,
    );

    server.close(async () => {
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
