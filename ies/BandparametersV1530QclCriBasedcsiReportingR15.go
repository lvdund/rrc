package ies

import "rrc/utils"

// BandParameters-v1530-qcl-CRI-BasedCSI-Reporting-r15 ::= ENUMERATED
type BandparametersV1530QclCriBasedcsiReportingR15 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1530QclCriBasedcsiReportingR15EnumeratedNothing = iota
	BandparametersV1530QclCriBasedcsiReportingR15EnumeratedSupported
)
