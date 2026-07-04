// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DRX-ConfigUC-Info-r17 (line 27123).

var sLDRXConfigUCInfoR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIndex-r17", Optional: true},
		{Name: "sl-DRX-ConfigUC-r17", Optional: true},
	},
}

type SL_DRX_ConfigUC_Info_r17 struct {
	Sl_DestinationIndex_r17 *SL_DestinationIndex_r16
	Sl_DRX_ConfigUC_r17     *SL_DRX_ConfigUC_r17
}

func (ie *SL_DRX_ConfigUC_Info_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXConfigUCInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DestinationIndex_r17 != nil, ie.Sl_DRX_ConfigUC_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DestinationIndex-r17: ref
	{
		if ie.Sl_DestinationIndex_r17 != nil {
			if err := ie.Sl_DestinationIndex_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-DRX-ConfigUC-r17: ref
	{
		if ie.Sl_DRX_ConfigUC_r17 != nil {
			if err := ie.Sl_DRX_ConfigUC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_DRX_ConfigUC_Info_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXConfigUCInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DestinationIndex-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_DestinationIndex_r17 = new(SL_DestinationIndex_r16)
			if err := ie.Sl_DestinationIndex_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-DRX-ConfigUC-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_DRX_ConfigUC_r17 = new(SL_DRX_ConfigUC_r17)
			if err := ie.Sl_DRX_ConfigUC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
