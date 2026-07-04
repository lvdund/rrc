// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RedCap-ConfigCommonSIB-r17 (line 2091).

var redCapConfigCommonSIBR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "halfDuplexRedCapAllowed-r17", Optional: true},
		{Name: "cellBarredRedCap-r17", Optional: true},
	},
}

const (
	RedCap_ConfigCommonSIB_r17_HalfDuplexRedCapAllowed_r17_True = 0
)

var redCapConfigCommonSIBR17HalfDuplexRedCapAllowedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCap_ConfigCommonSIB_r17_HalfDuplexRedCapAllowed_r17_True},
}

const (
	RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap1Rx_r17_Barred    = 0
	RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap1Rx_r17_NotBarred = 1
)

var redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap1RxR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap1Rx_r17_Barred, RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap1Rx_r17_NotBarred},
}

const (
	RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap2Rx_r17_Barred    = 0
	RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap2Rx_r17_NotBarred = 1
)

var redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap2RxR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap2Rx_r17_Barred, RedCap_ConfigCommonSIB_r17_CellBarredRedCap_r17_CellBarredRedCap2Rx_r17_NotBarred},
}

type RedCap_ConfigCommonSIB_r17 struct {
	HalfDuplexRedCapAllowed_r17 *int64
	CellBarredRedCap_r17        *struct {
		CellBarredRedCap1Rx_r17 int64
		CellBarredRedCap2Rx_r17 int64
	}
}

func (ie *RedCap_ConfigCommonSIB_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(redCapConfigCommonSIBR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HalfDuplexRedCapAllowed_r17 != nil, ie.CellBarredRedCap_r17 != nil}); err != nil {
		return err
	}

	// 3. halfDuplexRedCapAllowed-r17: enumerated
	{
		if ie.HalfDuplexRedCapAllowed_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HalfDuplexRedCapAllowed_r17, redCapConfigCommonSIBR17HalfDuplexRedCapAllowedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. cellBarredRedCap-r17: sequence
	{
		if ie.CellBarredRedCap_r17 != nil {
			c := ie.CellBarredRedCap_r17
			if err := e.EncodeEnumerated(c.CellBarredRedCap1Rx_r17, redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap1RxR17Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.CellBarredRedCap2Rx_r17, redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap2RxR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RedCap_ConfigCommonSIB_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(redCapConfigCommonSIBR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. halfDuplexRedCapAllowed-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(redCapConfigCommonSIBR17HalfDuplexRedCapAllowedR17Constraints)
			if err != nil {
				return err
			}
			ie.HalfDuplexRedCapAllowed_r17 = &idx
		}
	}

	// 4. cellBarredRedCap-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.CellBarredRedCap_r17 = &struct {
				CellBarredRedCap1Rx_r17 int64
				CellBarredRedCap2Rx_r17 int64
			}{}
			c := ie.CellBarredRedCap_r17
			{
				v, err := d.DecodeEnumerated(redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap1RxR17Constraints)
				if err != nil {
					return err
				}
				c.CellBarredRedCap1Rx_r17 = v
			}
			{
				v, err := d.DecodeEnumerated(redCapConfigCommonSIBR17CellBarredRedCapR17CellBarredRedCap2RxR17Constraints)
				if err != nil {
					return err
				}
				c.CellBarredRedCap2Rx_r17 = v
			}
		}
	}

	return nil
}
