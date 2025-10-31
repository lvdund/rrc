package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-aperiodicCsi-ReportingSTTI-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15AperiodiccsiReportingsttiR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15AperiodiccsiReportingsttiR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15AperiodiccsiReportingsttiR15EnumeratedSupported
)
