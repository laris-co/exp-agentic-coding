package main

import (
    "log"
    "log/slog"
    "os"
    "time"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    // basic structured logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        // lightweight request logging middleware
        se.Router.BindFunc(func(e *core.RequestEvent) error {
            start := time.Now()
            err := e.Next()
            dur := time.Since(start)
            slog.Info("request",
                "method", e.Request.Method,
                "path", e.Request.URL.Path,
                "status", e.Status(),
                "duration_ms", dur.Milliseconds(),
                "remote_ip", e.RemoteIP(),
            )
            return err
        })

        // routes
        se.Router.GET("/hello", handleHello)
        se.Router.GET("/healthz", handleHealthz)
        se.Router.GET("/version", handleVersion)

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

