package ies

// UL-AM-RLC ::= SEQUENCE
type UlAmRlc struct {
	TPollretransmit  TPollretransmit
	Pollpdu          Pollpdu
	Pollbyte         Pollbyte
	Maxretxthreshold UlAmRlcMaxretxthreshold
}
