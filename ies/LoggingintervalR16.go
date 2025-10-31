package ies

import "rrc/utils"

// LoggingInterval-r16 ::= ENUMERATED
type LoggingintervalR16 struct {
	Value utils.ENUMERATED
}

const (
	LoggingintervalR16EnumeratedNothing = iota
	LoggingintervalR16EnumeratedMs320
	LoggingintervalR16EnumeratedMs640
	LoggingintervalR16EnumeratedMs1280
	LoggingintervalR16EnumeratedMs2560
	LoggingintervalR16EnumeratedMs5120
	LoggingintervalR16EnumeratedMs10240
	LoggingintervalR16EnumeratedMs20480
	LoggingintervalR16EnumeratedMs30720
	LoggingintervalR16EnumeratedMs40960
	LoggingintervalR16EnumeratedMs61440
	LoggingintervalR16EnumeratedInfinity
)
