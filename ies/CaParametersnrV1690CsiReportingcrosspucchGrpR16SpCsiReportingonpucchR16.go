package ies

import "rrc/utils"

// CA-ParametersNR-v1690-csi-ReportingCrossPUCCH-Grp-r16-sp-CSI-ReportingOnPUCCH-r16 ::= ENUMERATED
type CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpucchR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpucchR16EnumeratedNothing = iota
	CaParametersnrV1690CsiReportingcrosspucchGrpR16SpCsiReportingonpucchR16EnumeratedSupported
)
