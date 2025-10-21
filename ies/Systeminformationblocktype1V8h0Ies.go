package ies

import "rrc/utils"

// SystemInformationBlockType1-v8h0-IEs ::= SEQUENCE
type Systeminformationblocktype1V8h0Ies struct {
	Multibandinfolist    *Multibandinfolist
	Noncriticalextension *Systeminformationblocktype1V9e0Ies
}
