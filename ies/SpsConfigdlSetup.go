package ies

import "rrc/utils"

// SPS-ConfigDL-setup ::= SEQUENCE
// Extensible
type SpsConfigdlSetup struct {
	Semipersistschedintervaldl SpsConfigdlSetupSemipersistschedintervaldl
	NumberofconfspsProcesses   utils.INTEGER `lb:0,ub:8`
	N1pucchAnPersistentlist    N1pucchAnPersistentlist
}
