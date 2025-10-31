package ies

import "rrc/utils"

// Phy-ParametersCommon-guardSymbolReportReception-IAB-r17 ::= ENUMERATED
type PhyParameterscommonGuardsymbolreportreceptionIabR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonGuardsymbolreportreceptionIabR17EnumeratedNothing = iota
	PhyParameterscommonGuardsymbolreportreceptionIabR17EnumeratedSupported
)
