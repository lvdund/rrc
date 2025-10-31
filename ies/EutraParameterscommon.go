package ies

import "rrc/utils"

// EUTRA-ParametersCommon ::= SEQUENCE
// Extensible
type EutraParameterscommon struct {
	MfbiEutra                *EutraParameterscommonMfbiEutra
	ModifiedmprBehavioreutra *utils.BITSTRING `lb:32,ub:32`
	MultinsPmaxEutra         *EutraParameterscommonMultinsPmaxEutra
	RsSinrMeaseutra          *EutraParameterscommonRsSinrMeaseutra
	NeDc                     *EutraParameterscommonNeDc          `ext`
	NrHoToenDcR16            *EutraParameterscommonNrHoToenDcR16 `ext`
}
