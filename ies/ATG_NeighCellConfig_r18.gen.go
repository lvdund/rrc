// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ATG-NeighCellConfig-r18 (line 4622).

var aTGNeighCellConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "atg-gNB-Location-r18", Optional: true},
		{Name: "height-gNB-r18", Optional: true},
		{Name: "carrierFreq-r18", Optional: true},
		{Name: "physCellId-r18", Optional: true},
	},
}

var aTGNeighCellConfigR18HeightGNBR18Constraints = per.Constrained(-16384, 16383)

type ATG_NeighCellConfig_r18 struct {
	Atg_GNB_Location_r18 *ReferenceLocation_r17
	Height_GNB_r18       *int64
	CarrierFreq_r18      *ARFCN_ValueNR
	PhysCellId_r18       *PhysCellId
}

func (ie *ATG_NeighCellConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aTGNeighCellConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Atg_GNB_Location_r18 != nil, ie.Height_GNB_r18 != nil, ie.CarrierFreq_r18 != nil, ie.PhysCellId_r18 != nil}); err != nil {
		return err
	}

	// 2. atg-gNB-Location-r18: ref
	{
		if ie.Atg_GNB_Location_r18 != nil {
			if err := ie.Atg_GNB_Location_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. height-gNB-r18: integer
	{
		if ie.Height_GNB_r18 != nil {
			if err := e.EncodeInteger(*ie.Height_GNB_r18, aTGNeighCellConfigR18HeightGNBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. carrierFreq-r18: ref
	{
		if ie.CarrierFreq_r18 != nil {
			if err := ie.CarrierFreq_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. physCellId-r18: ref
	{
		if ie.PhysCellId_r18 != nil {
			if err := ie.PhysCellId_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ATG_NeighCellConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aTGNeighCellConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. atg-gNB-Location-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Atg_GNB_Location_r18 = new(ReferenceLocation_r17)
			if err := ie.Atg_GNB_Location_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. height-gNB-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(aTGNeighCellConfigR18HeightGNBR18Constraints)
			if err != nil {
				return err
			}
			ie.Height_GNB_r18 = &v
		}
	}

	// 4. carrierFreq-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CarrierFreq_r18 = new(ARFCN_ValueNR)
			if err := ie.CarrierFreq_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. physCellId-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.PhysCellId_r18 = new(PhysCellId)
			if err := ie.PhysCellId_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
