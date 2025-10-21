package ies

import "rrc/utils"

// MeasResultsCDMA2000 ::= SEQUENCE
type Measresultscdma2000 struct {
	Preregistrationstatushrpd bool
	Measresultlistcdma2000    Measresultlistcdma2000
}
