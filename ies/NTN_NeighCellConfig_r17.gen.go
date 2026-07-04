// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-NeighCellConfig-r17 (line 4531).

var nTNNeighCellConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-Config-r17", Optional: true},
		{Name: "carrierFreq-r17", Optional: true},
		{Name: "physCellId-r17", Optional: true},
	},
}

type NTN_NeighCellConfig_r17 struct {
	Ntn_Config_r17  *NTN_Config_r17
	CarrierFreq_r17 *ARFCN_ValueNR
	PhysCellId_r17  *PhysCellId
}

func (ie *NTN_NeighCellConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNNeighCellConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_Config_r17 != nil, ie.CarrierFreq_r17 != nil, ie.PhysCellId_r17 != nil}); err != nil {
		return err
	}

	// 2. ntn-Config-r17: ref
	{
		if ie.Ntn_Config_r17 != nil {
			if err := ie.Ntn_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. carrierFreq-r17: ref
	{
		if ie.CarrierFreq_r17 != nil {
			if err := ie.CarrierFreq_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. physCellId-r17: ref
	{
		if ie.PhysCellId_r17 != nil {
			if err := ie.PhysCellId_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_NeighCellConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNNeighCellConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-Config-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ntn_Config_r17 = new(NTN_Config_r17)
			if err := ie.Ntn_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. carrierFreq-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CarrierFreq_r17 = new(ARFCN_ValueNR)
			if err := ie.CarrierFreq_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. physCellId-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PhysCellId_r17 = new(PhysCellId)
			if err := ie.PhysCellId_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
