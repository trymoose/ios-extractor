package main

import (
	"errors"
	"github.com/dunhamsteve/ios/backup"
	"io"
	"iter"
	"log/slog"
	"os"
	"path/filepath"
)

func ExtractAllFiles() {
	for rec := range FilterRecs(Backup()) {
		DecodeRec(rec)
	}
}

func DecodeRec(rec *backup.Record) {
	slog.Info("decoding", slog.String("domain", rec.Domain), slog.String("path", rec.Path), slog.Uint64("length", rec.Length))
	r := OpenDBFile(rec)
	defer Defer(r.Close)
	w := CreateOutputFile(rec)
	defer Defer(w.Close)
	Get(io.Copy(w, r))
}

func CreateOutputFile(rec *backup.Record) io.WriteCloser {
	fp := filepath.Join(Args().OutputDir, Get(filepath.Localize(rec.Path)))
	fd, _ := filepath.Split(fp)
	Must(os.MkdirAll(fd, os.ModePerm))
	o := Get(os.Create(fp))
	return o
}

func OpenDBFile(rec *backup.Record) io.ReadCloser {
	r, err := Backup().FileReader(*rec)
	if err != nil && !errors.Is(err, io.EOF) {
		Must(err)
	}
	return r
}

func FilterRecs(bk *backup.MobileBackup) iter.Seq[*backup.Record] {
	return func(yield func(*backup.Record) bool) {
		for _, rec := range bk.Records {
			if rec.Length > 0 && rec.Domain != "" && rec.Path != "" {
				if !yield(&rec) {
					return
				}
			}
		}
	}
}
