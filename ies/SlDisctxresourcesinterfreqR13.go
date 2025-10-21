package ies

import "rrc/utils"

// SL-DiscTxResourcesInterFreq-r13 ::= CHOICE
type SlDisctxresourcesinterfreqR13 interface {
	isSlDisctxresourcesinterfreqR13()
}

type SlDisctxresourcesinterfreqR13AcquiresiFromcarrierR13 struct {
	Value struct{}
}

func (*SlDisctxresourcesinterfreqR13AcquiresiFromcarrierR13) isSlDisctxresourcesinterfreqR13() {}

type SlDisctxresourcesinterfreqR13DisctxpoolcommonR13 struct {
	Value SlDisctxpoollistR12
}

func (*SlDisctxresourcesinterfreqR13DisctxpoolcommonR13) isSlDisctxresourcesinterfreqR13() {}

type SlDisctxresourcesinterfreqR13RequestdedicatedR13 struct {
	Value struct{}
}

func (*SlDisctxresourcesinterfreqR13RequestdedicatedR13) isSlDisctxresourcesinterfreqR13() {}

type SlDisctxresourcesinterfreqR13NotxoncarrierR13 struct {
	Value struct{}
}

func (*SlDisctxresourcesinterfreqR13NotxoncarrierR13) isSlDisctxresourcesinterfreqR13() {}
