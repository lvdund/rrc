package ies

import "rrc/utils"

// LoggingDuration-r10 ::= ENUMERATED
type LoggingdurationR10 struct {
	Value utils.ENUMERATED
}

const (
	LoggingdurationR10EnumeratedNothing = iota
	LoggingdurationR10EnumeratedMin10
	LoggingdurationR10EnumeratedMin20
	LoggingdurationR10EnumeratedMin40
	LoggingdurationR10EnumeratedMin60
	LoggingdurationR10EnumeratedMin90
	LoggingdurationR10EnumeratedMin120
	LoggingdurationR10EnumeratedSpare2
	LoggingdurationR10EnumeratedSpare1
)
