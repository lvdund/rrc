package ies

// SL-DiscConfigRelayUE-r13 ::= SEQUENCE
type SlDiscconfigrelayueR13 struct {
	ThreshhighR13 *RsrpRangesl4R13
	ThreshlowR13  *RsrpRangesl4R13
	HystmaxR13    *SlDiscconfigrelayueR13HystmaxR13
	HystminR13    *SlDiscconfigrelayueR13HystminR13
}
