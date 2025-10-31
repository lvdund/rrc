package ies

import "rrc/utils"

// AntennaInfoDedicatedSTTI-r15-setup ::= SEQUENCE
type AntennainfodedicatedsttiR15Setup struct {
	TransmissionmodedlMbsfnR15     *AntennainfodedicatedsttiR15SetupTransmissionmodedlMbsfnR15
	TransmissionmodedlNonmbsfnR15  *AntennainfodedicatedsttiR15SetupTransmissionmodedlNonmbsfnR15
	Codebooksubsetrestriction      *AntennainfodedicatedsttiR15SetupCodebooksubsetrestriction
	MaxlayersmimoSttiR15           *AntennainfodedicatedsttiR15SetupMaxlayersmimoSttiR15
	SlotsubslotpdschTxdiv2layerR15 utils.BOOLEAN
	SlotsubslotpdschTxdiv4layerR15 utils.BOOLEAN
}
