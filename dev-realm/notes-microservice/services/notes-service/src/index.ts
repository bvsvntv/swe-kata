import app from './app';
import { env } from './config/env.config';
import { logger } from './lib/logger';
import prisma from './lib/prisma';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;
const SERVICE_NAME = env.SERVICE_NAME || 'notes-service';

const server = app.listen(PORT, () => {
    logger.info(`${SERVICE_NAME} listening at http://localhost:${PORT}/api/v1`);
    logger.info(`Environment: ${ENVIRONMENT}`);
    logger.info(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});

async function gracefulShutdown(signal: string) {
    logger.info(
        `\n> ${signal} received. Shutting down ${SERVICE_NAME} gracefully...`,
    );

    server.close(async () => {
        // Release db connection
        await prisma.$disconnect();
        logger.info('> Closing database connection.');
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
