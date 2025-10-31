package ies

// CC-Group-r17-offsetToDefault-r17 ::= CHOICE
const (
	CcGroupR17OffsettodefaultR17ChoiceNothing = iota
	CcGroupR17OffsettodefaultR17ChoiceOffsetvalue
	CcGroupR17OffsettodefaultR17ChoiceOffsetlist
)

type CcGroupR17OffsettodefaultR17 struct {
	Choice      uint64
	Offsetvalue *OffsetvalueR17
	Offsetlist  *[]OffsetvalueR17 `lb:1,ub:maxNrofReqComDCLocationR17`
}
