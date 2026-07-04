// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1770 (line 17608).

var cAParametersNRV1770Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "parallelTxPUCCH-PUSCH-SamePriority-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1770_ParallelTxPUCCH_PUSCH_SamePriority_r17_Supported = 0
)

var cAParametersNRV1770ParallelTxPUCCHPUSCHSamePriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1770_ParallelTxPUCCH_PUSCH_SamePriority_r17_Supported},
}

type CA_ParametersNR_v1770 struct {
	ParallelTxPUCCH_PUSCH_SamePriority_r17 *int64
}

func (ie *CA_ParametersNR_v1770) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1770Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 != nil}); err != nil {
		return err
	}

	// 2. parallelTxPUCCH-PUSCH-SamePriority-r17: enumerated
	{
		if ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxPUCCH_PUSCH_SamePriority_r17, cAParametersNRV1770ParallelTxPUCCHPUSCHSamePriorityR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1770) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1770Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. parallelTxPUCCH-PUSCH-SamePriority-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1770ParallelTxPUCCHPUSCHSamePriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 = &idx
		}
	}

	return nil
}
