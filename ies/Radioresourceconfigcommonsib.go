package ies

import "rrc/utils"

// RadioResourceConfigCommonSIB ::= SEQUENCE
// Extensible
type Radioresourceconfigcommonsib struct {
	RachConfigcommon         RachConfigcommon
	BcchConfig               BcchConfig
	PcchConfig               PcchConfig
	PrachConfig              PrachConfigsib
	PdschConfigcommon        PdschConfigcommon
	PuschConfigcommon        PuschConfigcommon
	PucchConfigcommon        PucchConfigcommon
	SoundingrsUlConfigcommon SoundingrsUlConfigcommon
	Uplinkpowercontrolcommon Uplinkpowercontrolcommon
	UlCyclicprefixlength     UlCyclicprefixlength
}
