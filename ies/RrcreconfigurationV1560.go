package ies

import "rrc/utils"

// RRCReconfiguration-v1560-IEs ::= SEQUENCE
type RrcreconfigurationV1560 struct {
	MrdcSecondarycellgroupconfig *utils.Setuprelease[MrdcSecondarycellgroupconfig]
	Radiobearerconfig2           *utils.OCTETSTRING
	SkCounter                    *SkCounter
	Noncriticalextension         *RrcreconfigurationV1610
}
