package ies

import "rrc/utils"

// RRCReject-IEs ::= SEQUENCE
type Rrcreject struct {
	Waittime                 *Rejectwaittime
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcrejectIesNoncriticalextension
}
