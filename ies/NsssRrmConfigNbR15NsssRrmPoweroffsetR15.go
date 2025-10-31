package ies

import "rrc/utils"

// NSSS-RRM-Config-NB-r15-nsss-RRM-PowerOffset-r15 ::= ENUMERATED
type NsssRrmConfigNbR15NsssRrmPoweroffsetR15 struct {
	Value utils.ENUMERATED
}

const (
	NsssRrmConfigNbR15NsssRrmPoweroffsetR15EnumeratedNothing = iota
	NsssRrmConfigNbR15NsssRrmPoweroffsetR15EnumeratedDb_3
	NsssRrmConfigNbR15NsssRrmPoweroffsetR15EnumeratedDb0
	NsssRrmConfigNbR15NsssRrmPoweroffsetR15EnumeratedDb3
)
