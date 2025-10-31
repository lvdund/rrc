package ies

// UE-BasedPerfMeas-Parameters-r16 ::= SEQUENCE
// Extensible
type UeBasedperfmeasParametersR16 struct {
	BarometermeasreportR16           *UeBasedperfmeasParametersR16BarometermeasreportR16
	ImmmeasbtR16                     *UeBasedperfmeasParametersR16ImmmeasbtR16
	ImmmeaswlanR16                   *UeBasedperfmeasParametersR16ImmmeaswlanR16
	LoggedmeasbtR16                  *UeBasedperfmeasParametersR16LoggedmeasbtR16
	LoggedmeasurementsR16            *UeBasedperfmeasParametersR16LoggedmeasurementsR16
	LoggedmeaswlanR16                *UeBasedperfmeasParametersR16LoggedmeaswlanR16
	OrientationmeasreportR16         *UeBasedperfmeasParametersR16OrientationmeasreportR16
	SpeedmeasreportR16               *UeBasedperfmeasParametersR16SpeedmeasreportR16
	GnssLocationR16                  *UeBasedperfmeasParametersR16GnssLocationR16
	UlpdcpDelayR16                   *UeBasedperfmeasParametersR16UlpdcpDelayR16
	SigbasedlogmdtOverrideprotectR17 *UeBasedperfmeasParametersR16SigbasedlogmdtOverrideprotectR17 `ext`
	MultiplecefReportR17             *UeBasedperfmeasParametersR16MultiplecefReportR17             `ext`
	ExcesspacketdelayR17             *UeBasedperfmeasParametersR16ExcesspacketdelayR17             `ext`
	EarlymeaslogR17                  *UeBasedperfmeasParametersR16EarlymeaslogR17                  `ext`
}
