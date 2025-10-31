package ies

import "rrc/utils"

// CG-SDT-Configuration-r17-sdt-DMRS-Ports-r17 ::= CHOICE
const (
	CgSdtConfigurationR17SdtDmrsPortsR17ChoiceNothing = iota
	CgSdtConfigurationR17SdtDmrsPortsR17ChoiceDmrstype1R17
	CgSdtConfigurationR17SdtDmrsPortsR17ChoiceDmrstype2R17
)

type CgSdtConfigurationR17SdtDmrsPortsR17 struct {
	Choice       uint64
	Dmrstype1R17 *utils.BITSTRING `lb:8,ub:8`
	Dmrstype2R17 *utils.BITSTRING `lb:12,ub:12`
}
