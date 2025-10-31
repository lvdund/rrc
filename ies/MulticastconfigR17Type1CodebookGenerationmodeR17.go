package ies

import "rrc/utils"

// MulticastConfig-r17-type1-Codebook-GenerationMode-r17 ::= ENUMERATED
type MulticastconfigR17Type1CodebookGenerationmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	MulticastconfigR17Type1CodebookGenerationmodeR17EnumeratedNothing = iota
	MulticastconfigR17Type1CodebookGenerationmodeR17EnumeratedMode1
	MulticastconfigR17Type1CodebookGenerationmodeR17EnumeratedMode2
)
