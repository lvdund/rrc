package ies

// DRX-Config-r13 ::= SEQUENCE
type DrxConfigR13 struct {
	OndurationtimerV1310        *DrxConfigR13OndurationtimerV1310
	DrxRetransmissiontimerV1310 *DrxConfigR13DrxRetransmissiontimerV1310
	DrxUlretransmissiontimerR13 *DrxConfigR13DrxUlretransmissiontimerR13
}
