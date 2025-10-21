package ies

import "rrc/utils"

// RAN-AreaConfig-r15 ::= SEQUENCE
type RanAreaconfigR15 struct {
	Trackingareacode5gcR15 Trackingareacode5gcR15
	RanAreacodelistR15     *interface{}
}
