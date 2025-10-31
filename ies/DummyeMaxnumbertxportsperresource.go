package ies

import "rrc/utils"

// DummyE-maxNumberTxPortsPerResource ::= ENUMERATED
type DummyeMaxnumbertxportsperresource struct {
	Value utils.ENUMERATED
}

const (
	DummyeMaxnumbertxportsperresourceEnumeratedNothing = iota
	DummyeMaxnumbertxportsperresourceEnumeratedP4
	DummyeMaxnumbertxportsperresourceEnumeratedP8
	DummyeMaxnumbertxportsperresourceEnumeratedP12
	DummyeMaxnumbertxportsperresourceEnumeratedP16
	DummyeMaxnumbertxportsperresourceEnumeratedP24
	DummyeMaxnumbertxportsperresourceEnumeratedP32
)
