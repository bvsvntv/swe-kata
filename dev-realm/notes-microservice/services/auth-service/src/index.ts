import 'dotenv/config';
import app from './app';
import prisma from './lib/prisma';

const PORT = process.env.SERVER_PORT || 18000;

const server = app.listen(PORT, () => {
    console.log(`Auth service listening at http://localhost:${PORT}/api/v1`);
    console.log(`Environment: ${process.env.SERVER_ENV}`);
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
