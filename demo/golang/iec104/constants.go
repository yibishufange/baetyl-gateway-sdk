package iec104

import "github.com/baetyl/baetyl-go/v2/errors"

var (
	ErrClientInvalid        = errors.New("device client is invalid")
	ErrUnsupportedFieldType = errors.New("unsupported field type")

	ErrWorkerNotExist     = errors.New("worker not exist")
	ErrDriverNameNotExist = errors.New("failed to get driverName in msg")
	ErrDevNameNotExist    = errors.New("failed to get deviceName in msg")
	ErrGetDeviceProps     = errors.New("failed to get device props in msg")
)

const (
	SlaveOffline = 0
	SlaveOnline  = 1

	DefaultAntsPoolSize = 1000

	DI_Start = "DI_Start"
	AI_Start = "AI_Start"
	PI_Start = "PI_Start"
	DO_Start = "DO_Start"
	AO_Start = "AO_Start"

	AI = "AI"
	DI = "DI"
	PI = "PI"
	DO = "DO"
	AO = "AO"
)
