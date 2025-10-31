package ies

import "rrc/utils"

// AntennaInfoDedicatedSTTI-r15-setup-codebookSubsetRestriction ::= CHOICE
const (
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceNothing = iota
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm3R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm3R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm4R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm4R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm5R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm5R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm6R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm6R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm8R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm8R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN2txantennaTm9and10R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN4txantennaTm9and10R15
	AntennainfodedicatedsttiR15SetupCodebooksubsetrestrictionChoiceN8txantennaTm9and10R15
)

type AntennainfodedicatedsttiR15SetupCodebooksubsetrestriction struct {
	Choice                 uint64
	N2txantennaTm3R15      *utils.BITSTRING `lb:2,ub:2`
	N4txantennaTm3R15      *utils.BITSTRING `lb:4,ub:4`
	N2txantennaTm4R15      *utils.BITSTRING `lb:6,ub:6`
	N4txantennaTm4R15      *utils.BITSTRING `lb:64,ub:64`
	N2txantennaTm5R15      *utils.BITSTRING `lb:4,ub:4`
	N4txantennaTm5R15      *utils.BITSTRING `lb:16,ub:16`
	N2txantennaTm6R15      *utils.BITSTRING `lb:4,ub:4`
	N4txantennaTm6R15      *utils.BITSTRING `lb:16,ub:16`
	N2txantennaTm8R15      *utils.BITSTRING `lb:6,ub:6`
	N4txantennaTm8R15      *utils.BITSTRING `lb:64,ub:64`
	N2txantennaTm9and10R15 *utils.BITSTRING `lb:6,ub:6`
	N4txantennaTm9and10R15 *utils.BITSTRING `lb:96,ub:96`
	N8txantennaTm9and10R15 *utils.BITSTRING `lb:109,ub:109`
}
