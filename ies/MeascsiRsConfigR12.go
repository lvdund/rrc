package ies

import "rrc/utils"

// MeasCSI-RS-Config-r12 ::= SEQUENCE
// Extensible
type MeascsiRsConfigR12 struct {
	MeascsiRsIdR12           MeascsiRsIdR12
	PhyscellidR12            utils.INTEGER `lb:0,ub:503`
	ScramblingidentityR12    utils.INTEGER `lb:0,ub:503`
	ResourceconfigR12        utils.INTEGER `lb:0,ub:31`
	SubframeoffsetR12        utils.INTEGER `lb:0,ub:4`
	CsiRsIndividualoffsetR12 QOffsetrange
}
