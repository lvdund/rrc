package ies

import "rrc/utils"

// AdditionalSI-Info-r9-csg-MemberStatus-r9 ::= ENUMERATED
type AdditionalsiInfoR9CsgMemberstatusR9 struct {
	Value utils.ENUMERATED
}

const (
	AdditionalsiInfoR9CsgMemberstatusR9EnumeratedNothing = iota
	AdditionalsiInfoR9CsgMemberstatusR9EnumeratedMember
)
