package ies

// AffectedCarrierFreqCombInfoMRDC-r15 ::= SEQUENCE
type AffectedcarrierfreqcombinfomrdcR15 struct {
	VictimsystemtypeR15            VictimsystemtypeR11
	InterferencedirectionmrdcR15   AffectedcarrierfreqcombinfomrdcR15InterferencedirectionmrdcR15
	AffectedcarrierfreqcombmrdcR15 *AffectedcarrierfreqcombinfomrdcR15AffectedcarrierfreqcombmrdcR15
}
