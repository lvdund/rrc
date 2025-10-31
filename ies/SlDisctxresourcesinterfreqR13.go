package ies

// SL-DiscTxResourcesInterFreq-r13 ::= CHOICE
const (
	SlDisctxresourcesinterfreqR13ChoiceNothing = iota
	SlDisctxresourcesinterfreqR13ChoiceAcquiresiFromcarrierR13
	SlDisctxresourcesinterfreqR13ChoiceDisctxpoolcommonR13
	SlDisctxresourcesinterfreqR13ChoiceRequestdedicatedR13
	SlDisctxresourcesinterfreqR13ChoiceNotxoncarrierR13
)

type SlDisctxresourcesinterfreqR13 struct {
	Choice                  uint64
	AcquiresiFromcarrierR13 *struct{}
	DisctxpoolcommonR13     *SlDisctxpoollistR12
	RequestdedicatedR13     *struct{}
	NotxoncarrierR13        *struct{}
}
