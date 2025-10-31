package ies

// TRS-ResourceSet-r17-scramblingID-Info-r17 ::= CHOICE
// Extensible
const (
	TrsResourcesetR17ScramblingidInfoR17ChoiceNothing = iota
	TrsResourcesetR17ScramblingidInfoR17ChoiceScramblingidforcommonR17
	TrsResourcesetR17ScramblingidInfoR17ChoiceScramblingidperresourcelistwith2R17
	TrsResourcesetR17ScramblingidInfoR17ChoiceScramblingidperresourcelistwith4R17
)

type TrsResourcesetR17ScramblingidInfoR17 struct {
	Choice                              uint64
	ScramblingidforcommonR17            *Scramblingid
	Scramblingidperresourcelistwith2R17 *[]Scramblingid `lb:2,ub:2`
	Scramblingidperresourcelistwith4R17 *[]Scramblingid `lb:4,ub:4`
}
