package ies

import "rrc/utils"

// RRCRelease-IEs ::= SEQUENCE
type Rrcrelease struct {
	Redirectedcarrierinfo     *Redirectedcarrierinfo
	Cellreselectionpriorities *Cellreselectionpriorities
	Suspendconfig             *Suspendconfig
	Deprioritisationreq       *RrcreleaseIesDeprioritisationreq
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *RrcreleaseV1540
}
