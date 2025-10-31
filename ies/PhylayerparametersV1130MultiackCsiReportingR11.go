package ies

import "rrc/utils"

// PhyLayerParameters-v1130-multiACK-CSI-Reporting-r11 ::= ENUMERATED
type PhylayerparametersV1130MultiackCsiReportingR11 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1130MultiackCsiReportingR11EnumeratedNothing = iota
	PhylayerparametersV1130MultiackCsiReportingR11EnumeratedSupported
)
