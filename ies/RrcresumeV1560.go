package ies

import "rrc/utils"

// RRCResume-v1560-IEs ::= SEQUENCE
type RrcresumeV1560 struct {
	Radiobearerconfig2   *utils.OCTETSTRING
	SkCounter            *SkCounter
	Noncriticalextension *RrcresumeV1610
}
