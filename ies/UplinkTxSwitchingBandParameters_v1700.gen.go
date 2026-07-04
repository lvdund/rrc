// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingBandParameters-v1700 (line 16994).

var uplinkTxSwitchingBandParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandIndex-r17"},
		{Name: "uplinkTxSwitching2T2T-PUSCH-TransCoherence-r17", Optional: true},
	},
}

var uplinkTxSwitchingBandParametersV1700BandIndexR17Constraints = per.Constrained(1, common.MaxSimultaneousBands)

const (
	UplinkTxSwitchingBandParameters_v1700_UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17_NonCoherent  = 0
	UplinkTxSwitchingBandParameters_v1700_UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17_FullCoherent = 1
)

var uplinkTxSwitchingBandParametersV1700UplinkTxSwitching2T2TPUSCHTransCoherenceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingBandParameters_v1700_UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17_NonCoherent, UplinkTxSwitchingBandParameters_v1700_UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17_FullCoherent},
}

type UplinkTxSwitchingBandParameters_v1700 struct {
	BandIndex_r17                                  int64
	UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 *int64
}

func (ie *UplinkTxSwitchingBandParameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingBandParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 != nil}); err != nil {
		return err
	}

	// 2. bandIndex-r17: integer
	{
		if err := e.EncodeInteger(ie.BandIndex_r17, uplinkTxSwitchingBandParametersV1700BandIndexR17Constraints); err != nil {
			return err
		}
	}

	// 3. uplinkTxSwitching2T2T-PUSCH-TransCoherence-r17: enumerated
	{
		if ie.UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17, uplinkTxSwitchingBandParametersV1700UplinkTxSwitching2T2TPUSCHTransCoherenceR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingBandParameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingBandParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandIndex-r17: integer
	{
		v0, err := d.DecodeInteger(uplinkTxSwitchingBandParametersV1700BandIndexR17Constraints)
		if err != nil {
			return err
		}
		ie.BandIndex_r17 = v0
	}

	// 3. uplinkTxSwitching2T2T-PUSCH-TransCoherence-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uplinkTxSwitchingBandParametersV1700UplinkTxSwitching2T2TPUSCHTransCoherenceR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 = &idx
		}
	}

	return nil
}
