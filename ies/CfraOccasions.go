package ies

// CFRA-occasions ::= SEQUENCE
type CfraOccasions struct {
	RachConfiggeneric  RachConfiggeneric
	SsbPerrachOccasion *CfraOccasionsSsbPerrachOccasion
}
