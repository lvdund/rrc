// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ModulationOrder (line 22540).
// ModulationOrder ::= ENUMERATED {bpsk-halfpi, bpsk, qpsk, qam16, qam64, qam256}

const (
	ModulationOrder_Bpsk_Halfpi = 0
	ModulationOrder_Bpsk        = 1
	ModulationOrder_Qpsk        = 2
	ModulationOrder_Qam16       = 3
	ModulationOrder_Qam64       = 4
	ModulationOrder_Qam256      = 5
)

var modulationOrderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ModulationOrder_Bpsk_Halfpi, ModulationOrder_Bpsk, ModulationOrder_Qpsk, ModulationOrder_Qam16, ModulationOrder_Qam64, ModulationOrder_Qam256},
}

type ModulationOrder struct {
	Value int64
}

func (v *ModulationOrder) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, modulationOrderConstraints)
}

func (v *ModulationOrder) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(modulationOrderConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
