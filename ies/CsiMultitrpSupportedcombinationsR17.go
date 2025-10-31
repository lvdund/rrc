package ies

import "rrc/utils"

// CSI-MultiTRP-SupportedCombinations-r17 ::= SEQUENCE
type CsiMultitrpSupportedcombinationsR17 struct {
	MaxnumtxPortsR17              CsiMultitrpSupportedcombinationsR17MaxnumtxPortsR17
	MaxtotalnumcmrR17             utils.INTEGER `lb:0,ub:64`
	MaxtotalnumtxPortsnzpCsiRsR17 utils.INTEGER `lb:0,ub:256`
}
