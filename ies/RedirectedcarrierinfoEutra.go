package ies

// RedirectedCarrierInfo-EUTRA ::= SEQUENCE
type RedirectedcarrierinfoEutra struct {
	Eutrafrequency ArfcnValueeutra
	Cntype         *RedirectedcarrierinfoEutraCntype
}
