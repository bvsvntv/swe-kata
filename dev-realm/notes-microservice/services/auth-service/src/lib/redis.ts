import { env } from '@/config/env.config';
import Redis from 'ioredis';
import { logger } from './logger';

const redis = new Redis(env.REDIS_URL);

redis.on('connect', () => {
    logger.info('Redis connection established.');
});

redis.on('error', (error) => {
    logger.error(`Error while connecting redis: ${error}`);
});

export default redis;
