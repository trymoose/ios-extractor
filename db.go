package main

import (
	"fmt"
	"github.com/dunhamsteve/ios/backup"
	"github.com/dunhamsteve/ios/keybag"
	"github.com/dunhamsteve/plist"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

var Backup = sync.OnceValue(OpenBackup)

func OpenBackup() *backup.MobileBackup {
	var bk backup.MobileBackup
	bk.Dir = Args().BackupDir
	f := Get(os.Open(filepath.Join(bk.Dir, "Manifest.plist")))
	defer Defer(f.Close)
	Must(plist.Unmarshal(f, &bk.Manifest))
	bk.Keybag = keybag.Read(bk.Manifest.BackupKeyBag)
	LoadDB(&bk)
	return &bk
}

func LoadDB(bk *backup.MobileBackup) {
	if bk.Manifest.IsEncrypted {
		Must(bk.SetPassword(GetPassword()))
	}
	Must(bk.Load())
}

func GetPassword() string {
	if Args().Password != nil {
		return *Args().Password
	}
	Get(fmt.Fprint(os.Stderr, "Backup Password: "))
	pw := Get(terminal.ReadPassword(syscall.Stdin))
	fmt.Println()
	return string(pw)
}
