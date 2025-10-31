package ies

import "rrc/utils"

// BandNR-bfd-Relaxation-r17 ::= ENUMERATED
type BandnrBfdRelaxationR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrBfdRelaxationR17EnumeratedNothing = iota
	BandnrBfdRelaxationR17EnumeratedSupported
)
