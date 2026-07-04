// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-DL-PRS-PDC-Info-r17 (line 10603).

var nRDLPRSPDCInfoR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-DL-PRS-PDC-ResourceSet-r17", Optional: true},
	},
}

type NR_DL_PRS_PDC_Info_r17 struct {
	Nr_DL_PRS_PDC_ResourceSet_r17 *NR_DL_PRS_PDC_ResourceSet_r17
}

func (ie *NR_DL_PRS_PDC_Info_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDLPRSPDCInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nr_DL_PRS_PDC_ResourceSet_r17 != nil}); err != nil {
		return err
	}

	// 3. nr-DL-PRS-PDC-ResourceSet-r17: ref
	{
		if ie.Nr_DL_PRS_PDC_ResourceSet_r17 != nil {
			if err := ie.Nr_DL_PRS_PDC_ResourceSet_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_DL_PRS_PDC_Info_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDLPRSPDCInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. nr-DL-PRS-PDC-ResourceSet-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Nr_DL_PRS_PDC_ResourceSet_r17 = new(NR_DL_PRS_PDC_ResourceSet_r17)
			if err := ie.Nr_DL_PRS_PDC_ResourceSet_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
