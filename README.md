# Library for tp-link HS100/HS110

Yet another tp-link HS100 library for golang

## Badges

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bce147bd9ade43cbae3b62157cc75aac)](https://app.codacy.com/app/jaedle/golang-tplink-hs100?utm_source=github.com&utm_medium=referral&utm_content=jaedle/golang-tplink-hs100&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://travis-ci.com/jaedle/golang-tplink-hs100.svg?branch=master)](https://travis-ci.com/jaedle/golang-tplink-hs100)
[![Coverage Status](https://coveralls.io/repos/github/jaedle/golang-tplink-hs100/badge.svg?branch=master)](https://coveralls.io/github/jaedle/golang-tplink-hs100?branch=master)
[![GolangCI](https://golangci.com/badges/github.com/jaedle/golang-tplink-hs100.svg)](https://golangci.com)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=golang-tplink-hs100&metric=code_smells)](https://sonarcloud.io/dashboard?id=golang-tplink-hs100)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=golang-tplink-hs100&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=golang-tplink-hs100)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=golang-tplink-hs100&metric=sqale_index)](https://sonarcloud.io/dashboard?id=golang-tplink-hs100)

## Usage

### With a project using go modules

install library `go get github.com/jaedle/golang-tplink-hs100`

### Without using go modules

install library `go get github.com/jaedle/golang-tplink-hs100/...`

### Usage example

use the following code as main and replace `YOUR_HS100_DEVICE` with the 
address of your HS100-device.

```golang
package main

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"os"
)

func main() {
	h := hs100.NewHs100("YOUR_HS100_DEVICE", configuration.Default())

	name, err := h.GetName()
	if err != nil {
		println("Error on accessing device")
		os.Exit(1)
	}

	println("Name of device: " + name)
}
```

## Acknowledgements

-   [tplink-smarthome-api](https://github.com/plasticrake/tplink-smarthome-api): 
    Thanks for the inspiration!

-   [tplink-smarthome-crypto](https://github.com/plasticrake/tplink-smarthome-crypto) 
    Thanks for the excellent documentation/test-cases for encrypting/decrypting 
    the communication

-   [tplink-smarthome-simulator](https://github.com/plasticrake/tplink-smarthome-simulator) 
    Thanks for providing a device simulator for integration tests!

-   [hs1xxplug](https://github.com/sausheong/hs1xxplug): 
    Thanks for the blueprint in golang!

## Development

### Prerequisites

1.  go-task 
1.  docker

## Project structure

This project tries to stick as close as possible to the [golang standard project layout](https://github.com/golang-standards/project-layout)

The public parts for this library are located in `/pkg`.

All files in `/cmd` are for demo purposes only.

## License

[MIT](https://github.com/jaedle/golang-tplink-hs100/blob/master/LICENSE)
