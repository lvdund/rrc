package ies

import "rrc/utils"

// BandNR-parallelMeasurementWithoutRestriction-r17 ::= ENUMERATED
type BandnrParallelmeasurementwithoutrestrictionR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrParallelmeasurementwithoutrestrictionR17EnumeratedNothing = iota
	BandnrParallelmeasurementwithoutrestrictionR17EnumeratedSupported
)
