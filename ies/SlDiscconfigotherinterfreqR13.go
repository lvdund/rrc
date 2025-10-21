package ies

import "rrc/utils"

// SL-DiscConfigOtherInterFreq-r13 ::= SEQUENCE
type SlDiscconfigotherinterfreqR13 struct {
	TxpowerinfoR13           *SlDisctxpowerinfolistR12
	RefcarriercommonR13      *utils.ENUMERATED
	DiscsyncconfigR13        *SlSyncconfiglistnfreqR13
	DisccellselectioninfoR13 *CellselectioninfonfreqR13
}
