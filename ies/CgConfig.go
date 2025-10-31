package ies

import "rrc/utils"

// CG-Config-IEs ::= SEQUENCE
type CgConfig struct {
	ScgCellgroupconfig         *utils.OCTETSTRING
	ScgRbConfig                *utils.OCTETSTRING
	Configrestrictmodreq       *Configrestrictmodreqscg
	DrxInfoscg                 *DrxInfo
	Candidatecellinfolistsn    *utils.OCTETSTRING
	Measconfigsn               *Measconfigsn
	Selectedbandcombination    *Bandcombinationinfosn
	FrInfolistscg              *FrInfolist
	Candidateservingfreqlistnr *Candidateservingfreqlistnr
	Noncriticalextension       *CgConfigV1540
}
