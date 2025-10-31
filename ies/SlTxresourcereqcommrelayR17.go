package ies

// SL-TxResourceReqCommRelay-r17 ::= CHOICE
const (
	SlTxresourcereqcommrelayR17ChoiceNothing = iota
	SlTxresourcereqcommrelayR17ChoiceSlTxresourcereql2u2nRelayR17
	SlTxresourcereqcommrelayR17ChoiceSlTxresourcereql3u2nRelayR17
)

type SlTxresourcereqcommrelayR17 struct {
	Choice                       uint64
	SlTxresourcereql2u2nRelayR17 *SlTxresourcereql2u2nRelayR17
	SlTxresourcereql3u2nRelayR17 *SlTxresourcereqR16
}
