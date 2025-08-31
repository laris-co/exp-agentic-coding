package main

import (
    "github.com/pocketbase/pocketbase/core"
)

func handleHello(e *core.RequestEvent) error {
    return e.String(200, "Hello world!")
}

func handleHealthz(e *core.RequestEvent) error {
    return e.JSON(200, map[string]string{
        "status": "ok",
    })
}

func handleVersion(e *core.RequestEvent) error {
    return e.JSON(200, map[string]string{
        "version": version,
        "commit":  commit,
        "builtAt": builtAt,
    })
}

