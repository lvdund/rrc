// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersEUTRA-v1570 (line 17263).

var cAParametersEUTRAV1570Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-1024QAM-TotalWeightedLayers", Optional: true},
	},
}

var cAParametersEUTRAV1570Dl1024QAMTotalWeightedLayersConstraints = per.Constrained(0, 10)

type CA_ParametersEUTRA_v1570 struct {
	Dl_1024QAM_TotalWeightedLayers *int64
}

func (ie *CA_ParametersEUTRA_v1570) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersEUTRAV1570Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_1024QAM_TotalWeightedLayers != nil}); err != nil {
		return err
	}

	// 2. dl-1024QAM-TotalWeightedLayers: integer
	{
		if ie.Dl_1024QAM_TotalWeightedLayers != nil {
			if err := e.EncodeInteger(*ie.Dl_1024QAM_TotalWeightedLayers, cAParametersEUTRAV1570Dl1024QAMTotalWeightedLayersConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersEUTRA_v1570) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersEUTRAV1570Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-1024QAM-TotalWeightedLayers: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cAParametersEUTRAV1570Dl1024QAMTotalWeightedLayersConstraints)
			if err != nil {
				return err
			}
			ie.Dl_1024QAM_TotalWeightedLayers = &v
		}
	}

	return nil
}
