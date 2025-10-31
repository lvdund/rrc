package ies

// UL-AM-RLC ::= SEQUENCE
type UlAmRlc struct {
	SnFieldlength    *SnFieldlengtham
	TPollretransmit  TPollretransmit
	Pollpdu          Pollpdu
	Pollbyte         Pollbyte
	Maxretxthreshold UlAmRlcMaxretxthreshold
}
