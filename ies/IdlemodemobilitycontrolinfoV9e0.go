package ies

// IdleModeMobilityControlInfo-v9e0 ::= SEQUENCE
type IdlemodemobilitycontrolinfoV9e0 struct {
	FreqprioritylisteutraV9e0 []FreqpriorityeutraV9e0 `lb:1,ub:maxFreq`
}
