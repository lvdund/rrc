package ies

import "rrc/utils"

// BandNR-prs-MeasurementWithoutMG-r17 ::= ENUMERATED
type BandnrPrsMeasurementwithoutmgR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPrsMeasurementwithoutmgR17EnumeratedNothing = iota
	BandnrPrsMeasurementwithoutmgR17EnumeratedCplength
	BandnrPrsMeasurementwithoutmgR17EnumeratedQuartersymbol
	BandnrPrsMeasurementwithoutmgR17EnumeratedHalfsymbol
	BandnrPrsMeasurementwithoutmgR17EnumeratedHalfslot
)
