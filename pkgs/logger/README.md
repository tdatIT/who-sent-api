# logger

**logger** is a library used to print logs in a specified format that supports centralized logging. It is developed based on the `uber-zap` library.

![Example Import](example.png)

### Log Format

The log output follows this format:

```json
{"level":"warn","time":"2024-09-11T17:37:18.378+0700","caller":"v2@v2.52.5/router.go:145","func":"github.com/gofiber/fiber/v2.(*App).next","msg":"Client error","service_name":"example-service","ip":"127.0.0.1","latency":"19.334µs","status":404,"method":"GET","url":"/sdsad"}
```

```json
{"level":"info","time":"2024-09-11T17:36:52.743+0700","caller":"example/main.go:16","func":"main.main","msg":"Hello","service_name":"example-service"}
```
### Installation
To install, use the following command:

```bash
make install
```
### Usage
You can either create a new instance of logger or use the singleton instance that is initialized when the program starts.

Creating a New Logger Instance
```go
logger := NewLogger(&LogConfig{
    Level:      "debug",
    LogFormat:  JsonFormat,
    TimeFormat: ISO8601TimeEncoder,
    Filename:   "",
    ServiceName: "example-service",
})
```
You can also use the default configuration:

```go
logger := NewLogger(DefaultConfig)
```
The singleton logger supports both SugaredLogger for formatted output and Zap's regular logger for structured logging.

```go

logger.Debugf("Hello %s", "world")
logger.Debug("Hello world", zap.String("key", "value"))
logger.Infof("Hello %s", "world")
logger.Info("Hello world", zap.String("key", "value"))
logger.Warnf("Hello %s", "world")
logger.Warn("Hello world", zap.String("key", "value"))
logger.Errorf("Hello %s", "world")
logger.Error("Hello world", zap.String("key", "value"))
logger.DPanicf("Hello %s", "world")
logger.DPanic("Hello world", zap.String("key", "value"))
logger.Panicf("Hello %s", "world")
logger.Panic("Hello world", zap.String("key", "value"))
logger.Fatalf("Hello %s", "world")
logger.Fatal("Hello world", zap.String("key", "value"))
```

Updating the Singleton Logger

```go
newlogger := logger.NewLogger(&logger.LogConfig{
    ServiceName: "example-service",
    Level:       "Error",
    LogFormat:   logger.ConsoleFormat,
    TimeFormat:  logger.RFC3339NanoTimeEncoder,
    Filename:    "trace.log",
})
logger.SetLogger(newlogger)
```

### Integrating with Frameworks
For frameworks that can integrate with uber-zap, you can use the initialized logger. Below is an example using Fiber-v2:

```go
newlogger := logger.NewLogger(&logger.LogConfig{
    ServiceName: "example-service",
    Level:       "Error",
    LogFormat:   logger.ConsoleFormat,
    TimeFormat:  logger.RFC3339NanoTimeEncoder,
    Filename:    "trace.log",
})

app := fiber.New()
app.Use(fiberzap.New(fiberzap.Config{
    Logger: newlogger.GetZapInstance(),
}))

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

app.Listen(":3000")

```
For the singleton instance, you can access the Zap instance with ```logger.GetZapInstance()```

<hr>
Developed by tdat.it2k2@gmail.com