package ies

import "rrc/utils"

// RetxBSR-Timer-r12 ::= ENUMERATED
type RetxbsrTimerR12 struct {
	Value utils.ENUMERATED
}

const (
	RetxbsrTimerR12EnumeratedNothing = iota
	RetxbsrTimerR12EnumeratedSf320
	RetxbsrTimerR12EnumeratedSf640
	RetxbsrTimerR12EnumeratedSf1280
	RetxbsrTimerR12EnumeratedSf2560
	RetxbsrTimerR12EnumeratedSf5120
	RetxbsrTimerR12EnumeratedSf10240
	RetxbsrTimerR12EnumeratedSpare2
	RetxbsrTimerR12EnumeratedSpare1
)
