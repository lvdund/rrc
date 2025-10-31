package ies

import "rrc/utils"

// DummyG-maxNumberSSB-CSI-RS-ResourceTwoTx ::= ENUMERATED
type DummygMaxnumberssbCsiRsResourcetwotx struct {
	Value utils.ENUMERATED
}

const (
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedNothing = iota
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN0
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN4
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN8
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN16
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN32
	DummygMaxnumberssbCsiRsResourcetwotxEnumeratedN64
)
