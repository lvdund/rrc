package ies

// SR-WithoutHARQ-ACK-Config-NB-r15 ::= CHOICE
const (
	SrWithoutharqAckConfigNbR15ChoiceNothing = iota
	SrWithoutharqAckConfigNbR15ChoiceRelease
	SrWithoutharqAckConfigNbR15ChoiceSetup
)

type SrWithoutharqAckConfigNbR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SrWithoutharqAckConfigNbR15Setup
}
