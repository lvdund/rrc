package ies

// AppLayerMeasConfig-r17 ::= SEQUENCE
// Extensible
type ApplayermeasconfigR17 struct {
	MeasconfigapplayertoaddmodlistR17  *[]MeasconfigapplayerR17   `lb:1,ub:maxNrofAppLayerMeasR17`
	MeasconfigapplayertoreleaselistR17 *[]MeasconfigapplayeridR17 `lb:1,ub:maxNrofAppLayerMeasR17`
	RrcSegallowedR17                   *ApplayermeasconfigR17RrcSegallowedR17
}
