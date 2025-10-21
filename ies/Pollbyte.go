package ies

import "rrc/utils"

// PollByte ::= ENUMERATED
type Pollbyte struct {
	Value utils.ENUMERATED
}

const (
	PollbyteKb25       = 0
	PollbyteKb50       = 1
	PollbyteKb75       = 2
	PollbyteKb100      = 3
	PollbyteKb125      = 4
	PollbyteKb250      = 5
	PollbyteKb375      = 6
	PollbyteKb500      = 7
	PollbyteKb750      = 8
	PollbyteKb1000     = 9
	PollbyteKb1250     = 10
	PollbyteKb1500     = 11
	PollbyteKb2000     = 12
	PollbyteKb3000     = 13
	PollbyteKbinfinity = 14
	PollbyteSpare1     = 15
)
