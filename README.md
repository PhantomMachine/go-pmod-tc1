# Introduction
This application can read the temperature from an attached SPI peripheral, and exposes it on an HTTP endpoint. It was designed targeting the Digilent PmodTC1, using [this hardware design](https://github.com/phantommachine/tc1-hw) and [this Petalinux project](https://github.com/phantommachine/tc1-os).

## Build
This project can be cross-compiled by running `GOARCH=arm GOARM=7 go build`.
