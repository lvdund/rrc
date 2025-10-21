package ies

import "rrc/utils"

// TraceReference-r10 ::= SEQUENCE
type TracereferenceR10 struct {
	PlmnIdentityR10 PlmnIdentity
	TraceidR10      utils.OCTETSTRING
}
