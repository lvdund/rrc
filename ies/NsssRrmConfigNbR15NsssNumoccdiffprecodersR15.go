package ies

import "rrc/utils"

// NSSS-RRM-Config-NB-r15-nsss-NumOccDiffPrecoders-r15 ::= ENUMERATED
type NsssRrmConfigNbR15NsssNumoccdiffprecodersR15 struct {
	Value utils.ENUMERATED
}

const (
	NsssRrmConfigNbR15NsssNumoccdiffprecodersR15EnumeratedNothing = iota
	NsssRrmConfigNbR15NsssNumoccdiffprecodersR15EnumeratedN1
	NsssRrmConfigNbR15NsssNumoccdiffprecodersR15EnumeratedN2
	NsssRrmConfigNbR15NsssNumoccdiffprecodersR15EnumeratedN4
	NsssRrmConfigNbR15NsssNumoccdiffprecodersR15EnumeratedN8
)
