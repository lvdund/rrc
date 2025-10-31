package ies

// SL-DiscConfigRemoteUE-r13 ::= SEQUENCE
type SlDiscconfigremoteueR13 struct {
	ThreshhighR13        *RsrpRangesl4R13
	HystmaxR13           *SlDiscconfigremoteueR13HystmaxR13
	ReselectioninfoicR13 ReselectioninforelayR13
}
