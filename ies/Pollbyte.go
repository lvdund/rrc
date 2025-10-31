package ies

import "rrc/utils"

// PollByte ::= ENUMERATED
type Pollbyte struct {
	Value utils.ENUMERATED
}

const (
	PollbyteEnumeratedNothing = iota
	PollbyteEnumeratedKb25
	PollbyteEnumeratedKb50
	PollbyteEnumeratedKb75
	PollbyteEnumeratedKb100
	PollbyteEnumeratedKb125
	PollbyteEnumeratedKb250
	PollbyteEnumeratedKb375
	PollbyteEnumeratedKb500
	PollbyteEnumeratedKb750
	PollbyteEnumeratedKb1000
	PollbyteEnumeratedKb1250
	PollbyteEnumeratedKb1500
	PollbyteEnumeratedKb2000
	PollbyteEnumeratedKb3000
	PollbyteEnumeratedKbinfinity
	PollbyteEnumeratedSpare1
)
