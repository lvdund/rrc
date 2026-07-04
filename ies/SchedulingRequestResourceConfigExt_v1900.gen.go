// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestResourceConfigExt-v1900 (line 14317).

var schedulingRequestResourceConfigExtV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "symbolType-r19", Optional: true},
	},
}

const (
	SchedulingRequestResourceConfigExt_v1900_SymbolType_r19_Sbfd     = 0
	SchedulingRequestResourceConfigExt_v1900_SymbolType_r19_Non_Sbfd = 1
)

var schedulingRequestResourceConfigExtV1900SymbolTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingRequestResourceConfigExt_v1900_SymbolType_r19_Sbfd, SchedulingRequestResourceConfigExt_v1900_SymbolType_r19_Non_Sbfd},
}

type SchedulingRequestResourceConfigExt_v1900 struct {
	SymbolType_r19 *int64
}

func (ie *SchedulingRequestResourceConfigExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestResourceConfigExtV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SymbolType_r19 != nil}); err != nil {
		return err
	}

	// 2. symbolType-r19: enumerated
	{
		if ie.SymbolType_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SymbolType_r19, schedulingRequestResourceConfigExtV1900SymbolTypeR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestResourceConfigExtV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. symbolType-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(schedulingRequestResourceConfigExtV1900SymbolTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.SymbolType_r19 = &idx
		}
	}

	return nil
}
