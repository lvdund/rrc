package ies

import "rrc/utils"

// MeasResultsCDMA2000 ::= SEQUENCE
type Measresultscdma2000 struct {
	Preregistrationstatushrpd utils.BOOLEAN
	Measresultlistcdma2000    Measresultlistcdma2000
}
