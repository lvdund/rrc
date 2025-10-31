package ies

import "rrc/utils"

// DummyA ::= SEQUENCE
type Dummya struct {
	MaxnumbernzpCsiRsPercc                       utils.INTEGER `lb:0,ub:32`
	MaxnumberportsacrossnzpCsiRsPercc            DummyaMaxnumberportsacrossnzpCsiRsPercc
	MaxnumbercsImPercc                           DummyaMaxnumbercsImPercc
	MaxnumbersimultaneouscsiRsActbwpAllcc        DummyaMaxnumbersimultaneouscsiRsActbwpAllcc
	TotalnumberportssimultaneouscsiRsActbwpAllcc DummyaTotalnumberportssimultaneouscsiRsActbwpAllcc
}
