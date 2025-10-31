package ies

import "rrc/utils"

// DummyG-maxNumberSSB-CSI-RS-ResourceOneTx ::= ENUMERATED
type DummygMaxnumberssbCsiRsResourceonetx struct {
	Value utils.ENUMERATED
}

const (
	DummygMaxnumberssbCsiRsResourceonetxEnumeratedNothing = iota
	DummygMaxnumberssbCsiRsResourceonetxEnumeratedN8
	DummygMaxnumberssbCsiRsResourceonetxEnumeratedN16
	DummygMaxnumberssbCsiRsResourceonetxEnumeratedN32
	DummygMaxnumberssbCsiRsResourceonetxEnumeratedN64
)
