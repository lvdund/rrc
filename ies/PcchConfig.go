package ies

// PCCH-Config ::= SEQUENCE
type PcchConfig struct {
	Defaultpagingcycle PcchConfigDefaultpagingcycle
	Nb                 PcchConfigNb
}
