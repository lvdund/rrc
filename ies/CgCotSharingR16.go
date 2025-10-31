package ies

// CG-COT-Sharing-r16 ::= CHOICE
const (
	CgCotSharingR16ChoiceNothing = iota
	CgCotSharingR16ChoiceNocotSharingR16
	CgCotSharingR16ChoiceCotSharingR16
)

type CgCotSharingR16 struct {
	Choice          uint64
	NocotSharingR16 *struct{}
	CotSharingR16   *CgCotSharingR16CotSharingR16
}
