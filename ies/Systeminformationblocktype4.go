package ies

import "rrc/utils"

// SystemInformationBlockType4 ::= SEQUENCE
// Extensible
type Systeminformationblocktype4 struct {
	Intrafreqneighcelllist   *Intrafreqneighcelllist
	Intrafreqblackcelllist   *Intrafreqblackcelllist
	CsgPhyscellidrange       *Physcellidrange
	Latenoncriticalextension *utils.OCTETSTRING
}
