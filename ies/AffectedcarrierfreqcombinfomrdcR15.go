package ies

import "rrc/utils"

// AffectedCarrierFreqCombInfoMRDC-r15 ::= SEQUENCE
type AffectedcarrierfreqcombinfomrdcR15 struct {
	VictimsystemtypeR15            VictimsystemtypeR11
	InterferencedirectionmrdcR15   utils.ENUMERATED
	AffectedcarrierfreqcombmrdcR15 *interface{}
}
