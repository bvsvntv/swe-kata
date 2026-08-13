import app from './app';
import prisma from './lib/prisma';
import { env } from './config/env.config';

const PORT = env.SERVER_PORT;
const ENVIRONMENT = env.SERVER_ENV;

const server = app.listen(PORT, () => {
    console.log(`Auth service listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${ENVIRONMENT}`);
    console.log(`Heartbeat: http://localhost:${PORT}/api/v1/heartbeat`);
});

async function gracefulShutdown(signal: string) {
    console.log(`> ${signal} received. Shutting down gracefully...`);

    server.close(async () => {
        // Release db connection
        await prisma.$disconnect();
        console.log('> Closing database connection.');
        process.exit(0);
    });
}

process.on('SIGINT', () => gracefulShutdown('SIGINT'));

process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
