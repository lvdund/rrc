// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectNR-SL-r18 (line 9621).

var measObjectNRSLR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Frequency-r18"},
		{Name: "tx-PoolMeasToRemoveList-r18", Optional: true},
		{Name: "tx-PoolMeasToAddModList-r18", Optional: true},
	},
}

var measObjectNRSLR18SlFrequencyR18Constraints = per.Constrained(1, common.MaxNrofFreqSL_r16)

type MeasObjectNR_SL_r18 struct {
	Sl_Frequency_r18            int64
	Tx_PoolMeasToRemoveList_r18 *Tx_PoolMeasList_r16
	Tx_PoolMeasToAddModList_r18 *Tx_PoolMeasList_r16
}

func (ie *MeasObjectNR_SL_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectNRSLR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tx_PoolMeasToRemoveList_r18 != nil, ie.Tx_PoolMeasToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-Frequency-r18: integer
	{
		if err := e.EncodeInteger(ie.Sl_Frequency_r18, measObjectNRSLR18SlFrequencyR18Constraints); err != nil {
			return err
		}
	}

	// 3. tx-PoolMeasToRemoveList-r18: ref
	{
		if ie.Tx_PoolMeasToRemoveList_r18 != nil {
			if err := ie.Tx_PoolMeasToRemoveList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. tx-PoolMeasToAddModList-r18: ref
	{
		if ie.Tx_PoolMeasToAddModList_r18 != nil {
			if err := ie.Tx_PoolMeasToAddModList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasObjectNR_SL_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectNRSLR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-Frequency-r18: integer
	{
		v0, err := d.DecodeInteger(measObjectNRSLR18SlFrequencyR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Frequency_r18 = v0
	}

	// 3. tx-PoolMeasToRemoveList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Tx_PoolMeasToRemoveList_r18 = new(Tx_PoolMeasList_r16)
			if err := ie.Tx_PoolMeasToRemoveList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. tx-PoolMeasToAddModList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Tx_PoolMeasToAddModList_r18 = new(Tx_PoolMeasList_r16)
			if err := ie.Tx_PoolMeasToAddModList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
