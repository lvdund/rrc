// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1710 (line 17064).

var bandParametersV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-AntennaSwitchingBeyond4RX-r17", Optional: true},
	},
}

var bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSRS-TxPortSwitchBeyond4Rx-r17"},
		{Name: "entryNumberAffectBeyond4Rx-r17", Optional: true},
		{Name: "entryNumberSwitchBeyond4Rx-r17", Optional: true},
	},
}

type BandParameters_v1710 struct {
	Srs_AntennaSwitchingBeyond4RX_r17 *struct {
		SupportedSRS_TxPortSwitchBeyond4Rx_r17 per.BitString
		EntryNumberAffectBeyond4Rx_r17         *int64
		EntryNumberSwitchBeyond4Rx_r17         *int64
	}
}

func (ie *BandParameters_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_AntennaSwitchingBeyond4RX_r17 != nil}); err != nil {
		return err
	}

	// 2. srs-AntennaSwitchingBeyond4RX-r17: sequence
	{
		if ie.Srs_AntennaSwitchingBeyond4RX_r17 != nil {
			c := ie.Srs_AntennaSwitchingBeyond4RX_r17
			bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq := e.NewSequenceEncoder(bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Constraints)
			if err := bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq.EncodePreamble([]bool{c.EntryNumberAffectBeyond4Rx_r17 != nil, c.EntryNumberSwitchBeyond4Rx_r17 != nil}); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.SupportedSRS_TxPortSwitchBeyond4Rx_r17, per.FixedSize(11)); err != nil {
				return err
			}
			if c.EntryNumberAffectBeyond4Rx_r17 != nil {
				if err := e.EncodeInteger((*c.EntryNumberAffectBeyond4Rx_r17), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.EntryNumberSwitchBeyond4Rx_r17 != nil {
				if err := e.EncodeInteger((*c.EntryNumberSwitchBeyond4Rx_r17), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-AntennaSwitchingBeyond4RX-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_AntennaSwitchingBeyond4RX_r17 = &struct {
				SupportedSRS_TxPortSwitchBeyond4Rx_r17 per.BitString
				EntryNumberAffectBeyond4Rx_r17         *int64
				EntryNumberSwitchBeyond4Rx_r17         *int64
			}{}
			c := ie.Srs_AntennaSwitchingBeyond4RX_r17
			bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq := d.NewSequenceDecoder(bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Constraints)
			if err := bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(11))
				if err != nil {
					return err
				}
				c.SupportedSRS_TxPortSwitchBeyond4Rx_r17 = v
			}
			if bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq.IsComponentPresent(1) {
				c.EntryNumberAffectBeyond4Rx_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberAffectBeyond4Rx_r17) = v
			}
			if bandParametersV1710SrsAntennaSwitchingBeyond4RXR17Seq.IsComponentPresent(2) {
				c.EntryNumberSwitchBeyond4Rx_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberSwitchBeyond4Rx_r17) = v
			}
		}
	}

	return nil
}
