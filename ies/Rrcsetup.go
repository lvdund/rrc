package ies

import "rrc/utils"

// RRCSetup-IEs ::= SEQUENCE
type Rrcsetup struct {
	Radiobearerconfig        Radiobearerconfig
	Mastercellgroup          utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcsetupV1700
}
