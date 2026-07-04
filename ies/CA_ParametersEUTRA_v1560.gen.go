// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersEUTRA-v1560 (line 17259).

var cAParametersEUTRAV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fd-MIMO-TotalWeightedLayers", Optional: true},
	},
}

var cAParametersEUTRAV1560FdMIMOTotalWeightedLayersConstraints = per.Constrained(2, 128)

type CA_ParametersEUTRA_v1560 struct {
	Fd_MIMO_TotalWeightedLayers *int64
}

func (ie *CA_ParametersEUTRA_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersEUTRAV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fd_MIMO_TotalWeightedLayers != nil}); err != nil {
		return err
	}

	// 2. fd-MIMO-TotalWeightedLayers: integer
	{
		if ie.Fd_MIMO_TotalWeightedLayers != nil {
			if err := e.EncodeInteger(*ie.Fd_MIMO_TotalWeightedLayers, cAParametersEUTRAV1560FdMIMOTotalWeightedLayersConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersEUTRA_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersEUTRAV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fd-MIMO-TotalWeightedLayers: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cAParametersEUTRAV1560FdMIMOTotalWeightedLayersConstraints)
			if err != nil {
				return err
			}
			ie.Fd_MIMO_TotalWeightedLayers = &v
		}
	}

	return nil
}
