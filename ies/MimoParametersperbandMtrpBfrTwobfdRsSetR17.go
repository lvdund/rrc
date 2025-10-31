package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-BFR-twoBFD-RS-Set-r17 ::= SEQUENCE
type MimoParametersperbandMtrpBfrTwobfdRsSetR17 struct {
	MaxbfdRsResourcespersetperbwpR17     MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcespersetperbwpR17
	MaxbfrR17                            utils.INTEGER `lb:0,ub:9`
	MaxbfdRsResourcesacrosssetsperbwpR17 MimoParametersperbandMtrpBfrTwobfdRsSetR17MaxbfdRsResourcesacrosssetsperbwpR17
}
