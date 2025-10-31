package ies

import "rrc/utils"

// BandNR-maxDurationDMRS-Bundling-r17-fdd-r17 ::= ENUMERATED
type BandnrMaxdurationdmrsBundlingR17FddR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxdurationdmrsBundlingR17FddR17EnumeratedNothing = iota
	BandnrMaxdurationdmrsBundlingR17FddR17EnumeratedN4
	BandnrMaxdurationdmrsBundlingR17FddR17EnumeratedN8
	BandnrMaxdurationdmrsBundlingR17FddR17EnumeratedN16
	BandnrMaxdurationdmrsBundlingR17FddR17EnumeratedN32
)
