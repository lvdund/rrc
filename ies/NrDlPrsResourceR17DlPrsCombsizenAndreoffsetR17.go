package ies

import "rrc/utils"

// NR-DL-PRS-Resource-r17-dl-PRS-CombSizeN-AndReOffset-r17 ::= CHOICE
// Extensible
const (
	NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17ChoiceNothing = iota
	NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17ChoiceN2R17
	NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17ChoiceN4R17
	NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17ChoiceN6R17
	NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17ChoiceN12R17
)

type NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17 struct {
	Choice uint64
	N2R17  *utils.INTEGER `lb:0,ub:1`
	N4R17  *utils.INTEGER `lb:0,ub:3`
	N6R17  *utils.INTEGER `lb:0,ub:5`
	N12R17 *utils.INTEGER `lb:0,ub:11`
}
