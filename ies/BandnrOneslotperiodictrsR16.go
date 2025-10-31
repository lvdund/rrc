package ies

import "rrc/utils"

// BandNR-oneSlotPeriodicTRS-r16 ::= ENUMERATED
type BandnrOneslotperiodictrsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrOneslotperiodictrsR16EnumeratedNothing = iota
	BandnrOneslotperiodictrsR16EnumeratedSupported
)
