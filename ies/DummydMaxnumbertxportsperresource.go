package ies

import "rrc/utils"

// DummyD-maxNumberTxPortsPerResource ::= ENUMERATED
type DummydMaxnumbertxportsperresource struct {
	Value utils.ENUMERATED
}

const (
	DummydMaxnumbertxportsperresourceEnumeratedNothing = iota
	DummydMaxnumbertxportsperresourceEnumeratedP4
	DummydMaxnumbertxportsperresourceEnumeratedP8
	DummydMaxnumbertxportsperresourceEnumeratedP12
	DummydMaxnumbertxportsperresourceEnumeratedP16
	DummydMaxnumbertxportsperresourceEnumeratedP24
	DummydMaxnumbertxportsperresourceEnumeratedP32
)
