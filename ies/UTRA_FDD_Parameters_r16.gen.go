// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UTRA-FDD-Parameters-r16 (line 20927).

var uTRAFDDParametersR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandListUTRA-FDD-r16"},
	},
}

var uTRAFDDParametersR16SupportedBandListUTRAFDDR16Constraints = per.SizeRange(1, common.MaxBandsUTRA_FDD_r16)

type UTRA_FDD_Parameters_r16 struct {
	SupportedBandListUTRA_FDD_r16 []SupportedBandUTRA_FDD_r16
}

func (ie *UTRA_FDD_Parameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uTRAFDDParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. supportedBandListUTRA-FDD-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uTRAFDDParametersR16SupportedBandListUTRAFDDR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListUTRA_FDD_r16))); err != nil {
			return err
		}
		for i := range ie.SupportedBandListUTRA_FDD_r16 {
			if err := ie.SupportedBandListUTRA_FDD_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UTRA_FDD_Parameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uTRAFDDParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. supportedBandListUTRA-FDD-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uTRAFDDParametersR16SupportedBandListUTRAFDDR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedBandListUTRA_FDD_r16 = make([]SupportedBandUTRA_FDD_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SupportedBandListUTRA_FDD_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
