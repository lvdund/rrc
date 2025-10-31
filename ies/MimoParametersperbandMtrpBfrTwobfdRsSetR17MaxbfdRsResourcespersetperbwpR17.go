package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-BFR-twoBFD-RS-Set-r17-maxBFD-RS-resourcesPerSetPerBWP-r17 ::= ENUMERATED
type MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcespersetperbwpR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcespersetperbwpR17EnumeratedNothing = iota
	MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcespersetperbwpR17EnumeratedN1
	MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcespersetperbwpR17EnumeratedN2
)
