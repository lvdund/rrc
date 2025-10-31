package ies

// CG-COT-Sharing-r17 ::= CHOICE
const (
	CgCotSharingR17ChoiceNothing = iota
	CgCotSharingR17ChoiceNocotSharingR17
	CgCotSharingR17ChoiceCotSharingR17
)

type CgCotSharingR17 struct {
	Choice          uint64
	NocotSharingR17 *struct{}
	CotSharingR17   *CgCotSharingR17CotSharingR17
}
