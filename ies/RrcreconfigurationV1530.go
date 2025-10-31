package ies

import "rrc/utils"

// RRCReconfiguration-v1530-IEs ::= SEQUENCE
type RrcreconfigurationV1530 struct {
	Mastercellgroup                    *utils.OCTETSTRING
	Fullconfig                         *utils.BOOLEAN
	DedicatednasMessagelist            *[]DedicatednasMessage `lb:1,ub:maxDRB`
	Masterkeyupdate                    *Masterkeyupdate
	Dedicatedsib1Delivery              *utils.OCTETSTRING
	Dedicatedsysteminformationdelivery *utils.OCTETSTRING
	Otherconfig                        *Otherconfig
	Noncriticalextension               *RrcreconfigurationV1540
}
