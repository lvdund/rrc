package ies

import "rrc/utils"

// RRCReestablishment-IEs ::= SEQUENCE
type Rrcreestablishment struct {
	Nexthopchainingcount     Nexthopchainingcount
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreestablishmentV1700
}
