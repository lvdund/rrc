// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-Parameters (line 20901).

var eUTRAParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandListEUTRA"},
		{Name: "eutra-ParametersCommon", Optional: true},
		{Name: "eutra-ParametersXDD-Diff", Optional: true},
	},
}

var eUTRAParametersSupportedBandListEUTRAConstraints = per.SizeRange(1, common.MaxBandsEUTRA)

type EUTRA_Parameters struct {
	SupportedBandListEUTRA   []FreqBandIndicatorEUTRA
	Eutra_ParametersCommon   *EUTRA_ParametersCommon
	Eutra_ParametersXDD_Diff *EUTRA_ParametersXDD_Diff
}

func (ie *EUTRA_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_ParametersCommon != nil, ie.Eutra_ParametersXDD_Diff != nil}); err != nil {
		return err
	}

	// 3. supportedBandListEUTRA: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(eUTRAParametersSupportedBandListEUTRAConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListEUTRA))); err != nil {
			return err
		}
		for i := range ie.SupportedBandListEUTRA {
			if err := ie.SupportedBandListEUTRA[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. eutra-ParametersCommon: ref
	{
		if ie.Eutra_ParametersCommon != nil {
			if err := ie.Eutra_ParametersCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. eutra-ParametersXDD-Diff: ref
	{
		if ie.Eutra_ParametersXDD_Diff != nil {
			if err := ie.Eutra_ParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. supportedBandListEUTRA: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(eUTRAParametersSupportedBandListEUTRAConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedBandListEUTRA = make([]FreqBandIndicatorEUTRA, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SupportedBandListEUTRA[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. eutra-ParametersCommon: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Eutra_ParametersCommon = new(EUTRA_ParametersCommon)
			if err := ie.Eutra_ParametersCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. eutra-ParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Eutra_ParametersXDD_Diff = new(EUTRA_ParametersXDD_Diff)
			if err := ie.Eutra_ParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
