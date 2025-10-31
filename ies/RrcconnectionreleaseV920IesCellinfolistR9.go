package ies

// RRCConnectionRelease-v920-IEs-cellInfoList-r9 ::= CHOICE
// Extensible
const (
	RrcconnectionreleaseV920IesCellinfolistR9ChoiceNothing = iota
	RrcconnectionreleaseV920IesCellinfolistR9ChoiceGeranR9
	RrcconnectionreleaseV920IesCellinfolistR9ChoiceUtraFddR9
	RrcconnectionreleaseV920IesCellinfolistR9ChoiceUtraTddR9
	RrcconnectionreleaseV920IesCellinfolistR9ChoiceUtraTddR10
)

type RrcconnectionreleaseV920IesCellinfolistR9 struct {
	Choice     uint64
	GeranR9    *CellinfolistgeranR9
	UtraFddR9  *CellinfolistutraFddR9
	UtraTddR9  *CellinfolistutraTddR9
	UtraTddR10 *CellinfolistutraTddR10
}
