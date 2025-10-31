package ies

// MRDC-Parameters-v1620 ::= SEQUENCE
type MrdcParametersV1620 struct {
	MaxuplinkdutycycleInterbandendcTddPc2R16 *MrdcParametersV1620MaxuplinkdutycycleInterbandendcTddPc2R16
	TdmRestrictiontddEndcR16                 *MrdcParametersV1620TdmRestrictiontddEndcR16
	TdmRestrictionfddEndcR16                 *MrdcParametersV1620TdmRestrictionfddEndcR16
	SingleulHarqOffsettddPcellR16            *MrdcParametersV1620SingleulHarqOffsettddPcellR16
	TdmRestrictiondualtxFddEndcR16           *MrdcParametersV1620TdmRestrictiondualtxFddEndcR16
}
