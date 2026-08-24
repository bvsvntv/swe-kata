import app from './app';
import { env } from '@/configs/env.config';
import prisma from '@/lib/prisma';
import { logger } from './lib/logger';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;

const server = app.listen(PORT, () => {
    logger.info(`Notes service listening at http://localhost:${PORT}/api/v1`);
    logger.info(`Environment: ${ENVIRONMENT}`);
    logger.info(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});

async function gracefulShutdown(signal: string) {
    logger.info(
        `\n> ${signal} received. Shutting down notes-service gracefully...`,
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
