package ies

// SL-RelayUE-Config-r17 ::= SEQUENCE
type SlRelayueConfigR17 struct {
	ThreshhighrelayR17 *RsrpRange
	ThreshlowrelayR17  *RsrpRange
	HystmaxrelayR17    *Hysteresis
	HystminrelayR17    *Hysteresis
}
