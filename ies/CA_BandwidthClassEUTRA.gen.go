// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-BandwidthClassEUTRA (line 17234).
// CA-BandwidthClassEUTRA ::=          ENUMERATED {a, b, c, d, e, f, ...}

const (
	CA_BandwidthClassEUTRA_A = 0
	CA_BandwidthClassEUTRA_B = 1
	CA_BandwidthClassEUTRA_C = 2
	CA_BandwidthClassEUTRA_D = 3
	CA_BandwidthClassEUTRA_E = 4
	CA_BandwidthClassEUTRA_F = 5
)

var cABandwidthClassEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{CA_BandwidthClassEUTRA_A, CA_BandwidthClassEUTRA_B, CA_BandwidthClassEUTRA_C, CA_BandwidthClassEUTRA_D, CA_BandwidthClassEUTRA_E, CA_BandwidthClassEUTRA_F},
}

type CA_BandwidthClassEUTRA struct {
	Value int64
}

func (v *CA_BandwidthClassEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cABandwidthClassEUTRAConstraints)
}

func (v *CA_BandwidthClassEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cABandwidthClassEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
