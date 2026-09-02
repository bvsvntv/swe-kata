import { Registry } from '@prometheus-io/client';

// Create a Registry to hold all metrics
// The registry is responsible for collecting and exposing metrics
export const metricRegistry = new Registry();
