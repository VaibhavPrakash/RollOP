/// <reference types="node" />
import { Server } from 'net';
import { Registry } from 'prom-client';
import bodyParser from 'body-parser';
import { ExpressRouter } from './router';
import { Logger } from '../common/logger';
import { Metrics, MetricsSpec, StandardMetrics } from './metrics';
import { Options, OptionsSpec, StandardOptions } from './options';
export declare abstract class BaseServiceV2<TOptions extends Options, TMetrics extends Metrics, TServiceState> {
    private readonly params;
    private pollingTimeout;
    private mainPromise;
    protected loop: boolean;
    protected loopIntervalMs: number;
    protected running: boolean;
    protected healthy: boolean;
    protected logger: Logger;
    protected state: TServiceState;
    protected readonly options: TOptions & StandardOptions;
    protected readonly metrics: TMetrics & StandardMetrics;
    protected readonly metricsRegistry: Registry;
    protected server: Server;
    protected readonly port: number;
    protected readonly hostname: string;
    constructor(params: {
        name: string;
        version: string;
        optionsSpec: OptionsSpec<TOptions>;
        metricsSpec: MetricsSpec<TMetrics>;
        options?: Partial<TOptions & StandardOptions>;
        loop?: boolean;
        bodyParserParams?: bodyParser.OptionsJson;
    });
    run(): Promise<void>;
    stop(): Promise<void>;
    protected init?(): Promise<void>;
    protected routes?(router: ExpressRouter): Promise<void>;
    protected abstract main(): Promise<void>;
}
