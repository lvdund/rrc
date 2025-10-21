package ies

import "rrc/utils"

// SystemInformationBlockType1-v1530-IEs ::= SEQUENCE
type Systeminformationblocktype1V1530Ies struct {
	HsdnCellR15                 *utils.ENUMERATED
	CellselectioninfoceV1530    *CellselectioninfoceV1530
	CrsIntfmitigconfigR15       *interface{}
	CellbarredCrsR15            utils.ENUMERATED
	PlmnIdentitylistV1530       *PlmnIdentitylistV1530
	PosschedulinginfolistR15    *PosschedulinginfolistR15
	Cellaccessrelatedinfo5gcR15 *interface{}
	ImsEmergencysupport5gcR15   *utils.ENUMERATED
	EcalloverimsSupport5gcR15   *utils.ENUMERATED
	Noncriticalextension        *Systeminformationblocktype1V1540Ies
}
