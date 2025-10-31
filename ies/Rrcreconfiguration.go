package ies

import "rrc/utils"

// RRCReconfiguration-IEs ::= SEQUENCE
type Rrcreconfiguration struct {
	Radiobearerconfig        *Radiobearerconfig
	Secondarycellgroup       *utils.OCTETSTRING
	Measconfig               *Measconfig
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreconfigurationV1530
}
