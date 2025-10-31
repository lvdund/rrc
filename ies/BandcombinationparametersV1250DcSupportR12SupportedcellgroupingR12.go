package ies

import "rrc/utils"

// BandCombinationParameters-v1250-dc-Support-r12-supportedCellGrouping-r12 ::= CHOICE
const (
	BandcombinationparametersV1250DcSupportR12SupportedcellgroupingR12ChoiceNothing = iota
	BandcombinationparametersV1250DcSupportR12SupportedcellgroupingR12ChoiceThreeentriesR12
	BandcombinationparametersV1250DcSupportR12SupportedcellgroupingR12ChoiceFourentriesR12
	BandcombinationparametersV1250DcSupportR12SupportedcellgroupingR12ChoiceFiveentriesR12
)

type BandcombinationparametersV1250DcSupportR12SupportedcellgroupingR12 struct {
	Choice          uint64
	ThreeentriesR12 *utils.BITSTRING `lb:3,ub:3`
	FourentriesR12  *utils.BITSTRING `lb:7,ub:7`
	FiveentriesR12  *utils.BITSTRING `lb:15,ub:15`
}
