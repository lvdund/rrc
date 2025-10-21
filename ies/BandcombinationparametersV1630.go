package ies

import "rrc/utils"

// BandCombinationParameters-v1630 ::= SEQUENCE
type BandcombinationparametersV1630 struct {
	V2xSupportedtxbandcomblistperbcV1630 *utils.BITSTRING
	V2xSupportedrxbandcomblistperbcV1630 *utils.BITSTRING
	ScalingfactortxsidelinkR16           *interface{}
	ScalingfactorrxsidelinkR16           *interface{}
	InterbandpowersharingsyncdapsR16     *utils.ENUMERATED
	InterbandpowersharingasyncdapsR16    *utils.ENUMERATED
}
