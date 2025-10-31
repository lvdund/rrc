package ies

// SL-RLC-Config-r16-sl-AM-RLC-r16 ::= SEQUENCE
// Extensible
type SlRlcConfigR16SlAmRlcR16 struct {
	SlSnFieldlengthamR16  *SnFieldlengtham
	SlTPollretransmitR16  TPollretransmit
	SlPollpduR16          Pollpdu
	SlPollbyteR16         Pollbyte
	SlMaxretxthresholdR16 SlRlcConfigR16SlAmRlcR16SlMaxretxthresholdR16
}
