package ies

import "rrc/utils"

// SIB1 ::= SEQUENCE
type Sib1 struct {
	Cellselectioninfo        *Sib1Cellselectioninfo
	Cellaccessrelatedinfo    Cellaccessrelatedinfo
	Connestfailurecontrol    *Connestfailurecontrol
	SiSchedulinginfo         *SiSchedulinginfo
	Servingcellconfigcommon  *Servingcellconfigcommonsib
	ImsEmergencysupport      *utils.BOOLEAN
	EcalloverimsSupport      *utils.BOOLEAN
	UeTimersandconstants     *UeTimersandconstants
	UacBarringinfo           *Sib1UacBarringinfo
	Usefullresumeid          *utils.BOOLEAN
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Sib1V1610
}
