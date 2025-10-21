package ies

import "rrc/utils"

// SystemInformationBlockType4-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype4NbR13 struct {
	IntrafreqneighcelllistR13 *Intrafreqneighcelllist
	IntrafreqblackcelllistR13 *Intrafreqblackcelllist
	Latenoncriticalextension  *utils.OCTETSTRING
}
