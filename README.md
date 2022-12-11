# Standard Context
Standard Context (`stdctx`) is a Go library for associating implicit dependencies such as logging, filesystems, and network access with `context.Context` objects.

Standard contexts are a convention for keeping code well-behaved.
A library or function instrumented with `stdctx` is not going to annoy the caller with weird dependencies or logs that can't be turned off.
Instead code using `stdctx` will produce no logging or telemetry side effects by default.

If the caller decides it wants more information about how a `stdctx` instrumented function is behaving, then it can setup logging or telemetry on the `context.Context` and pass it to the callee.

## Packages
- `logctx` logging using `slog`
- `telctx` telemetry
- `netctx` networking
