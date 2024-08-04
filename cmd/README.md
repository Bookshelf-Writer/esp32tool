# CLI esptool

This is a console application for working with esptool-mod

## Build

```bash
go build -ldflags="-s -w" -o esp32tool.bin
chmod +x esp32tool.bin
ls -lh esp32tool.bin
```

## work

### Help
```bash
./esp32tool.bin h
```

### Version
```bash
./esp32tool.bin ver
```

### Serial Device List
```bash
./esp32tool.bin -list
```

---

### Info
```bash
./esp32tool.bin -info -port /dev/ttyACM0 logTrase
```

```bash
./esp32tool.bin -info -port /dev/ttyACM0
```