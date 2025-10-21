package ies

import "rrc/utils"

// BandParameters-v1430 ::= SEQUENCE
type BandparametersV1430 struct {
	BandparametersdlV1430           *MimoCaParametersperbobcV1430
	Ul256qamR14                     *utils.ENUMERATED
	Ul256qamPerccInfolistR14        *interface{}
	SrsCapabilityperbandpairlistR14 *interface{}
}
