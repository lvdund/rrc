package ies

import "rrc/utils"

// UDT-Restricting-r13 ::= SEQUENCE
type UdtRestrictingR13 struct {
	UdtRestrictingR13     *utils.ENUMERATED
	UdtRestrictingtimeR13 *utils.ENUMERATED
}
