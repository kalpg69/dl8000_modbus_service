package main

import (
	"encoding/binary"

	"github.com/goburrow/modbus"
)

const (
	//Auto is BCU Operating Mode
	Auto = 0
	//Manual is BCU Operating Mode
	Manual = 1
)

type RegInt struct {
	OperatingMode  uint16
	BatchingStatus uint16
}

type mdbusClient struct {
	regInt *RegInt
	client modbus.Client
}

func (m mdbusClient) GetCurrentMode() (*uint16, error) {
	byteResult, err := m.client.ReadHoldingRegisters(m.regInt.OperatingMode, 1)
	if err != nil {
		return nil, err
	}
	currentMode := binary.BigEndian.Uint16(byteResult)
	return &currentMode, nil
}

func (m mdbusClient) GetBatchingStatus() (*uint16, error) {
	byteResult, err := m.client.ReadHoldingRegisters(m.regInt.BatchingStatus, 1)
	if err != nil {
		return nil, err
	}

	currentBatchingStatus := binary.BigEndian.Uint16(byteResult)
	return &currentBatchingStatus, nil
}

func (m mdbusClient) ChangeOperatingMode() error {
	currentMode, err := m.GetCurrentMode()
	if err != nil {
		return err
	}

	switch *currentMode {
	case Auto:
		_, err = m.client.WriteSingleRegister(m.regInt.OperatingMode, Manual)

		if err != nil {
			return err
		}

	case Manual:

		_, err = m.client.WriteSingleRegister(m.regInt.OperatingMode, Auto)

		if err != nil {
			return err
		}

	default:
		_, err = m.client.WriteSingleRegister(m.regInt.OperatingMode, Auto)

		if err != nil {
			return err
		}

	}

	return nil
}

func (m mdbusClient) DisplayBCUIdleMessage() error {
	currentMode, err := m.GetCurrentMode()
	if err != nil {
		return err
	}

	currentBatchingStatus, err := m.GetBatchingStatus()
	if err != nil {
		return err
	}

}
