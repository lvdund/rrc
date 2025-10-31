package ies

import "rrc/utils"

// BeamFailureRecoveryConfig ::= SEQUENCE
// Extensible
type Beamfailurerecoveryconfig struct {
	RootsequenceindexBfr        *utils.INTEGER `lb:0,ub:137`
	RachConfigbfr               *RachConfiggeneric
	RsrpThresholdssb            *RsrpRange
	Candidatebeamrslist         *[]PrachResourcededicatedbfr `lb:1,ub:maxNrofCandidateBeams`
	SsbPerrachOccasion          *BeamfailurerecoveryconfigSsbPerrachOccasion
	RaSsbOccasionmaskindex      *utils.INTEGER `lb:0,ub:15`
	Recoverysearchspaceid       *Searchspaceid
	RaPrioritization            *RaPrioritization
	Beamfailurerecoverytimer    *BeamfailurerecoveryconfigBeamfailurerecoverytimer
	Msg1Subcarrierspacing       *Subcarrierspacing                             `ext`
	RaPrioritizationtwostepR16  *RaPrioritization                              `ext`
	CandidatebeamrslistextV1610 *utils.Setuprelease[CandidatebeamrslistextR16] `ext`
	SpcellBfrCbraR16            *utils.BOOLEAN                                 `ext`
}
