package ies

import "rrc/utils"

// BandSidelink-r16-psfch-FormatZeroSidelink-r16-psfch-TxNumber ::= ENUMERATED
type BandsidelinkR16PsfchFormatzerosidelinkR16PsfchTxnumber struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16PsfchFormatzerosidelinkR16PsfchTxnumberEnumeratedNothing = iota
	BandsidelinkR16PsfchFormatzerosidelinkR16PsfchTxnumberEnumeratedN4
	BandsidelinkR16PsfchFormatzerosidelinkR16PsfchTxnumberEnumeratedN8
	BandsidelinkR16PsfchFormatzerosidelinkR16PsfchTxnumberEnumeratedN16
)
