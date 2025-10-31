package ies

import "rrc/utils"

// BandCombinationParameters-v1630 ::= SEQUENCE
type BandcombinationparametersV1630 struct {
	V2xSupportedtxbandcomblistperbcV1630 *utils.BITSTRING            `lb:1,ub:maxBandCombSidelinkNRR16`
	V2xSupportedrxbandcomblistperbcV1630 *utils.BITSTRING            `lb:1,ub:maxBandCombSidelinkNRR16`
	ScalingfactortxsidelinkR16           *[]ScalingfactorsidelinkR16 `lb:1,ub:maxBandCombSidelinkNRR16`
	ScalingfactorrxsidelinkR16           *[]ScalingfactorsidelinkR16 `lb:1,ub:maxBandCombSidelinkNRR16`
	InterbandpowersharingsyncdapsR16     *BandcombinationparametersV1630InterbandpowersharingsyncdapsR16
	InterbandpowersharingasyncdapsR16    *BandcombinationparametersV1630InterbandpowersharingasyncdapsR16
}
