package ies

// CE-PDSCH-MultiTB-Config-r16 ::= SEQUENCE
type CePdschMultitbConfigR16 struct {
	InterleavingR16    *CePdschMultitbConfigR16InterleavingR16
	HarqAckbundlingR16 *CePdschMultitbConfigR16HarqAckbundlingR16
}
