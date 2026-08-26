type ServiceConfigType = {
    name: string;
    path: string;
    url: string;
    pathRewrite: Record<string, string>;
    timeout?: number;
};

export { ServiceConfigType };
