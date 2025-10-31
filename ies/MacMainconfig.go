package ies

// MAC-MainConfig ::= SEQUENCE
// Extensible
type MacMainconfig struct {
	UlSchConfig                 *MacMainconfigUlSchConfig
	DrxConfig                   *DrxConfig
	Timealignmenttimerdedicated Timealignmenttimer
	PhrConfig                   *MacMainconfigPhrConfig
}
