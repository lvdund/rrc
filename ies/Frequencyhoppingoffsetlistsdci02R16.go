package ies

import "rrc/utils"

// FrequencyHoppingOffsetListsDCI-0-2-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type Frequencyhoppingoffsetlistsdci02R16 struct {
	Value []utils.INTEGER `lb:1,ub:4`
}
