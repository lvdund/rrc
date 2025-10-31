package ies

import "rrc/utils"

// LoggingInterval-r10 ::= ENUMERATED
type LoggingintervalR10 struct {
	Value utils.ENUMERATED
}

const (
	LoggingintervalR10EnumeratedNothing = iota
	LoggingintervalR10EnumeratedMs1280
	LoggingintervalR10EnumeratedMs2560
	LoggingintervalR10EnumeratedMs5120
	LoggingintervalR10EnumeratedMs10240
	LoggingintervalR10EnumeratedMs20480
	LoggingintervalR10EnumeratedMs30720
	LoggingintervalR10EnumeratedMs40960
	LoggingintervalR10EnumeratedMs61440
)
