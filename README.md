# s3-event-logger

Provides an simple endpoint for s3 event notifications

Install:

```bash
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/s3-event-logger/main/samples/eventLoggerDeployment.yaml
> kubectl apply -f https://raw.githubusercontent.com/mwennrich/s3-event-logger/main/samples/eventLoggerService.yaml
```

Watch incoming events:

```bash
> stern s3-event-logger
```
