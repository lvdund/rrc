package ies

// SL-DiscConfigOtherInterFreq-r13 ::= SEQUENCE
type SlDiscconfigotherinterfreqR13 struct {
	TxpowerinfoR13           *SlDisctxpowerinfolistR12
	RefcarriercommonR13      *SlDiscconfigotherinterfreqR13RefcarriercommonR13
	DiscsyncconfigR13        *SlSyncconfiglistnfreqR13
	DisccellselectioninfoR13 *CellselectioninfonfreqR13
}
