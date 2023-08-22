# slogctx

Pass log/slog through contexts.

Usually it's an anti-pattern to pass values with the context, though we found
it was tedious to explicitly pass the logger to all structs and it was difficult
to get the logger across certain boundaries (http requests, tracing libraries,
etc.). I wrote a package almost identical to this based on
[zap](https://github.com/uber-go/zap) at previous position and received a lot of
positive feedback. I thought it would be helpful to rewrite it for the recently
announced [slog package](https://go.dev/blog/slog) from the Go team.

## Usage

```go
package main

import (
        "log/slog"

        "github.com/uhthomas/slogctx"
)

func main() {
        ctx := context.Background()

        slogctx.Info(ctx, "default logger", "some key", "some value")

        ctx = slogctx.With(ctx, slog.New(slog.NewTextHandler(nil, nil)))

        slogctx.Info(ctx, "custom logger")

        // get the logger back from the context
        _ = slogctx.From(ctx)
}
```
