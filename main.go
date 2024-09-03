package main

import (
	"log/slog"
	"os"
)

func main() {
	Must(os.MkdirAll(Args().OutputDir, os.ModePerm))
	success, defr := OnFail(func() {
		slog.Error("failed to extract backup", slog.Bool("allow-partial", Args().AllowPartial))
		if !Args().AllowPartial {
			Must(os.RemoveAll(Args().OutputDir))
		}
	})
	defer defr()
	ExtractAllFiles()
	success()
}
