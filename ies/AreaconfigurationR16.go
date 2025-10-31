package ies

// AreaConfiguration-r16 ::= SEQUENCE
type AreaconfigurationR16 struct {
	AreaconfigR16          AreaconfigR16
	InterfreqtargetlistR16 *[]InterfreqtargetinfoR16 `lb:1,ub:maxFreq`
}
