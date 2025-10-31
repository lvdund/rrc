package ies

import "rrc/utils"

// PerRASSBInfo-r16 ::= SEQUENCE
type PerrassbinfoR16 struct {
	SsbIndexR16                   SsbIndex
	NumberofpreamblessentonssbR16 utils.INTEGER `lb:0,ub:200`
	PerraattemptinfolistR16       PerraattemptinfolistR16
}
