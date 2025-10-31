package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-v1700 ::= SEQUENCE
type InterfreqcarrierfreqinfoV1700 struct {
	InterfreqneighhsdnCelllistR17 *InterfreqneighhsdnCelllistR17
	HighspeedmeasinterfreqR17     *utils.BOOLEAN
	RedcapaccessallowedR17        *utils.BOOLEAN
	SsbPositionqclCommonR17       *SsbPositionqclRelationR17
	InterfreqneighcelllistV1710   *InterfreqneighcelllistV1710
}
