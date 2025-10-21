package ies

import "rrc/utils"

// SPS-ConfigUL-ToAddModList-r14 ::= SEQUENCE OF SPS-ConfigUL
// SIZE (1..maxConfigSPS-r14)
type SpsConfigulToaddmodlistR14 struct {
	Value utils.Sequence[SpsConfigul]
}
