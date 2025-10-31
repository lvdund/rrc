package ies

import "rrc/utils"

// BandNR-maxDurationDMRS-Bundling-r17-tdd-r17 ::= ENUMERATED
type BandnrMaxdurationdmrsBundlingR17TddR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxdurationdmrsBundlingR17TddR17EnumeratedNothing = iota
	BandnrMaxdurationdmrsBundlingR17TddR17EnumeratedN2
	BandnrMaxdurationdmrsBundlingR17TddR17EnumeratedN4
	BandnrMaxdurationdmrsBundlingR17TddR17EnumeratedN8
	BandnrMaxdurationdmrsBundlingR17TddR17EnumeratedN16
)
