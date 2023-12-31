"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.BaseServiceV2 = void 0;
const bcfg_1 = __importDefault(require("bcfg"));
const dotenv = __importStar(require("dotenv"));
const commander_1 = require("commander");
const envalid_1 = require("envalid");
const snakeCase_1 = __importDefault(require("lodash/snakeCase"));
const express_1 = __importDefault(require("express"));
const prom_client_1 = __importDefault(require("prom-client"));
const express_prom_bundle_1 = __importDefault(require("express-prom-bundle"));
const body_parser_1 = __importDefault(require("body-parser"));
const morgan_1 = __importDefault(require("morgan"));
const logger_1 = require("../common/logger");
const metrics_1 = require("./metrics");
const options_1 = require("./options");
class BaseServiceV2 {
    constructor(params) {
        var _a, _b, _c, _d;
        this.params = params;
        this.loop = params.loop !== undefined ? params.loop : true;
        this.state = {};
        params.optionsSpec = Object.assign(Object.assign({}, params.optionsSpec), options_1.stdOptionsSpec);
        params.metricsSpec = Object.assign(Object.assign({}, params.metricsSpec), (0, metrics_1.makeStdMetricsSpec)(params.optionsSpec));
        const opSnakeCase = (str) => {
            const reg = /l_1|l_2/g;
            const repl = str.includes('l1') ? 'l1' : 'l2';
            return (0, snakeCase_1.default)(str).replace(reg, repl);
        };
        const program = new commander_1.Command().allowUnknownOption(true);
        for (const [optionName, optionSpec] of Object.entries(params.optionsSpec)) {
            if (['useEnv', 'useArgv'].includes(optionName)) {
                continue;
            }
            program.addOption(new commander_1.Option(`--${optionName.toLowerCase()}`, `${optionSpec.desc}`).env(`${opSnakeCase(params.name.replace(/-/g, '_')).toUpperCase()}__${opSnakeCase(optionName).toUpperCase()}`));
        }
        const longestMetricNameLength = Object.keys(params.metricsSpec).reduce((acc, key) => {
            const nameLength = (0, snakeCase_1.default)(key).length;
            if (nameLength > acc) {
                return nameLength;
            }
            else {
                return acc;
            }
        }, 0);
        program.addHelpText('after', `\nMetrics:\n${Object.entries(params.metricsSpec)
            .map(([metricName, metricSpec]) => {
            const parsedName = opSnakeCase(metricName);
            return `  ${parsedName}${' '.repeat(longestMetricNameLength - parsedName.length + 2)}${metricSpec.desc} (type: ${metricSpec.type.name})`;
        })
            .join('\n')}
      `);
        program.parse();
        dotenv.config();
        const config = new bcfg_1.default(params.name);
        config.load({
            env: (_b = (_a = params.options) === null || _a === void 0 ? void 0 : _a.useEnv) !== null && _b !== void 0 ? _b : true,
            argv: (_d = (_c = params.options) === null || _c === void 0 ? void 0 : _c.useEnv) !== null && _d !== void 0 ? _d : true,
        });
        const lowerCaseOptions = Object.entries(params.options).reduce((acc, [key, val]) => {
            acc[key.toLowerCase()] = val;
            return acc;
        }, {});
        const cleaned = (0, envalid_1.cleanEnv)(Object.assign(Object.assign(Object.assign({}, config.env), config.args), (lowerCaseOptions || {})), Object.entries(params.optionsSpec || {}).reduce((acc, [key, val]) => {
            acc[key.toLowerCase()] = val.validator({
                desc: val.desc,
                default: val.default,
            });
            return acc;
        }, {}));
        this.options = Object.keys(params.optionsSpec || {}).reduce((acc, key) => {
            acc[key] = cleaned[key.toLowerCase()];
            return acc;
        }, {});
        for (const [optionName, optionSpec] of Object.entries(params.optionsSpec)) {
            if (optionSpec.default === undefined &&
                this.options[optionName] === undefined) {
                throw new Error(`missing required option: ${optionName}`);
            }
        }
        this.metrics = Object.keys(params.metricsSpec || {}).reduce((acc, key) => {
            const spec = params.metricsSpec[key];
            acc[key] = new spec.type({
                name: `${opSnakeCase(params.name)}_${opSnakeCase(key)}`,
                help: spec.desc,
                labelNames: spec.labels || [],
            });
            return acc;
        }, {});
        this.metricsRegistry = prom_client_1.default.register;
        this.port = this.options.port;
        this.hostname = this.options.hostname;
        this.healthy = true;
        this.loopIntervalMs = this.options.loopIntervalMs;
        this.logger = new logger_1.Logger({
            name: params.name,
            level: this.options.logLevel,
        });
        const maxSignalCount = 3;
        let currSignalCount = 0;
        const stop = async (signal) => {
            currSignalCount++;
            if (currSignalCount === 1) {
                this.logger.info(`stopping service with signal`, { signal });
                await this.stop();
                process.exit(0);
            }
            else if (currSignalCount >= maxSignalCount) {
                this.logger.info(`performing hard stop`);
                process.exit(0);
            }
            else {
                this.logger.info(`send ${maxSignalCount - currSignalCount} more signal(s) to hard stop`);
            }
        };
        process.on('SIGTERM', stop);
        process.on('SIGINT', stop);
        this.metrics.metadata.set(Object.assign({ name: params.name, version: params.version }, (0, options_1.getPublicOptions)(params.optionsSpec).reduce((acc, key) => {
            if (key in options_1.stdOptionsSpec) {
                acc[key] = this.options[key].toString();
            }
            else {
                acc[key] = config.str(key);
            }
            return acc;
        }, {})), 1);
        prom_client_1.default.collectDefaultMetrics({
            register: this.metricsRegistry,
            labels: { name: params.name, version: params.version },
        });
    }
    async run() {
        var _a;
        if (!this.server) {
            this.logger.info('starting app server');
            const app = (0, express_1.default)();
            app.use(body_parser_1.default.urlencoded({ extended: true }));
            app.use(body_parser_1.default.json(Object.assign({ verify: (req, res, buf, encoding) => {
                    ;
                    req.rawBody = (buf === null || buf === void 0 ? void 0 : buf.toString(encoding || 'utf8')) || '';
                } }, ((_a = this.params.bodyParserParams) !== null && _a !== void 0 ? _a : {}))));
            app.use((0, morgan_1.default)('short', {
                stream: {
                    write: (str) => {
                        this.logger.info(`server log`, {
                            log: str,
                        });
                    },
                },
            }));
            app.get('/healthz', async (req, res) => {
                return res.json({
                    ok: this.healthy,
                    version: this.params.version,
                });
            });
            const router = express_1.default.Router();
            if (this.routes) {
                this.routes(router);
            }
            app.use((0, express_prom_bundle_1.default)({
                promRegistry: this.metricsRegistry,
                includeMethod: true,
                includePath: true,
                includeStatusCode: true,
                normalizePath: (req) => {
                    for (const layer of router.stack) {
                        if (layer.route && req.path.match(layer.regexp)) {
                            return layer.route.path;
                        }
                    }
                    return '/invalid_path_not_a_real_route';
                },
            }));
            app.use('/api', router);
            await new Promise((resolve) => {
                this.server = app.listen(this.port, this.hostname, () => {
                    resolve(null);
                });
            });
            this.logger.info(`app server started`, {
                port: this.port,
                hostname: this.hostname,
            });
        }
        if (this.init) {
            this.logger.info('initializing service');
            await this.init();
            this.logger.info('service initialized');
        }
        if (this.loop) {
            this.logger.info('starting main loop');
            this.running = true;
            const doLoop = async () => {
                try {
                    this.mainPromise = this.main();
                    await this.mainPromise;
                }
                catch (err) {
                    this.metrics.unhandledErrors.inc();
                    this.logger.error('caught an unhandled exception', {
                        message: err.message,
                        stack: err.stack,
                        code: err.code,
                    });
                }
                if (this.running) {
                    this.pollingTimeout = setTimeout(doLoop, this.loopIntervalMs);
                }
            };
            doLoop();
        }
        else {
            this.logger.info('running main function');
            await this.main();
        }
    }
    async stop() {
        this.logger.info('stopping main loop...');
        this.running = false;
        clearTimeout(this.pollingTimeout);
        this.logger.info('waiting for main to complete');
        await this.mainPromise;
        this.logger.info('main loop stoped.');
        if (this.server) {
            this.logger.info('stopping metrics server');
            await new Promise((resolve) => {
                this.server.close(() => {
                    resolve(null);
                });
            });
            this.logger.info('metrics server stopped');
            this.server = undefined;
        }
    }
}
exports.BaseServiceV2 = BaseServiceV2;
//# sourceMappingURL=base-service-v2.js.map