package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2-phaseAlphabetSize ::= ENUMERATED
type CodebookconfigCodebooktypeType2Phasealphabetsize struct {
	Value utils.ENUMERATED
}

const (
	CodebookconfigCodebooktypeType2PhasealphabetsizeEnumeratedNothing = iota
	CodebookconfigCodebooktypeType2PhasealphabetsizeEnumeratedN4
	CodebookconfigCodebooktypeType2PhasealphabetsizeEnumeratedN8
)
