package ies

import "rrc/utils"

// RRCResume-IEs ::= SEQUENCE
type Rrcresume struct {
	Radiobearerconfig        *Radiobearerconfig
	Mastercellgroup          *utils.OCTETSTRING
	Measconfig               *Measconfig
	Fullconfig               *utils.BOOLEAN
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcresumeV1560
}
