package ies

import "rrc/utils"

// DC-Parameters-r12 ::= SEQUENCE
type DcParametersR12 struct {
	DrbTypesplitR12 *utils.ENUMERATED
	DrbTypescgR12   *utils.ENUMERATED
}
