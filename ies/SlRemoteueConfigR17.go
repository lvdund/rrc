package ies

// SL-RemoteUE-Config-r17 ::= SEQUENCE
type SlRemoteueConfigR17 struct {
	ThreshhighremoteR17    *RsrpRange
	HystmaxremoteR17       *Hysteresis
	SlReselectionconfigR17 *SlReselectionconfigR17
}
