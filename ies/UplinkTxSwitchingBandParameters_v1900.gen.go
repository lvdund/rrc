// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingBandParameters-v1900 (line 17001).

var uplinkTxSwitchingBandParametersV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandIndex-r19"},
		{Name: "uplinkTxSwitching3Tx-PUSCH-TransCoherence-DualUL-v1900", Optional: true},
	},
}

var uplinkTxSwitchingBandParametersV1900BandIndexR19Constraints = per.Constrained(1, common.MaxSimultaneousBands)

const (
	UplinkTxSwitchingBandParameters_v1900_UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900_NonCoherent  = 0
	UplinkTxSwitchingBandParameters_v1900_UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900_FullCoherent = 1
)

var uplinkTxSwitchingBandParametersV1900UplinkTxSwitching3TxPUSCHTransCoherenceDualULV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingBandParameters_v1900_UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900_NonCoherent, UplinkTxSwitchingBandParameters_v1900_UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900_FullCoherent},
}

type UplinkTxSwitchingBandParameters_v1900 struct {
	BandIndex_r19                                          int64
	UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900 *int64
}

func (ie *UplinkTxSwitchingBandParameters_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingBandParametersV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900 != nil}); err != nil {
		return err
	}

	// 2. bandIndex-r19: integer
	{
		if err := e.EncodeInteger(ie.BandIndex_r19, uplinkTxSwitchingBandParametersV1900BandIndexR19Constraints); err != nil {
			return err
		}
	}

	// 3. uplinkTxSwitching3Tx-PUSCH-TransCoherence-DualUL-v1900: enumerated
	{
		if ie.UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900, uplinkTxSwitchingBandParametersV1900UplinkTxSwitching3TxPUSCHTransCoherenceDualULV1900Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingBandParameters_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingBandParametersV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandIndex-r19: integer
	{
		v0, err := d.DecodeInteger(uplinkTxSwitchingBandParametersV1900BandIndexR19Constraints)
		if err != nil {
			return err
		}
		ie.BandIndex_r19 = v0
	}

	// 3. uplinkTxSwitching3Tx-PUSCH-TransCoherence-DualUL-v1900: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uplinkTxSwitchingBandParametersV1900UplinkTxSwitching3TxPUSCHTransCoherenceDualULV1900Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching3Tx_PUSCH_TransCoherence_DualUL_v1900 = &idx
		}
	}

	return nil
}
