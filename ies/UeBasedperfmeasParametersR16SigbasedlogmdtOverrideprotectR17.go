package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-sigBasedLogMDT-OverrideProtect-r17 ::= ENUMERATED
type UeBasedperfmeasParametersR16SigbasedlogmdtOverrideprotectR17 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16SigbasedlogmdtOverrideprotectR17EnumeratedNothing = iota
	UeBasedperfmeasParametersR16SigbasedlogmdtOverrideprotectR17EnumeratedSupported
)
