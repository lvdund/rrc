// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellIndividualOffsetList-EUTRA-r18 (line 13519).

var cellIndividualOffsetListEUTRAR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r18"},
		{Name: "cellIndividualOffset-r18"},
		{Name: "carrierFreq-r18", Optional: true},
	},
}

type CellIndividualOffsetList_EUTRA_r18 struct {
	PhysCellId_r18           EUTRA_PhysCellId
	CellIndividualOffset_r18 EUTRA_Q_OffsetRange
	CarrierFreq_r18          *ARFCN_ValueEUTRA
}

func (ie *CellIndividualOffsetList_EUTRA_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellIndividualOffsetListEUTRAR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CarrierFreq_r18 != nil}); err != nil {
		return err
	}

	// 2. physCellId-r18: ref
	{
		if err := ie.PhysCellId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. cellIndividualOffset-r18: ref
	{
		if err := ie.CellIndividualOffset_r18.Encode(e); err != nil {
			return err
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

	return nil
}

func (ie *CellIndividualOffsetList_EUTRA_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellIndividualOffsetListEUTRAR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. physCellId-r18: ref
	{
		if err := ie.PhysCellId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. cellIndividualOffset-r18: ref
	{
		if err := ie.CellIndividualOffset_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. carrierFreq-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CarrierFreq_r18 = new(ARFCN_ValueEUTRA)
			if err := ie.CarrierFreq_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
