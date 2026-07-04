// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-ResourceConfigCLI-r16 (line 9351).

var sRSResourceConfigCLIR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-Resource-r16"},
		{Name: "srs-SCS-r16"},
		{Name: "refServCellIndex-r16", Optional: true},
		{Name: "refBWP-r16"},
	},
}

type SRS_ResourceConfigCLI_r16 struct {
	Srs_Resource_r16     SRS_Resource
	Srs_SCS_r16          SubcarrierSpacing
	RefServCellIndex_r16 *ServCellIndex
	RefBWP_r16           BWP_Id
}

func (ie *SRS_ResourceConfigCLI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSResourceConfigCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefServCellIndex_r16 != nil}); err != nil {
		return err
	}

	// 3. srs-Resource-r16: ref
	{
		if err := ie.Srs_Resource_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. srs-SCS-r16: ref
	{
		if err := ie.Srs_SCS_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. refServCellIndex-r16: ref
	{
		if ie.RefServCellIndex_r16 != nil {
			if err := ie.RefServCellIndex_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. refBWP-r16: ref
	{
		if err := ie.RefBWP_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_ResourceConfigCLI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSResourceConfigCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-Resource-r16: ref
	{
		if err := ie.Srs_Resource_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. srs-SCS-r16: ref
	{
		if err := ie.Srs_SCS_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. refServCellIndex-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.RefServCellIndex_r16 = new(ServCellIndex)
			if err := ie.RefServCellIndex_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. refBWP-r16: ref
	{
		if err := ie.RefBWP_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
