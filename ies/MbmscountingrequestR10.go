package ies

import "rrc/utils"

// MBMSCountingRequest-r10 ::= SEQUENCE
type MbmscountingrequestR10 struct {
	CountingrequestlistR10   CountingrequestlistR10
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
