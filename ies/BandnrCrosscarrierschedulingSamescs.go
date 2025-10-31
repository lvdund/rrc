package ies

import "rrc/utils"

// BandNR-crossCarrierScheduling-SameSCS ::= ENUMERATED
type BandnrCrosscarrierschedulingSamescs struct {
	Value utils.ENUMERATED
}

const (
	BandnrCrosscarrierschedulingSamescsEnumeratedNothing = iota
	BandnrCrosscarrierschedulingSamescsEnumeratedSupported
)
