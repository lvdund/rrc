// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasObjectNR-SL-r16 (line 9616).

var measObjectNRSLR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tx-PoolMeasToRemoveList-r16", Optional: true},
		{Name: "tx-PoolMeasToAddModList-r16", Optional: true},
	},
}

type MeasObjectNR_SL_r16 struct {
	Tx_PoolMeasToRemoveList_r16 *Tx_PoolMeasList_r16
	Tx_PoolMeasToAddModList_r16 *Tx_PoolMeasList_r16
}

func (ie *MeasObjectNR_SL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectNRSLR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tx_PoolMeasToRemoveList_r16 != nil, ie.Tx_PoolMeasToAddModList_r16 != nil}); err != nil {
		return err
	}

	// 2. tx-PoolMeasToRemoveList-r16: ref
	{
		if ie.Tx_PoolMeasToRemoveList_r16 != nil {
			if err := ie.Tx_PoolMeasToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. tx-PoolMeasToAddModList-r16: ref
	{
		if ie.Tx_PoolMeasToAddModList_r16 != nil {
			if err := ie.Tx_PoolMeasToAddModList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasObjectNR_SL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectNRSLR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tx-PoolMeasToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Tx_PoolMeasToRemoveList_r16 = new(Tx_PoolMeasList_r16)
			if err := ie.Tx_PoolMeasToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. tx-PoolMeasToAddModList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Tx_PoolMeasToAddModList_r16 = new(Tx_PoolMeasList_r16)
			if err := ie.Tx_PoolMeasToAddModList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
