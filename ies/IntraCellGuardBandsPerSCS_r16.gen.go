// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntraCellGuardBandsPerSCS-r16 (line 14772).

var intraCellGuardBandsPerSCSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "guardBandSCS-r16"},
		{Name: "intraCellGuardBands-r16"},
	},
}

var intraCellGuardBandsPerSCSR16IntraCellGuardBandsR16Constraints = per.SizeRange(1, 4)

type IntraCellGuardBandsPerSCS_r16 struct {
	GuardBandSCS_r16        SubcarrierSpacing
	IntraCellGuardBands_r16 []GuardBand_r16
}

func (ie *IntraCellGuardBandsPerSCS_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intraCellGuardBandsPerSCSR16Constraints)
	_ = seq

	// 1. guardBandSCS-r16: ref
	{
		if err := ie.GuardBandSCS_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. intraCellGuardBands-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(intraCellGuardBandsPerSCSR16IntraCellGuardBandsR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.IntraCellGuardBands_r16))); err != nil {
			return err
		}
		for i := range ie.IntraCellGuardBands_r16 {
			if err := ie.IntraCellGuardBands_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IntraCellGuardBandsPerSCS_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intraCellGuardBandsPerSCSR16Constraints)
	_ = seq

	// 1. guardBandSCS-r16: ref
	{
		if err := ie.GuardBandSCS_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. intraCellGuardBands-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(intraCellGuardBandsPerSCSR16IntraCellGuardBandsR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.IntraCellGuardBands_r16 = make([]GuardBand_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.IntraCellGuardBands_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
