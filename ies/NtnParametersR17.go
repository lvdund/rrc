package ies

// NTN-Parameters-r17 ::= SEQUENCE
type NtnParametersR17 struct {
	InactivestatentnR17             *NtnParametersR17InactivestatentnR17
	RaSdtNtnR17                     *NtnParametersR17RaSdtNtnR17
	SrbSdtNtnR17                    *NtnParametersR17SrbSdtNtnR17
	MeasandmobparametersntnR17      *Measandmobparameters
	MacParametersntnR17             *MacParameters
	PhyParametersntnR17             *PhyParameters
	FddAddUeNrCapabilitiesntnR17    *UeNrCapabilityaddxddMode
	Fr1AddUeNrCapabilitiesntnR17    *UeNrCapabilityaddfrxMode
	UeBasedperfmeasParametersntnR17 *UeBasedperfmeasParametersR16
	SonParametersntnR17             *SonParametersR16
}
