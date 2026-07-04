// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandParameters-v1730 (line 17076).

var bandParametersV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-SwitchingAffectedBandsListNR-r17"},
	},
}

var bandParametersV1730SrsSwitchingAffectedBandsListNRR17Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandParameters_v1730 struct {
	Srs_SwitchingAffectedBandsListNR_r17 []SRS_SwitchingAffectedBandsNR_r17
}

func (ie *BandParameters_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1730Constraints)
	_ = seq

	// 1. srs-SwitchingAffectedBandsListNR-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(bandParametersV1730SrsSwitchingAffectedBandsListNRR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Srs_SwitchingAffectedBandsListNR_r17))); err != nil {
			return err
		}
		for i := range ie.Srs_SwitchingAffectedBandsListNR_r17 {
			if err := ie.Srs_SwitchingAffectedBandsListNR_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1730Constraints)
	_ = seq

	// 1. srs-SwitchingAffectedBandsListNR-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(bandParametersV1730SrsSwitchingAffectedBandsListNRR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Srs_SwitchingAffectedBandsListNR_r17 = make([]SRS_SwitchingAffectedBandsNR_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Srs_SwitchingAffectedBandsListNR_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
