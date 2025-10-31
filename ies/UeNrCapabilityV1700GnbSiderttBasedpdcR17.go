package ies

import "rrc/utils"

// UE-NR-Capability-v1700-gNB-SideRTT-BasedPDC-r17 ::= ENUMERATED
type UeNrCapabilityV1700GnbSiderttBasedpdcR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700GnbSiderttBasedpdcR17EnumeratedNothing = iota
	UeNrCapabilityV1700GnbSiderttBasedpdcR17EnumeratedSupported
)
