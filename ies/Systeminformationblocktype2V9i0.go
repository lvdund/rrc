package ies

import "rrc/utils"

// SystemInformationBlockType2-v9i0-IEs ::= SEQUENCE
type Systeminformationblocktype2V9i0 struct {
	Noncriticalextension *utils.OCTETSTRING
	Dummy                *Systeminformationblocktype2V9i0IesDummy
}
