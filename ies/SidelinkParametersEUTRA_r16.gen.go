// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SidelinkParametersEUTRA-r16 (line 25050).

var sidelinkParametersEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ParametersEUTRA1-r16", Optional: true},
		{Name: "sl-ParametersEUTRA2-r16", Optional: true},
		{Name: "sl-ParametersEUTRA3-r16", Optional: true},
		{Name: "supportedBandListSidelinkEUTRA-r16", Optional: true},
	},
}

var sidelinkParametersEUTRAR16SlParametersEUTRA1R16Constraints = per.SizeConstraints{}

var sidelinkParametersEUTRAR16SlParametersEUTRA2R16Constraints = per.SizeConstraints{}

var sidelinkParametersEUTRAR16SlParametersEUTRA3R16Constraints = per.SizeConstraints{}

var sidelinkParametersEUTRAR16SupportedBandListSidelinkEUTRAR16Constraints = per.SizeRange(1, common.MaxBandsEUTRA)

type SidelinkParametersEUTRA_r16 struct {
	Sl_ParametersEUTRA1_r16            []byte
	Sl_ParametersEUTRA2_r16            []byte
	Sl_ParametersEUTRA3_r16            []byte
	SupportedBandListSidelinkEUTRA_r16 []BandSidelinkEUTRA_r16
}

func (ie *SidelinkParametersEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkParametersEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ParametersEUTRA1_r16 != nil, ie.Sl_ParametersEUTRA2_r16 != nil, ie.Sl_ParametersEUTRA3_r16 != nil, ie.SupportedBandListSidelinkEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-ParametersEUTRA1-r16: octet-string
	{
		if ie.Sl_ParametersEUTRA1_r16 != nil {
			if err := e.EncodeOctetString(ie.Sl_ParametersEUTRA1_r16, sidelinkParametersEUTRAR16SlParametersEUTRA1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-ParametersEUTRA2-r16: octet-string
	{
		if ie.Sl_ParametersEUTRA2_r16 != nil {
			if err := e.EncodeOctetString(ie.Sl_ParametersEUTRA2_r16, sidelinkParametersEUTRAR16SlParametersEUTRA2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-ParametersEUTRA3-r16: octet-string
	{
		if ie.Sl_ParametersEUTRA3_r16 != nil {
			if err := e.EncodeOctetString(ie.Sl_ParametersEUTRA3_r16, sidelinkParametersEUTRAR16SlParametersEUTRA3R16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. supportedBandListSidelinkEUTRA-r16: sequence-of
	{
		if ie.SupportedBandListSidelinkEUTRA_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sidelinkParametersEUTRAR16SupportedBandListSidelinkEUTRAR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListSidelinkEUTRA_r16))); err != nil {
				return err
			}
			for i := range ie.SupportedBandListSidelinkEUTRA_r16 {
				if err := ie.SupportedBandListSidelinkEUTRA_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SidelinkParametersEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkParametersEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-ParametersEUTRA1-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(sidelinkParametersEUTRAR16SlParametersEUTRA1R16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ParametersEUTRA1_r16 = v
		}
	}

	// 4. sl-ParametersEUTRA2-r16: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sidelinkParametersEUTRAR16SlParametersEUTRA2R16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ParametersEUTRA2_r16 = v
		}
	}

	// 5. sl-ParametersEUTRA3-r16: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sidelinkParametersEUTRAR16SlParametersEUTRA3R16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ParametersEUTRA3_r16 = v
		}
	}

	// 6. supportedBandListSidelinkEUTRA-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sidelinkParametersEUTRAR16SupportedBandListSidelinkEUTRAR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandListSidelinkEUTRA_r16 = make([]BandSidelinkEUTRA_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandListSidelinkEUTRA_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
