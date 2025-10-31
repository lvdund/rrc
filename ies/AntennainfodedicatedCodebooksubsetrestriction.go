package ies

import "rrc/utils"

// AntennaInfoDedicated-codebookSubsetRestriction ::= CHOICE
const (
	AntennainfodedicatedCodebooksubsetrestrictionChoiceNothing = iota
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN2txantennaTm3
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN4txantennaTm3
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN2txantennaTm4
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN4txantennaTm4
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN2txantennaTm5
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN4txantennaTm5
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN2txantennaTm6
	AntennainfodedicatedCodebooksubsetrestrictionChoiceN4txantennaTm6
)

type AntennainfodedicatedCodebooksubsetrestriction struct {
	Choice         uint64
	N2txantennaTm3 *utils.BITSTRING `lb:2,ub:2`
	N4txantennaTm3 *utils.BITSTRING `lb:4,ub:4`
	N2txantennaTm4 *utils.BITSTRING `lb:6,ub:6`
	N4txantennaTm4 *utils.BITSTRING `lb:64,ub:64`
	N2txantennaTm5 *utils.BITSTRING `lb:4,ub:4`
	N4txantennaTm5 *utils.BITSTRING `lb:16,ub:16`
	N2txantennaTm6 *utils.BITSTRING `lb:4,ub:4`
	N4txantennaTm6 *utils.BITSTRING `lb:16,ub:16`
}
