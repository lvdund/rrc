package ies

import "rrc/utils"

// SystemInformationBlockType8 ::= SEQUENCE
// Extensible
type Systeminformationblocktype8 struct {
	Systemtimeinfo           *Systemtimeinfocdma2000
	Searchwindowsize         *utils.INTEGER
	Parametershrpd           *interface{}
	Parameters1xrtt          *interface{}
	Latenoncriticalextension *utils.OCTETSTRING
}
