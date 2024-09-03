package main

import (
    "github.com/trymoose/errors"
    "log/slog"
    "os"
)

func main() {
    errors.Check(os.MkdirAll(Args().OutputDir, os.ModePerm))
    success, defr := errors.OnFail(func() {
        slog.Error("failed to extract backup", slog.Bool("allow-partial", Args().AllowPartial))
        if !Args().AllowPartial {
            errors.Check(os.RemoveAll(Args().OutputDir))
        }
    })
    defer defr()
    ExtractAllFiles()
    success()
}
