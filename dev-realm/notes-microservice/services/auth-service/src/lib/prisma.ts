import 'dotenv/config';
import { PrismaPg } from '@prisma/adapter-pg';
import { PrismaClient } from '../../generated/prisma/client';

const connectionString = `${process.env.DATABASE_URL}`;

const adapter = new PrismaPg({ connectionString });
const prisma = new PrismaClient({
    adapter,
    log:
        process.env.SERVER_ENV === 'development'
            ? ['query', 'info', 'warn', 'error']
            : ['error'],
});

// handle graceful shutdown
process.on('beforeExit', async () => {
    console.log('> Closing database connection.');
    await prisma.$disconnect();
});

process.on('SIGINT', async () => {
    console.log('> Closing database connection.');
    await prisma.$disconnect();
    process.exit(0);
});

process.on('SIGTERM', async () => {
    console.log('> Closing database connection.');
    await prisma.$disconnect();
    process.exit(0);
});

async function fetchTables() {
    const res =
        await prisma.$queryRaw`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'`;
    console.table(res);
}

export { prisma, fetchTables };
