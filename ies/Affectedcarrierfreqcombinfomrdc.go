package ies

// AffectedCarrierFreqCombInfoMRDC ::= SEQUENCE
type Affectedcarrierfreqcombinfomrdc struct {
	Victimsystemtype            Victimsystemtype
	Interferencedirectionmrdc   AffectedcarrierfreqcombinfomrdcInterferencedirectionmrdc
	Affectedcarrierfreqcombmrdc *AffectedcarrierfreqcombinfomrdcAffectedcarrierfreqcombmrdc
}
