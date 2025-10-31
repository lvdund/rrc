package ies

// CFRA-TwoStep-r16 ::= SEQUENCE
// Extensible
type CfraTwostepR16 struct {
	OccasionstwostepraR16 *CfraTwostepR16OccasionstwostepraR16
	MsgaCfraPuschR16      MsgaPuschResourceR16
	MsgaTransmaxR16       *CfraTwostepR16MsgaTransmaxR16
	ResourcestwostepR16   CfraTwostepR16ResourcestwostepR16
}
