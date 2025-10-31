package ies

import "rrc/utils"

// DummyC-maxNumberTxPortsPerResource ::= ENUMERATED
type DummycMaxnumbertxportsperresource struct {
	Value utils.ENUMERATED
}

const (
	DummycMaxnumbertxportsperresourceEnumeratedNothing = iota
	DummycMaxnumbertxportsperresourceEnumeratedP8
	DummycMaxnumbertxportsperresourceEnumeratedP16
	DummycMaxnumbertxportsperresourceEnumeratedP32
)
