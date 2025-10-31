package ies

// RedirectedCarrierInfo ::= CHOICE
// Extensible
const (
	RedirectedcarrierinfoChoiceNothing = iota
	RedirectedcarrierinfoChoiceNr
	RedirectedcarrierinfoChoiceEutra
)

type Redirectedcarrierinfo struct {
	Choice uint64
	Nr     *Carrierinfonr
	Eutra  *RedirectedcarrierinfoEutra
}
