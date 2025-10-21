package ies

import "rrc/utils"

// MBSFNAreaConfiguration-v930-IEs ::= SEQUENCE
type MbsfnareaconfigurationV930Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbsfnareaconfigurationV1250Ies
}
