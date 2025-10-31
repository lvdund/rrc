package ies

import "rrc/utils"

// SPS-ConfigDL-STTI-r15-setup ::= SEQUENCE
// Extensible
type SpsConfigdlSttiR15Setup struct {
	SemipersistschedintervaldlSttiR15 SpsConfigdlSttiR15SetupSemipersistschedintervaldlSttiR15
	NumberofconfspsProcessesSttiR15   utils.INTEGER `lb:0,ub:12`
	TwoantennaportactivatedR15        *SpsConfigdlSttiR15SetupTwoantennaportactivatedR15
	SttiStarttimedlR15                utils.INTEGER `lb:0,ub:5`
	TpcPdcchConfigpucchSpsR15         *TpcPdcchConfig
}
