package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2 ::= SEQUENCE
type CodebookconfigCodebooktypeType2 struct {
	Subtype           *CodebookconfigCodebooktypeType2Subtype
	Phasealphabetsize CodebookconfigCodebooktypeType2Phasealphabetsize
	Subbandamplitude  utils.BOOLEAN
	Numberofbeams     CodebookconfigCodebooktypeType2Numberofbeams
}
