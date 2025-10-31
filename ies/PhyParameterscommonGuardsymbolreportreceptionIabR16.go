package ies

import "rrc/utils"

// Phy-ParametersCommon-guardSymbolReportReception-IAB-r16 ::= ENUMERATED
type PhyParameterscommonGuardsymbolreportreceptionIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonGuardsymbolreportreceptionIabR16EnumeratedNothing = iota
	PhyParameterscommonGuardsymbolreportreceptionIabR16EnumeratedSupported
)
