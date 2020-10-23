// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 Advantech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
)

const (
	defMinInt32, defMaxInt32 = -2147483648, 2147483647
)

type serialDevice struct {
	minInt32 int
	maxInt32 int
	currInt int
}

func (d *serialDevice) value(valueType string) (int, error) {
	switch valueType {
	case "Int32":
		if d.maxInt32 <= d.minInt32 {
			return 0, fmt.Errorf("serialDevice.value: maximum: %d <= minimum : %d", d.maxInt32, d.minInt32)
		} else {
			d.currInt = d.currInt + 1
			return d.currInt, nil
		}
	default:
		return 0, fmt.Errorf("serialDevice.value: wrong value type: %s", valueType)
	}
}

func newSerialDevice() *serialDevice {
	return &serialDevice{
		minInt32: defMinInt32,
		maxInt32: defMaxInt32,
		currInt: 0,
	}
}
