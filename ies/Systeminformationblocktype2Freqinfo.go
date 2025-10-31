package ies

// SystemInformationBlockType2-freqInfo ::= SEQUENCE
type Systeminformationblocktype2Freqinfo struct {
	UlCarrierfreq              *ArfcnValueeutra
	UlBandwidth                *Systeminformationblocktype2FreqinfoUlBandwidth
	Additionalspectrumemission Additionalspectrumemission
}
