package ies

import "rrc/utils"

// CG-ConfigInfo-IEs ::= SEQUENCE
type CgConfiginfo struct {
	UeCapabilityinfo         *utils.OCTETSTRING
	Candidatecellinfolistmn  *Measresultlist2nr
	Candidatecellinfolistsn  *utils.OCTETSTRING
	MeasresultcelllistsftdNr *MeasresultcelllistsftdNr
	Scgfailureinfo           *CgConfiginfoIesScgfailureinfo
	Configrestrictinfo       *Configrestrictinfoscg
	DrxInfomcg               *DrxInfo
	Measconfigmn             *Measconfigmn
	Sourceconfigscg          *utils.OCTETSTRING
	ScgRbConfig              *utils.OCTETSTRING
	McgRbConfig              *utils.OCTETSTRING
	MrdcAssistanceinfo       *MrdcAssistanceinfo
	Noncriticalextension     *CgConfiginfoV1540
}
