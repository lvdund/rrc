package ies

import "rrc/utils"

// SL-PSFCH-Config-r16 ::= SEQUENCE
// Extensible
type SlPsfchConfigR16 struct {
	SlPsfchPeriodR16                *SlPsfchConfigR16SlPsfchPeriodR16
	SlPsfchRbSetR16                 *utils.BITSTRING `lb:10,ub:275`
	SlNummuxcsPairR16               *SlPsfchConfigR16SlNummuxcsPairR16
	SlMintimegappsfchR16            *SlPsfchConfigR16SlMintimegappsfchR16
	SlPsfchHopidR16                 *utils.INTEGER `lb:0,ub:1023`
	SlPsfchCandidateresourcetypeR16 *SlPsfchConfigR16SlPsfchCandidateresourcetypeR16
}
