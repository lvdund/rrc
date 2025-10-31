package ies

import "rrc/utils"

// BandCombinationParameters-r13-dc-Support-r13-supportedCellGrouping-r13 ::= CHOICE
const (
	BandcombinationparametersR13DcSupportR13SupportedcellgroupingR13ChoiceNothing = iota
	BandcombinationparametersR13DcSupportR13SupportedcellgroupingR13ChoiceThreeentriesR13
	BandcombinationparametersR13DcSupportR13SupportedcellgroupingR13ChoiceFourentriesR13
	BandcombinationparametersR13DcSupportR13SupportedcellgroupingR13ChoiceFiveentriesR13
)

type BandcombinationparametersR13DcSupportR13SupportedcellgroupingR13 struct {
	Choice          uint64
	ThreeentriesR13 *utils.BITSTRING `lb:3,ub:3`
	FourentriesR13  *utils.BITSTRING `lb:7,ub:7`
	FiveentriesR13  *utils.BITSTRING `lb:15,ub:15`
}
