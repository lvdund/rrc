package ies

import "rrc/utils"

// PUR-Config-r16-pur-StartTimeParameters-r16 ::= SEQUENCE
type PurConfigR16PurStarttimeparametersR16 struct {
	PeriodicityandoffsetR16 PurPeriodicityandoffsetR16
	StartsfnR16             utils.INTEGER   `lb:0,ub:1023`
	StartsubframeR16        utils.INTEGER   `lb:0,ub:9`
	HsfnLsbInfoR16          utils.BITSTRING `lb:1,ub:1`
}
