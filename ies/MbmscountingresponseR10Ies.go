package ies

import "rrc/utils"

// MBMSCountingResponse-r10-IEs ::= SEQUENCE
type MbmscountingresponseR10Ies struct {
	MbsfnAreaindexR10        *utils.INTEGER
	CountingresponselistR10  *CountingresponselistR10
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
