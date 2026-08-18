import app from './app';
import { env } from './config/env.config';
import prisma from './lib/prisma';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;

const server = app.listen(PORT, () => {
    console.log(`Notes service listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${ENVIRONMENT}`);
    console.log(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});

async function gracefulShutdown(signal: string) {
    console.log(
        `> ${signal} received. Shutting down notes-service gracefully...`,
    );

    server.close(async () => {
        // Release db connection
        await prisma.$disconnect();
        console.log('> Closing database connection.');
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
