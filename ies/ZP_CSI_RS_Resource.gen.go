// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ZP-CSI-RS-Resource (line 16445).

var zPCSIRSResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "zp-CSI-RS-ResourceId"},
		{Name: "resourceMapping"},
		{Name: "periodicityAndOffset", Optional: true},
	},
}

type ZP_CSI_RS_Resource struct {
	Zp_CSI_RS_ResourceId ZP_CSI_RS_ResourceId
	ResourceMapping      CSI_RS_ResourceMapping
	PeriodicityAndOffset *CSI_ResourcePeriodicityAndOffset
}

func (ie *ZP_CSI_RS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(zPCSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicityAndOffset != nil}); err != nil {
		return err
	}

	// 3. zp-CSI-RS-ResourceId: ref
	{
		if err := ie.Zp_CSI_RS_ResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 4. resourceMapping: ref
	{
		if err := ie.ResourceMapping.Encode(e); err != nil {
			return err
		}
	}

	// 5. periodicityAndOffset: ref
	{
		if ie.PeriodicityAndOffset != nil {
			if err := ie.PeriodicityAndOffset.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ZP_CSI_RS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(zPCSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. zp-CSI-RS-ResourceId: ref
	{
		if err := ie.Zp_CSI_RS_ResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 4. resourceMapping: ref
	{
		if err := ie.ResourceMapping.Decode(d); err != nil {
			return err
		}
	}

	// 5. periodicityAndOffset: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
			if err := ie.PeriodicityAndOffset.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
