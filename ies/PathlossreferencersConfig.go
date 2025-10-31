package ies

// PathlossReferenceRS-Config ::= CHOICE
const (
	PathlossreferencersConfigChoiceNothing = iota
	PathlossreferencersConfigChoiceSsbIndex
	PathlossreferencersConfigChoiceCsiRsIndex
)

type PathlossreferencersConfig struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
