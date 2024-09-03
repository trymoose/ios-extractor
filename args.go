package main

import (
    "github.com/jessevdk/go-flags"
    "github.com/trymoose/errors"
    "os"
    "sync"
)

var Args = sync.OnceValue(func() (args struct {
    BackupDir    string  `short:"i" long:"input" default:"./input" description:"Encrypted iOS backup directory"`
    OutputDir    string  `short:"o" long:"output" default:"./output" description:"Decrypted output directory"`
    Password     *string `short:"p" long:"password" description:"Encrypted iOS backup password, if not specified will be queried from console"`
    AllowPartial bool    `short:"a" long:"allow-partial" description:"Don't delete backup directory on failed extract.'"`
}) {
    p := flags.NewParser(&args, flags.HelpFlag|flags.PassDoubleDash)
    success, defr := errors.OnFail(func() {
        p.WriteHelp(os.Stderr)
        os.Exit(1)
    })
    defer defr()
    errors.Get(p.Parse())
    success()
    return
})
