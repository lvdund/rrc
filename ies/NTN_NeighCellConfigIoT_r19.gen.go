// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-NeighCellConfigIoT-r19 (line 4795).

var nTNNeighCellConfigIoTR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-Config-r19", Optional: true},
		{Name: "carrierFreqIoT-r19", Optional: true},
	},
}

type NTN_NeighCellConfigIoT_r19 struct {
	Ntn_Config_r19     *NTN_Config_r19
	CarrierFreqIoT_r19 *CarrierFreqIoT_r19
}

func (ie *NTN_NeighCellConfigIoT_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNNeighCellConfigIoTR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_Config_r19 != nil, ie.CarrierFreqIoT_r19 != nil}); err != nil {
		return err
	}

	// 2. ntn-Config-r19: ref
	{
		if ie.Ntn_Config_r19 != nil {
			if err := ie.Ntn_Config_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. carrierFreqIoT-r19: ref
	{
		if ie.CarrierFreqIoT_r19 != nil {
			if err := ie.CarrierFreqIoT_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_NeighCellConfigIoT_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNNeighCellConfigIoTR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-Config-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ntn_Config_r19 = new(NTN_Config_r19)
			if err := ie.Ntn_Config_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. carrierFreqIoT-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CarrierFreqIoT_r19 = new(CarrierFreqIoT_r19)
			if err := ie.CarrierFreqIoT_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
