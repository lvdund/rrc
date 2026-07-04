// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqBandIndicatorEUTRA (line 20725).
// FreqBandIndicatorEUTRA ::=  INTEGER (1..maxBandsEUTRA)

var freqBandIndicatorEUTRAConstraints = per.Constrained(1, common.MaxBandsEUTRA)

type FreqBandIndicatorEUTRA struct {
	Value int64
}

func (v *FreqBandIndicatorEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, freqBandIndicatorEUTRAConstraints)
}

func (v *FreqBandIndicatorEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(freqBandIndicatorEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
