package ies

// AntennaInfoDedicated ::= SEQUENCE
type Antennainfodedicated struct {
	Transmissionmode           AntennainfodedicatedTransmissionmode
	Codebooksubsetrestriction  *AntennainfodedicatedCodebooksubsetrestriction
	UeTransmitantennaselection AntennainfodedicatedUeTransmitantennaselection
}
