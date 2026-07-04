// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-IndirectPathAddChange-r18 (line 27337).

var sLIndirectPathAddChangeR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-IndirectPathRelayUE-Identity-r18"},
		{Name: "sl-IndirectPathCellIdentity-r18"},
		{Name: "t421-r18", Optional: true},
	},
}

const (
	SL_IndirectPathAddChange_r18_T421_r18_Ms50    = 0
	SL_IndirectPathAddChange_r18_T421_r18_Ms100   = 1
	SL_IndirectPathAddChange_r18_T421_r18_Ms150   = 2
	SL_IndirectPathAddChange_r18_T421_r18_Ms200   = 3
	SL_IndirectPathAddChange_r18_T421_r18_Ms500   = 4
	SL_IndirectPathAddChange_r18_T421_r18_Ms1000  = 5
	SL_IndirectPathAddChange_r18_T421_r18_Ms2000  = 6
	SL_IndirectPathAddChange_r18_T421_r18_Ms10000 = 7
)

var sLIndirectPathAddChangeR18T421R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_IndirectPathAddChange_r18_T421_r18_Ms50, SL_IndirectPathAddChange_r18_T421_r18_Ms100, SL_IndirectPathAddChange_r18_T421_r18_Ms150, SL_IndirectPathAddChange_r18_T421_r18_Ms200, SL_IndirectPathAddChange_r18_T421_r18_Ms500, SL_IndirectPathAddChange_r18_T421_r18_Ms1000, SL_IndirectPathAddChange_r18_T421_r18_Ms2000, SL_IndirectPathAddChange_r18_T421_r18_Ms10000},
}

type SL_IndirectPathAddChange_r18 struct {
	Sl_IndirectPathRelayUE_Identity_r18 SL_SourceIdentity_r17
	Sl_IndirectPathCellIdentity_r18     CellIdentity
	T421_r18                            *int64
}

func (ie *SL_IndirectPathAddChange_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLIndirectPathAddChangeR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T421_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-IndirectPathRelayUE-Identity-r18: ref
	{
		if err := ie.Sl_IndirectPathRelayUE_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-IndirectPathCellIdentity-r18: ref
	{
		if err := ie.Sl_IndirectPathCellIdentity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. t421-r18: enumerated
	{
		if ie.T421_r18 != nil {
			if err := e.EncodeEnumerated(*ie.T421_r18, sLIndirectPathAddChangeR18T421R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_IndirectPathAddChange_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLIndirectPathAddChangeR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-IndirectPathRelayUE-Identity-r18: ref
	{
		if err := ie.Sl_IndirectPathRelayUE_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-IndirectPathCellIdentity-r18: ref
	{
		if err := ie.Sl_IndirectPathCellIdentity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. t421-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLIndirectPathAddChangeR18T421R18Constraints)
			if err != nil {
				return err
			}
			ie.T421_r18 = &idx
		}
	}

	return nil
}
