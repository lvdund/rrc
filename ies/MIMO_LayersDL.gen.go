// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MIMO-LayersDL (line 21472).
// MIMO-LayersDL ::=   ENUMERATED {twoLayers, fourLayers, eightLayers}

const (
	MIMO_LayersDL_TwoLayers   = 0
	MIMO_LayersDL_FourLayers  = 1
	MIMO_LayersDL_EightLayers = 2
)

var mIMOLayersDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_LayersDL_TwoLayers, MIMO_LayersDL_FourLayers, MIMO_LayersDL_EightLayers},
}

type MIMO_LayersDL struct {
	Value int64
}

func (v *MIMO_LayersDL) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, mIMOLayersDLConstraints)
}

func (v *MIMO_LayersDL) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(mIMOLayersDLConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
