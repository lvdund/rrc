package ies

import "rrc/utils"

// UE-NR-Capability-v1700 ::= SEQUENCE
type UeNrCapabilityV1700 struct {
	InactivestatepoDeterminationR17     *UeNrCapabilityV1700InactivestatepoDeterminationR17
	HighspeedparametersV1700            *HighspeedparametersV1700
	PowsavParametersV1700               *PowsavParametersV1700
	MacParametersV1700                  *MacParametersV1700
	ImsParametersV1700                  *ImsParametersV1700
	MeasandmobparametersV1700           MeasandmobparametersV1700
	ApplayermeasparametersR17           *ApplayermeasparametersR17
	RedcapparametersR17                 *RedcapparametersR17
	RaSdtR17                            *UeNrCapabilityV1700RaSdtR17
	SrbSdtR17                           *UeNrCapabilityV1700SrbSdtR17
	GnbSiderttBasedpdcR17               *UeNrCapabilityV1700GnbSiderttBasedpdcR17
	BhRlfDetectionrecoveryIndicationR17 *UeNrCapabilityV1700BhRlfDetectionrecoveryIndicationR17
	NrdcParametersV1700                 *NrdcParametersV1700
	BapParametersV1700                  *BapParametersV1700
	MusimGappreferenceR17               *UeNrCapabilityV1700MusimGappreferenceR17
	MusimleaveconnectedR17              *UeNrCapabilityV1700MusimleaveconnectedR17
	MbsParametersR17                    MbsParametersR17
	NonterrestrialnetworkR17            *UeNrCapabilityV1700NonterrestrialnetworkR17
	NtnScenariosupportR17               *UeNrCapabilityV1700NtnScenariosupportR17
	SliceinfoforcellreselectionR17      *UeNrCapabilityV1700SliceinfoforcellreselectionR17
	UeRadiopaginginfoR17                *UeRadiopaginginfoR17
	UlGapfr2PatternR17                  *utils.BITSTRING `lb:4,ub:4`
	NtnParametersR17                    *NtnParametersR17
	Noncriticalextension                *UeNrCapabilityV1700Noncriticalextension
}
