package ies

// SSB-InfoNcell-r16 ::= SEQUENCE
type SsbInfoncellR16 struct {
	PhysicalcellidR16   Physcellid
	SsbIndexncellR16    *SsbIndex
	SsbConfigurationR16 *SsbConfigurationR16
}
