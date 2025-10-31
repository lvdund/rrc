package ies

import "rrc/utils"

// CG-ConfigInfo-v1560-IEs ::= SEQUENCE
type CgConfiginfoV1560 struct {
	CandidatecellinfolistmnEutra *utils.OCTETSTRING
	CandidatecellinfolistsnEutra *utils.OCTETSTRING
	SourceconfigscgEutra         *utils.OCTETSTRING
	Scgfailureinfoeutra          *CgConfiginfoV1560IesScgfailureinfoeutra
	DrxConfigmcg                 *DrxConfig
	MeasresultreportcgiEutra     *CgConfiginfoV1560IesMeasresultreportcgiEutra
	MeasresultcelllistsftdEutra  *MeasresultcelllistsftdEutra
	FrInfolistmcg                *FrInfolist
	Noncriticalextension         *CgConfiginfoV1570
}
