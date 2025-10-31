package ies

import "rrc/utils"

// FeatureSetUplink-v1610-ul-FullPwrMode2-TPMIGroup-r16 ::= SEQUENCE
type FeaturesetuplinkV1610UlFullpwrmode2TpmigroupR16 struct {
	TwoportsR16                 *utils.BITSTRING `lb:2,ub:2`
	FourportsnoncoherentR16     *FeaturesetuplinkV1610UlFullpwrmode2TpmigroupR16FourportsnoncoherentR16
	FourportspartialcoherentR16 *FeaturesetuplinkV1610UlFullpwrmode2TpmigroupR16FourportspartialcoherentR16
}
