import { env } from '@/config/env.config';
import { logger } from '@/lib/logger';
import { ServiceConfigType } from '@/types';
import { Router } from 'express';
import { createProxyMiddleware, Options } from 'http-proxy-middleware';

const serviceConfigs: ServiceConfigType[] = [
    {
        name: 'auth-service',
        path: '/auth',
        url: env.AUTH_SERVICE_URL,
        pathRewrite: { '^/auth': '/api/v1/auth' },
        timeout: 5000,
    },
    {
        name: 'user-service',
        path: '/users',
        url: env.USER_SERVICE_URL,
        pathRewrite: { '^/users': '/api/v1/users' },
    },
    {
        name: 'notes-service',
        path: '/notes',
        url: env.NOTES_SERVICE_URL,
        pathRewrite: { '^/notes': '/api/v1/notes' },
    },
];

function createProxyOptions(service: ServiceConfigType): Options {
    return {
        target: service.url,
        changeOrigin: true,
        pathRewrite: service.pathRewrite,
        timeout: service.timeout ?? env.DEFAULT_TIMEOUT,
        logger: logger,
        on: {
            error: (err: Error, req: any, res: any): void => {
                logger.error(`Proxy error: ${err.message}`);

                if (!res.headersSent) {
                    res.status(503).end(
                        JSON.stringify({
                            success: false,
                            message: 'Service unavailable.',
                        }),
                    );
                }
            },
            proxyReq: (req: any, proxyReq: any): void => {
                logger.info(
                    `Proxying request: ${req.method} ${req.originalUrl} to ${service.url}`,
                );
            },
            proxyRes: (req: any, proxyRes: any): void => {
                logger.info(
                    `Received response from ${service.url}: ${proxyRes.statusCode} for ${req.method} ${req.originalUrl}`,
                );
            },
        },
    };
}

function setupProxy(app: Router): void {
    serviceConfigs.forEach((service) => {
        const proxyOptions = createProxyOptions(service);
        app.use(service.path, createProxyMiddleware(proxyOptions));
        logger.info(`Configured proxy for ${service.name} at ${service.path}`);
    });
}

export { setupProxy };
