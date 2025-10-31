package ies

import "rrc/utils"

// CG-Config-v1560-IEs ::= SEQUENCE
type CgConfigV1560 struct {
	Pscellfrequencyeutra          *ArfcnValueeutra
	ScgCellgroupconfigeutra       *utils.OCTETSTRING
	CandidatecellinfolistsnEutra  *utils.OCTETSTRING
	Candidateservingfreqlisteutra *Candidateservingfreqlisteutra
	Needforgaps                   *utils.BOOLEAN
	DrxConfigscg                  *DrxConfig
	ReportcgiRequesteutra         *CgConfigV1560IesReportcgiRequesteutra
	Noncriticalextension          *CgConfigV1590
}
