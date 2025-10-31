package ies

import "rrc/utils"

// MBMS-Parameters-v1470-mbms-MaxBW-r14 ::= CHOICE
const (
	MbmsParametersV1470MbmsMaxbwR14ChoiceNothing = iota
	MbmsParametersV1470MbmsMaxbwR14ChoiceImplicitvalue
	MbmsParametersV1470MbmsMaxbwR14ChoiceExplicitvalue
)

type MbmsParametersV1470MbmsMaxbwR14 struct {
	Choice        uint64
	Implicitvalue *struct{}
	Explicitvalue *utils.INTEGER `lb:2,ub:20`
}
