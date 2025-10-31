package ies

import "rrc/utils"

// MulticastRLC-BearerConfig-r17 ::= SEQUENCE
type MulticastrlcBearerconfigR17 struct {
	ServedmbsRadiobearerR17 MrbIdentityR17
	IsptmEntityR17          *utils.BOOLEAN
}
