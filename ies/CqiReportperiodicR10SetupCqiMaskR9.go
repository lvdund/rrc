package ies

import "rrc/utils"

// CQI-ReportPeriodic-r10-setup-cqi-Mask-r9 ::= ENUMERATED
type CqiReportperiodicR10SetupCqiMaskR9 struct {
	Value utils.ENUMERATED
}

const (
	CqiReportperiodicR10SetupCqiMaskR9EnumeratedNothing = iota
	CqiReportperiodicR10SetupCqiMaskR9EnumeratedSetup
)
