package ies

import "rrc/utils"

// SystemInformationBlockType2-v9e0-IEs ::= SEQUENCE
type Systeminformationblocktype2V9e0Ies struct {
	UlCarrierfreqV9e0    *ArfcnValueeutraV9e0
	Noncriticalextension *Systeminformationblocktype2V9i0Ies
}
