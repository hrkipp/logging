# Context Based Logging
[![Build Status](https://travis-ci.org/hrkipp/logging.svg)](https://travis-ci.org/hrkipp/logging)
This is a simple context based logger for go.

## Usage

If all you want is to simply log to standard out using a pretty simple format then you can simply 
```
ctx := context.Background()
logger.Info(ctx, "%s : %s", "foo", "bar")
```

and you will see 
```
INFO foo : bar
```
logged to stdout. You can also provide a custom writer for the
logger, and it will log to that instead.
```
var logs bytes.Buffer
ctx := logging.WithWriter(context.Background(), &logs)
logger.Info(ctx, "foo")
```

Currently only `INFO`, `WARNING`, and `ERROR` are supported, since i'm lazy, but feel free to fork it and add more. 

If you want to disable output for logging below a threshold:
```
ctx := logging.WithLevel(context.Background(), logging.WARNING)
// this wont print anything
logger.Info(ctx, "foo")
// but this will
logger.Warning(ctx, "foo")
```

You can also define custom formatting to the output strings. Simply define a Builder function and add it to the context.
```
builder := func(ctx context.Context, level logging.Level, args ...interface{}) []byte {
    return []byte(fmt.Sprintf("custom %s %v", level.String(), args))
}
ctx := logging.WithBuilder(context.Background(), builder)
// prints out "custom INFO [foo]"
logging.Info(ctx, "foo")
```

## License

Logging is MIT-Licensed
