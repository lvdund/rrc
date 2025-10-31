package ies

// BandParameters ::= CHOICE
const (
	BandparametersChoiceNothing = iota
	BandparametersChoiceEutra
	BandparametersChoiceNr
)

type Bandparameters struct {
	Choice uint64
	Eutra  *BandparametersEutra
	Nr     *BandparametersNr
}
