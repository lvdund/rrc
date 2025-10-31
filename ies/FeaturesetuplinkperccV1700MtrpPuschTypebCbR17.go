package ies

import "rrc/utils"

// FeatureSetUplinkPerCC-v1700-mTRP-PUSCH-TypeB-CB-r17 ::= ENUMERATED
type FeaturesetuplinkperccV1700MtrpPuschTypebCbR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkperccV1700MtrpPuschTypebCbR17EnumeratedNothing = iota
	FeaturesetuplinkperccV1700MtrpPuschTypebCbR17EnumeratedN1
	FeaturesetuplinkperccV1700MtrpPuschTypebCbR17EnumeratedN2
	FeaturesetuplinkperccV1700MtrpPuschTypebCbR17EnumeratedN4
)
