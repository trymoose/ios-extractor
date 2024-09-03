# ios-extractor

## Install

```
go install github.com/trymoose/ios-extractor
```

## Usage

```
Usage:
  ios-extractor [OPTIONS]

Application Options:
  -i, --input=         Encrypted iOS backup directory (default: ./input)
  -o, --output=        Decrypted output directory (default: ./output)
  -p, --password=      Encrypted iOS backup password, if not specified will be queried from console
  -a, --allow-partial  Don't delete backup directory on failed extract.'

Help Options:
  -h, --help           Show this help message
```