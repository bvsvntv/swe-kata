import client, { Registry } from '@prometheus-io/client';
import { DBLabels } from '../types';

function createDBMetrics<T>(register: Registry, service: string) {
    const duration = new client.Histogram({
        name: 'db_query_duration_seconds',
        help: 'Duration of database queries in seconds',
        labelNames: ['service', 'model', 'operation'],
        buckets: [0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5],
        registers: [register],
    });

    const total = new client.Counter({
        name: 'db_query_total',
        help: 'Total number of database queries',
        labelNames: ['service', 'model', 'operation'],
        registers: [register],
    });

    const errors = new client.Counter({
        name: 'db_errors_total',
        help: 'Total number of failed database queries',
        labelNames: ['service', 'model', 'operation'],
        registers: [register],
    });

    async function measure<T>(
        model: string,
        operation: string,
        fn: () => Promise<T>,
    ): Promise<T> {
        const labels: DBLabels = { service, model, operation };
        const start = process.hrtime.bigint();

        try {
            total.inc(labels);
            const result = await fn();
            duration.observe(
                labels,
                Number(process.hrtime.bigint() - start) / 1e9,
            );
            return result;
        } catch (err) {
            errors.inc(labels);
            duration.observe(
                labels,
                Number(process.hrtime.bigint() - start) / 1e9,
            );
            throw err;
        }
    }

    return { measure };
}

export { createDBMetrics };
