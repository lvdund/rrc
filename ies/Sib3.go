package ies

import "rrc/utils"

// SIB3 ::= SEQUENCE
// Extensible
type Sib3 struct {
	Intrafreqneighcelllist        *Intrafreqneighcelllist
	Intrafreqexcludedcelllist     *Intrafreqexcludedcelllist
	Latenoncriticalextension      *utils.OCTETSTRING
	IntrafreqneighcelllistV1610   *IntrafreqneighcelllistV1610      `ext`
	IntrafreqallowedcelllistR16   *IntrafreqallowedcelllistR16      `ext`
	IntrafreqcagCelllistR16       *[]IntrafreqcagCelllistperplmnR16 `lb:1,ub:maxPLMN,ext`
	IntrafreqneighhsdnCelllistR17 *IntrafreqneighhsdnCelllistR17    `ext`
	IntrafreqneighcelllistV1710   *IntrafreqneighcelllistV1710      `ext`
	Channelaccessmode2R17         *Sib3Channelaccessmode2R17        `ext`
}
