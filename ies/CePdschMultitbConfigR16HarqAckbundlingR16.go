package ies

import "rrc/utils"

// CE-PDSCH-MultiTB-Config-r16-harq-AckBundling-r16 ::= ENUMERATED
type CePdschMultitbConfigR16HarqAckbundlingR16 struct {
	Value utils.ENUMERATED
}

const (
	CePdschMultitbConfigR16HarqAckbundlingR16EnumeratedNothing = iota
	CePdschMultitbConfigR16HarqAckbundlingR16EnumeratedOn
)
