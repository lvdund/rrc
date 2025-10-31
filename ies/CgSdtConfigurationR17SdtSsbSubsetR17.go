package ies

import "rrc/utils"

// CG-SDT-Configuration-r17-sdt-SSB-Subset-r17 ::= CHOICE
const (
	CgSdtConfigurationR17SdtSsbSubsetR17ChoiceNothing = iota
	CgSdtConfigurationR17SdtSsbSubsetR17ChoiceShortbitmapR17
	CgSdtConfigurationR17SdtSsbSubsetR17ChoiceMediumbitmapR17
	CgSdtConfigurationR17SdtSsbSubsetR17ChoiceLongbitmapR17
)

type CgSdtConfigurationR17SdtSsbSubsetR17 struct {
	Choice          uint64
	ShortbitmapR17  *utils.BITSTRING `lb:4,ub:4`
	MediumbitmapR17 *utils.BITSTRING `lb:8,ub:8`
	LongbitmapR17   *utils.BITSTRING `lb:64,ub:64`
}
