package main

import (
    "github.com/dunhamsteve/ios/backup"
    "github.com/trymoose/errors"
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
    defer errors.Do(r.Close)
    w := CreateOutputFile(rec)
    defer errors.Do(w.Close)
    errors.Get(io.Copy(w, r))
}

func CreateOutputFile(rec *backup.Record) io.WriteCloser {
    fp := filepath.Join(Args().OutputDir, errors.Get(filepath.Localize(rec.Path)))
    fd, _ := filepath.Split(fp)
    errors.Check(os.MkdirAll(fd, os.ModePerm))
    o := errors.Get(os.Create(fp))
    return o
}

func OpenDBFile(rec *backup.Record) io.ReadCloser {
    r, err := Backup().FileReader(*rec)
    if err != nil && !errors.Is(err, io.EOF) {
        errors.Check(err)
    }
    return r
}

func FilterRecs(bk *backup.MobileBackup) iter.Seq[*backup.Record] {
    return func(yield func(*backup.Record) bool) {
        for _, rec := range bk.Records {
            if RecOk(&rec) {
                if !yield(&rec) {
                    return
                }
            }
        }
    }
}

func RecOk(rec *backup.Record) bool {
    return rec.Length > 0 && rec.Domain != "" && rec.Path != ""
}
