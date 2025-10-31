package ies

import "rrc/utils"

// MBSFNAreaConfiguration-v930-IEs ::= SEQUENCE
type MbsfnareaconfigurationV930 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbsfnareaconfigurationV1250
}
