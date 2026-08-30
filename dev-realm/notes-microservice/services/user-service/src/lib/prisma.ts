import 'dotenv/config';
import { PrismaPg } from '@prisma/adapter-pg';
import { PrismaClient } from '../../generated/prisma/client';
import { env } from '@/config/env.config';
import { logger } from './logger';

const connectionString = `${process.env.DATABASE_URL}`;

const adapter = new PrismaPg({ connectionString });
const prisma = new PrismaClient({
    adapter,
    log:
        env.SERVER_ENV === 'development'
            ? ['query', 'info', 'warn', 'error']
            : ['error'],
});

// handle graceful shutdown
process.on('beforeExit', async () => {
    logger.info('> Closing database connection.');
    await prisma.$disconnect();
});

export default prisma;
