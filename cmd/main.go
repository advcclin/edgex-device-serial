// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 Advantech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-serial"
	"github.com/edgexfoundry/device-serial/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	version     string = device_serial.Version
	serviceName string = "device-serial"
)

func main() {
	d := driver.SerialDriver{}
	startup.Bootstrap(serviceName, version, &d)
}
