package ies

import "rrc/utils"

// SL-DiscConfigRemoteUE-r13 ::= SEQUENCE
type SlDiscconfigremoteueR13 struct {
	ThreshhighR13        *RsrpRangesl4R13
	HystmaxR13           *utils.ENUMERATED
	ReselectioninfoicR13 ReselectioninforelayR13
}
