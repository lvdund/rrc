// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosConfigPerULCarrier-r18 (line 1470).

var sRSPosConfigPerULCarrierR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "freqInfo-r18"},
		{Name: "srs-PosConfig-r18"},
		{Name: "scs-SpecificCarrier-r18", Optional: true},
		{Name: "bwp-r18", Optional: true},
	},
}

type SRS_PosConfigPerULCarrier_r18 struct {
	FreqInfo_r18            ARFCN_ValueNR
	Srs_PosConfig_r18       SRS_PosConfig_r17
	Scs_SpecificCarrier_r18 *SCS_SpecificCarrier
	Bwp_r18                 *BWP
}

func (ie *SRS_PosConfigPerULCarrier_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosConfigPerULCarrierR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Scs_SpecificCarrier_r18 != nil, ie.Bwp_r18 != nil}); err != nil {
		return err
	}

	// 3. freqInfo-r18: ref
	{
		if err := ie.FreqInfo_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. srs-PosConfig-r18: ref
	{
		if err := ie.Srs_PosConfig_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. scs-SpecificCarrier-r18: ref
	{
		if ie.Scs_SpecificCarrier_r18 != nil {
			if err := ie.Scs_SpecificCarrier_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. bwp-r18: ref
	{
		if ie.Bwp_r18 != nil {
			if err := ie.Bwp_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_PosConfigPerULCarrier_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosConfigPerULCarrierR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. freqInfo-r18: ref
	{
		if err := ie.FreqInfo_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. srs-PosConfig-r18: ref
	{
		if err := ie.Srs_PosConfig_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. scs-SpecificCarrier-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Scs_SpecificCarrier_r18 = new(SCS_SpecificCarrier)
			if err := ie.Scs_SpecificCarrier_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. bwp-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Bwp_r18 = new(BWP)
			if err := ie.Bwp_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
