import { logger } from '@/lib/logger';
import { createRequestLogger } from '@shared/src/middlewares/requestLogger.middleware';

const requestLogger = createRequestLogger(logger);

export { requestLogger };
