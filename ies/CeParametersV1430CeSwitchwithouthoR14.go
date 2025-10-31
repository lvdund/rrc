package ies

import "rrc/utils"

// CE-Parameters-v1430-ce-SwitchWithoutHO-r14 ::= ENUMERATED
type CeParametersV1430CeSwitchwithouthoR14 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1430CeSwitchwithouthoR14EnumeratedNothing = iota
	CeParametersV1430CeSwitchwithouthoR14EnumeratedSupported
)
