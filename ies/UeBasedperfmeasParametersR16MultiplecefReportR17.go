package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-multipleCEF-Report-r17 ::= ENUMERATED
type UeBasedperfmeasParametersR16MultiplecefReportR17 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16MultiplecefReportR17EnumeratedNothing = iota
	UeBasedperfmeasParametersR16MultiplecefReportR17EnumeratedSupported
)
