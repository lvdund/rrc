// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasObjectUTRA-FDD-r16 (line 9665).

var measObjectUTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "utra-FDD-Q-OffsetRange-r16", Optional: true},
		{Name: "cellsToRemoveList-r16", Optional: true},
		{Name: "cellsToAddModList-r16", Optional: true},
	},
}

type MeasObjectUTRA_FDD_r16 struct {
	CarrierFreq_r16            ARFCN_ValueUTRA_FDD_r16
	Utra_FDD_Q_OffsetRange_r16 *UTRA_FDD_Q_OffsetRange_r16
	CellsToRemoveList_r16      *UTRA_FDD_CellIndexList_r16
	CellsToAddModList_r16      *CellsToAddModListUTRA_FDD_r16
}

func (ie *MeasObjectUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectUTRAFDDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Utra_FDD_Q_OffsetRange_r16 != nil, ie.CellsToRemoveList_r16 != nil, ie.CellsToAddModList_r16 != nil}); err != nil {
		return err
	}

	// 3. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. utra-FDD-Q-OffsetRange-r16: ref
	{
		if ie.Utra_FDD_Q_OffsetRange_r16 != nil {
			if err := ie.Utra_FDD_Q_OffsetRange_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. cellsToRemoveList-r16: ref
	{
		if ie.CellsToRemoveList_r16 != nil {
			if err := ie.CellsToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. cellsToAddModList-r16: ref
	{
		if ie.CellsToAddModList_r16 != nil {
			if err := ie.CellsToAddModList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasObjectUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectUTRAFDDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. utra-FDD-Q-OffsetRange-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Utra_FDD_Q_OffsetRange_r16 = new(UTRA_FDD_Q_OffsetRange_r16)
			if err := ie.Utra_FDD_Q_OffsetRange_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. cellsToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CellsToRemoveList_r16 = new(UTRA_FDD_CellIndexList_r16)
			if err := ie.CellsToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. cellsToAddModList-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.CellsToAddModList_r16 = new(CellsToAddModListUTRA_FDD_r16)
			if err := ie.CellsToAddModList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
