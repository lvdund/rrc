package ies

// MAC-MainConfig-NB-r13 ::= SEQUENCE
// Extensible
type MacMainconfigNbR13 struct {
	UlSchConfigR13                 *MacMainconfigNbR13UlSchConfigR13
	DrxConfigR13                   *DrxConfigNbR13
	TimealignmenttimerdedicatedR13 Timealignmenttimer
	LogicalchannelsrConfigR13      *MacMainconfigNbR13LogicalchannelsrConfigR13
}
