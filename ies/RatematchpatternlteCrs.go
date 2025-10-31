package ies

import "rrc/utils"

// RateMatchPatternLTE-CRS ::= SEQUENCE
type RatematchpatternlteCrs struct {
	Carrierfreqdl           utils.INTEGER `lb:0,ub:16383`
	Carrierbandwidthdl      RatematchpatternlteCrsCarrierbandwidthdl
	MbsfnSubframeconfiglist *EutraMbsfnSubframeconfiglist
	NrofcrsPorts            RatematchpatternlteCrsNrofcrsPorts
	VShift                  RatematchpatternlteCrsVShift
}
