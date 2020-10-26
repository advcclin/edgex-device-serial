// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Advantech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/edgex-go/pkg/clients/logging"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

type SerialDriver struct {
	lc            logger.LoggingClient
	asyncCh       chan<- *dsModels.AsyncValues
	serialDevices map[string]*serialDevice
}

func (d *SerialDriver) DisconnectDevice(address *models.Addressable) error {
	d.lc.Info(fmt.Sprintf("SerialDriver.DisconnectDevice: device-serial driver is disconnecting to %v", address))
	return nil
}

func (d *SerialDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	d.lc = lc
	d.asyncCh = asyncCh
	d.serialDevices = make(map[string]*serialDevice)
	return nil
}

func (d *SerialDriver) HandleReadCommands(addr *models.Addressable, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	rd, ok := d.serialDevices[addr.Name]
	if !ok {
		rd = newSerialDevice()
		d.serialDevices[addr.Name] = rd
	}

	res = make([]*dsModels.CommandValue, len(reqs))
	now := time.Now().UnixNano() / int64(time.Millisecond)

	for i, req := range reqs {
		t := req.DeviceObject.Properties.Value.Type
		v, err := rd.value(t)
		d.lc.Info(fmt.Sprintf("SerialDriver.HandleReadCommands: value=%d", v))
		if err != nil {
			return nil, err
		}
		var cv *dsModels.CommandValue
		switch t {
		case "Int32":
			cv, _ = dsModels.NewInt32Value(&req.RO, now, int32(v))
		}
		res[i] = cv
	}

	return res, nil
}

func (d *SerialDriver) HandleWriteCommands(addr *models.Addressable, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	rd, ok := d.serialDevices[addr.Name]
	if !ok {
		rd = newSerialDevice()
		d.serialDevices[addr.Name] = rd
	}

	for _, param := range params {
		switch param.RO.Object {
		case "Min_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("SerialDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt32 {
				return fmt.Errorf("SerialDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.minInt32 = int(v)
		case "Max_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("SerialDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt32 {
				return fmt.Errorf("SerialDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.maxInt32 = int(v)
		default:
			return fmt.Errorf("SerialDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
		}
	}

	return nil
}

func (d *SerialDriver) Stop(force bool) error {
	d.lc.Info("SerialDriver.Stop: device-serial driver is stopping...")
	return nil
}
