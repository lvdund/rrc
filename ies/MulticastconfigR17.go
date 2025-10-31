package ies

import "rrc/utils"

// MulticastConfig-r17 ::= SEQUENCE
type MulticastconfigR17 struct {
	PdschHarqAckCodebooklistmulticastR17 *utils.Setuprelease[PdschHarqAckCodebooklistR16]
	Type1CodebookGenerationmodeR17       *MulticastconfigR17Type1CodebookGenerationmodeR17
}
