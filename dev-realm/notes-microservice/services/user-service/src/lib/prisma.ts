import 'dotenv/config';
import { PrismaPg } from '@prisma/adapter-pg';
import { env } from '@/configs/env.config';

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
    console.log('> Closing database connection.');
    await prisma.$disconnect();
});

export default prisma;
