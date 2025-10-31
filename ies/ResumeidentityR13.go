package ies

import "rrc/utils"

// ResumeIdentity-r13 ::= BIT STRING (SIZE (40))
type ResumeidentityR13 struct {
	Value utils.BITSTRING `lb:40,ub:40`
}
