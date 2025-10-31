package ies

// SystemInformationBlockType1-v1530-IEs ::= SEQUENCE
type Systeminformationblocktype1V1530 struct {
	HsdnCellR15                 *bool
	CellselectioninfoceV1530    *CellselectioninfoceV1530
	CrsIntfmitigconfigR15       *Systeminformationblocktype1V1530IesCrsIntfmitigconfigR15
	CellbarredCrsR15            Systeminformationblocktype1V1530IesCellbarredCrsR15
	PlmnIdentitylistV1530       *PlmnIdentitylistV1530
	PosschedulinginfolistR15    *PosschedulinginfolistR15
	Cellaccessrelatedinfo5gcR15 *Systeminformationblocktype1V1530IesCellaccessrelatedinfo5gcR15
	ImsEmergencysupport5gcR15   *bool
	EcalloverimsSupport5gcR15   *bool
	Noncriticalextension        *Systeminformationblocktype1V1540
}
