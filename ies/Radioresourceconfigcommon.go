package ies

// RadioResourceConfigCommon ::= SEQUENCE
// Extensible
type Radioresourceconfigcommon struct {
	RachConfigcommon         *RachConfigcommon
	PrachConfig              PrachConfig
	PdschConfigcommon        *PdschConfigcommon
	PuschConfigcommon        PuschConfigcommon
	PhichConfig              *PhichConfig
	PucchConfigcommon        *PucchConfigcommon
	SoundingrsUlConfigcommon *SoundingrsUlConfigcommon
	Uplinkpowercontrolcommon *Uplinkpowercontrolcommon
	Antennainfocommon        *Antennainfocommon
	PMax                     *PMax
	TddConfig                *TddConfig
	UlCyclicprefixlength     UlCyclicprefixlength
}
