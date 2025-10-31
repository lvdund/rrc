package ies

import "rrc/utils"

// CA-ParametersNR-v1690-csi-ReportingCrossPUCCH-Grp-r16-sp-CSI-ReportingOnPUSCH-r16 ::= ENUMERATED
type CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpuschR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpuschR16EnumeratedNothing = iota
	CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpuschR16EnumeratedSupported
)
