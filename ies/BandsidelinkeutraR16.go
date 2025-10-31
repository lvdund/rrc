package ies

// BandSidelinkEUTRA-r16 ::= SEQUENCE
type BandsidelinkeutraR16 struct {
	FreqbandsidelinkeutraR16          Freqbandindicatoreutra
	GnbScheduledmode3sidelinkeutraR16 *BandsidelinkeutraR16GnbScheduledmode3sidelinkeutraR16
	GnbScheduledmode4sidelinkeutraR16 *BandsidelinkeutraR16GnbScheduledmode4sidelinkeutraR16
}
