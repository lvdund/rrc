package ies

import "rrc/utils"

// AntennaInfoDedicated-v920-codebookSubsetRestriction-v920 ::= CHOICE
const (
	AntennainfodedicatedV920CodebooksubsetrestrictionV920ChoiceNothing = iota
	AntennainfodedicatedV920CodebooksubsetrestrictionV920ChoiceN2txantennaTm8R9
	AntennainfodedicatedV920CodebooksubsetrestrictionV920ChoiceN4txantennaTm8R9
)

type AntennainfodedicatedV920CodebooksubsetrestrictionV920 struct {
	Choice           uint64
	N2txantennaTm8R9 *utils.BITSTRING `lb:6,ub:6`
	N4txantennaTm8R9 *utils.BITSTRING `lb:32,ub:32`
}
