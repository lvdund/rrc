package ies

import "rrc/utils"

// SystemInformationBlockType5-v13a0-IEs ::= SEQUENCE
type Systeminformationblocktype5V13a0Ies struct {
	Latenoncriticalextension      *utils.OCTETSTRING
	InterfreqcarrierfreqlistV13a0 *InterfreqcarrierfreqlistV13a0
	Noncriticalextension          *interface{}
}
