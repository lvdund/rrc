// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MIMO-LayersUL (line 21474).
// MIMO-LayersUL ::=   ENUMERATED {oneLayer, twoLayers, fourLayers}

const (
	MIMO_LayersUL_OneLayer   = 0
	MIMO_LayersUL_TwoLayers  = 1
	MIMO_LayersUL_FourLayers = 2
)

var mIMOLayersULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_LayersUL_OneLayer, MIMO_LayersUL_TwoLayers, MIMO_LayersUL_FourLayers},
}

type MIMO_LayersUL struct {
	Value int64
}

func (v *MIMO_LayersUL) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, mIMOLayersULConstraints)
}

func (v *MIMO_LayersUL) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(mIMOLayersULConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
