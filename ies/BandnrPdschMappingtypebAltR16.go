package ies

import "rrc/utils"

// BandNR-pdsch-MappingTypeB-Alt-r16 ::= ENUMERATED
type BandnrPdschMappingtypebAltR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdschMappingtypebAltR16EnumeratedNothing = iota
	BandnrPdschMappingtypebAltR16EnumeratedSupported
)
