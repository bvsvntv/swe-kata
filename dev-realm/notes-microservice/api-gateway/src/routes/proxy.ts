import { env } from '@/config/env.config';
import { logger } from '@/lib/logger';
import { ServiceConfigType } from '@/types';
import { Application } from 'express';
import { createProxyMiddleware, Options } from 'http-proxy-middleware';

const serviceConfigs: ServiceConfigType[] = [
    {
        name: 'auth-service',
        path: '/api/v1/auth/',
        url: env.AUTH_SERVICE_URL,
        pathRewrite: { '^/api/v1/auth': '/api/v1/auth/' },
        timeout: 5000,
    },
    {
        name: 'user-service',
        path: '/api/v1/user/',
        url: env.USER_SERVICE_URL,
        pathRewrite: { '^/api/v1/user': '/api/v1/user/' },
    },
    {
        name: 'notes-service',
        path: '/api/v1/notes/',
        url: env.NOTES_SERVICE_URL,
        pathRewrite: { '^/api/v1/notes': '/api/v1/notes/' },
    },
];

function createProxyOptions(service: ServiceConfigType): Options {
    return {
        target: service.url,
        changeOrigin: true,
        pathRewrite: service.pathRewrite,
        timeout: service.timeout ?? 3000,
        logger: logger,
        on: {
            error: handleProxyError,
            // proxyReq: handleProxyRequest,
            // proxyRes: handleProxyResponse,
        },
    };
}

function handleProxyError(err: Error, req: any, res: any): void {
    logger.error(`Proxy error: ${err.message}`);

    if (!res.headersSent) {
        res.status(503).end(
            JSON.stringify({
                success: false,
                message: 'Service unavailable.',
            }),
        );
    }
}

// function handleProxyRequest(proxyReq: any, req: any): void {}
// function handleProxyResponse(proxyRes: any, res: any): void {}

function setupProxy(app: Application): void {
    serviceConfigs.forEach((service) => {
        const proxyOptions = createProxyOptions(service);
        app.use(service.path, createProxyMiddleware(proxyOptions));
        logger.info(`Configured proxy for ${service.name} at ${service.path}`);
    });
}

export { setupProxy };
