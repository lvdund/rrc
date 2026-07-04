// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-Parameters-v1820 (line 22690).

var nTNParametersV1820Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr2-Add-UE-NR-CapabilitiesNTN-r18", Optional: true},
	},
}

type NTN_Parameters_v1820 struct {
	Fr2_Add_UE_NR_CapabilitiesNTN_r18 *UE_NR_CapabilityAddFRX_Mode
}

func (ie *NTN_Parameters_v1820) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNParametersV1820Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fr2_Add_UE_NR_CapabilitiesNTN_r18 != nil}); err != nil {
		return err
	}

	// 2. fr2-Add-UE-NR-CapabilitiesNTN-r18: ref
	{
		if ie.Fr2_Add_UE_NR_CapabilitiesNTN_r18 != nil {
			if err := ie.Fr2_Add_UE_NR_CapabilitiesNTN_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_Parameters_v1820) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNParametersV1820Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fr2-Add-UE-NR-CapabilitiesNTN-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Fr2_Add_UE_NR_CapabilitiesNTN_r18 = new(UE_NR_CapabilityAddFRX_Mode)
			if err := ie.Fr2_Add_UE_NR_CapabilitiesNTN_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
