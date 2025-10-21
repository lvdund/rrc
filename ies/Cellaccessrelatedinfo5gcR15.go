package ies

import "rrc/utils"

// CellAccessRelatedInfo-5GC-r15 ::= SEQUENCE
type Cellaccessrelatedinfo5gcR15 struct {
	PlmnIdentitylistR15    PlmnIdentitylistR15
	RanAreacodeR15         *RanAreacodeR15
	Trackingareacode5gcR15 Trackingareacode5gcR15
	Cellidentity5gcR15     Cellidentity5gcR15
}
