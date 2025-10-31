package ies

// AreaConfiguration-v1700 ::= SEQUENCE
type AreaconfigurationV1700 struct {
	AreaconfigR17          *AreaconfigR16
	InterfreqtargetlistR17 *[]InterfreqtargetinfoR16 `lb:1,ub:maxFreq`
}
