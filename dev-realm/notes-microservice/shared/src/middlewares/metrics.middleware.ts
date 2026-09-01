import { NextFunction, Request, Response } from 'express';
import client, {
    Registry,
    Counter,
    Histogram,
    Gauge,
} from '@prometheus-io/client';

// Returns the Express middleware AND the registered metrics so the caller
// can keep references to them. Per-service registry is passed in.
function createMetricsMiddleware(register: Registry, service: string) {
    const requestTotal = new client.Counter({
        name: 'http_requests_total',
        help: 'Total number of HTTP requests',
        labelNames: ['service', 'method', 'route', 'status'],
        registers: [register],
    });

    const requestDuration = new client.Histogram({
        name: 'http_request_duration_seconds',
        help: 'HTTP request duration in seconds',
        labelNames: ['service', 'method', 'route'],
        buckets: [0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10],
        registers: [register],
    });

    const inFlight = new client.Gauge({
        name: 'http_requests_in_flight',
        help: 'HTTP requests currently being processed',
        labelNames: ['service', 'method'],
        registers: [register],
    });

    return function metricsMiddleware(
        req: Request,
        res: Response,
        next: NextFunction,
    ) {
        const start = process.hrtime.bigint();

        // Prefer the route pattern (e.g. /notes/:id) over the raw URL so labels
        // don't explode with dynamic ids. Fall back to originalUrl when no route.
        const route = req.baseUrl + (req.route?.path ?? req.path);

        inFlight.inc({ service, method: req.method });

        // Collect metrics after response is sent
        res.on('finish', () => {
            const seconds = Number(process.hrtime.bigint() - start) / 1e9;
            const status = String(res.statusCode);

            requestTotal.inc({ service, method: req.method, route, status });
            requestDuration.observe(
                { service, method: req.method, route },
                seconds,
            );
            inFlight.dec({ service, method: req.method });
        });

        next();
    };
}

export { createMetricsMiddleware };
