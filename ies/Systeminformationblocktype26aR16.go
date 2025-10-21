package ies

import "rrc/utils"

// SystemInformationBlockType26a-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype26aR16 struct {
	PlmnInfolistR16          PlmnInfolistR16
	BandlistendcR16          BandlistendcR16
	Latenoncriticalextension *utils.OCTETSTRING
}
