# Tsumugi

This is a small utility to modify the titles of recordings made with various Panasonic TVs.
It is confirmed to work with the model "TX-L37D25E".
Other models from around that era (2010) should work as well.

## Building

```bash
go build cmd/tsumugi.go
```

## Usage

1. Connect the drive to your computer.
2. Mount the drive. The file system of the drive's partition is UFS2. This repository contains a helper script called `mnt.sh` in case you don't want to type out the whole mount command by yourself.
3. Run `tsumugi` and supply all the `.dat` files you want to modify as arguments.
4. You'll be prompted to enter a new title for each file. If you supply an empty title, it won't be changed.
5. Unmount the drive once you're finished. This repository contains a helper script for this called `umt.sh`.

## Useful links
- https://www.hackerboard.de/threads/verschlüsselung-knacken-von-tv-aufnahmen-panasonic-gw20-v20.41849/page-2
- https://www.elektroda.pl/rtvforum/topic2137030.html

## License

This project is licensed under the BSD 3-Clause License. See [LICENSE](LICENSE) for more details.
