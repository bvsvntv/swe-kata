import app from './app';
import { env } from './config/env.config';
import { logger } from './lib/logger';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;
const SERVICE_NAME = env.SERVICE_NAME;

const server = app.listen(PORT, () => {
    logger.info(`${SERVICE_NAME} listening at http://localhost:${PORT}/api/v1`);
    logger.info(`Environment: ${ENVIRONMENT}`);
});

async function gracefulShutdown(signal: string) {
    logger.info(
        `\n> ${signal} received. Shutting down api-gateway gracefully...`,
    );

    server.close(async () => {
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
