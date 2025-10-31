package ies

import "rrc/utils"

// SystemInformationBlockType8 ::= SEQUENCE
// Extensible
type Systeminformationblocktype8 struct {
	Systemtimeinfo           *Systemtimeinfocdma2000
	Searchwindowsize         *utils.INTEGER `lb:0,ub:15`
	Parametershrpd           *Systeminformationblocktype8Parametershrpd
	Parameters1xrtt          *Systeminformationblocktype8Parameters1xrtt
	Latenoncriticalextension *utils.OCTETSTRING
}
