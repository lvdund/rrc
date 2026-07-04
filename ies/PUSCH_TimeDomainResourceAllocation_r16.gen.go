// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-TimeDomainResourceAllocation-r16 (line 12711).

var pUSCHTimeDomainResourceAllocationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "k2-r16", Optional: true},
		{Name: "puschAllocationList-r16"},
	},
}

var pUSCHTimeDomainResourceAllocationR16K2R16Constraints = per.Constrained(0, 32)

var pUSCHTimeDomainResourceAllocationR16PuschAllocationListR16Constraints = per.SizeRange(1, common.MaxNrofMultiplePUSCHs_r16)

type PUSCH_TimeDomainResourceAllocation_r16 struct {
	K2_r16                  *int64
	PuschAllocationList_r16 []PUSCH_Allocation_r16
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHTimeDomainResourceAllocationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.K2_r16 != nil}); err != nil {
		return err
	}

	// 3. k2-r16: integer
	{
		if ie.K2_r16 != nil {
			if err := e.EncodeInteger(*ie.K2_r16, pUSCHTimeDomainResourceAllocationR16K2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. puschAllocationList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pUSCHTimeDomainResourceAllocationR16PuschAllocationListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.PuschAllocationList_r16))); err != nil {
			return err
		}
		for i := range ie.PuschAllocationList_r16 {
			if err := ie.PuschAllocationList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHTimeDomainResourceAllocationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. k2-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUSCHTimeDomainResourceAllocationR16K2R16Constraints)
			if err != nil {
				return err
			}
			ie.K2_r16 = &v
		}
	}

	// 4. puschAllocationList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pUSCHTimeDomainResourceAllocationR16PuschAllocationListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.PuschAllocationList_r16 = make([]PUSCH_Allocation_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.PuschAllocationList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
