package ies

import "rrc/utils"

// MTC-SSB-NR-r15-ssb-Duration-r15 ::= ENUMERATED
type MtcSsbNrR15SsbDurationR15 struct {
	Value utils.ENUMERATED
}

const (
	MtcSsbNrR15SsbDurationR15EnumeratedNothing = iota
	MtcSsbNrR15SsbDurationR15EnumeratedSf1
	MtcSsbNrR15SsbDurationR15EnumeratedSf2
	MtcSsbNrR15SsbDurationR15EnumeratedSf3
	MtcSsbNrR15SsbDurationR15EnumeratedSf4
	MtcSsbNrR15SsbDurationR15EnumeratedSf5
)
