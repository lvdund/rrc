package ies

import "rrc/utils"

// BandParameters-v1710-srs-AntennaSwitchingBeyond4RX-r17 ::= SEQUENCE
type BandparametersV1710SrsAntennaswitchingbeyond4rxR17 struct {
	SupportedsrsTxportswitchbeyond4rxR17 utils.BITSTRING `lb:11,ub:11`
	Entrynumberaffectbeyond4rxR17        *utils.INTEGER  `lb:0,ub:32`
	Entrynumberswitchbeyond4rxR17        *utils.INTEGER  `lb:0,ub:32`
}
