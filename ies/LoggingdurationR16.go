package ies

import "rrc/utils"

// LoggingDuration-r16 ::= ENUMERATED
type LoggingdurationR16 struct {
	Value utils.ENUMERATED
}

const (
	LoggingdurationR16EnumeratedNothing = iota
	LoggingdurationR16EnumeratedMin10
	LoggingdurationR16EnumeratedMin20
	LoggingdurationR16EnumeratedMin40
	LoggingdurationR16EnumeratedMin60
	LoggingdurationR16EnumeratedMin90
	LoggingdurationR16EnumeratedMin120
	LoggingdurationR16EnumeratedSpare2
	LoggingdurationR16EnumeratedSpare1
)
