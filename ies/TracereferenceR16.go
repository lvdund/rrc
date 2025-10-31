package ies

import "rrc/utils"

// TraceReference-r16 ::= SEQUENCE
type TracereferenceR16 struct {
	PlmnIdentityR16 PlmnIdentity
	TraceidR16      utils.OCTETSTRING `lb:3,ub:3`
}
