package ies

import "rrc/utils"

// MBMSCountingResponse-r10-IEs ::= SEQUENCE
type MbmscountingresponseR10 struct {
	MbsfnAreaindexR10        *utils.INTEGER `lb:0,ub:maxMBSFNArea1`
	CountingresponselistR10  *CountingresponselistR10
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbmscountingresponseR10IesNoncriticalextension
}
