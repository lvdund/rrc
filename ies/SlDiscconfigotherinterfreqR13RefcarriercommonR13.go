package ies

import "rrc/utils"

// SL-DiscConfigOtherInterFreq-r13-refCarrierCommon-r13 ::= ENUMERATED
type SlDiscconfigotherinterfreqR13RefcarriercommonR13 struct {
	Value utils.ENUMERATED
}

const (
	SlDiscconfigotherinterfreqR13RefcarriercommonR13EnumeratedNothing = iota
	SlDiscconfigotherinterfreqR13RefcarriercommonR13EnumeratedPcell
)
