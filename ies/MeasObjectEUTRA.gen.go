// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectEUTRA (line 9387).

var measObjectEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "allowedMeasBandwidth"},
		{Name: "cellsToRemoveListEUTRAN", Optional: true},
		{Name: "cellsToAddModListEUTRAN", Optional: true},
		{Name: "excludedCellsToRemoveListEUTRAN", Optional: true},
		{Name: "excludedCellsToAddModListEUTRAN", Optional: true},
		{Name: "eutra-PresenceAntennaPort1"},
		{Name: "eutra-Q-OffsetRange", Optional: true},
		{Name: "widebandRSRQ-Meas"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var measObjectEUTRACellsToAddModListEUTRANConstraints = per.SizeRange(1, common.MaxCellMeasEUTRA)

var measObjectEUTRAExcludedCellsToAddModListEUTRANConstraints = per.SizeRange(1, common.MaxCellMeasEUTRA)

type MeasObjectEUTRA struct {
	CarrierFreq                     ARFCN_ValueEUTRA
	AllowedMeasBandwidth            EUTRA_AllowedMeasBandwidth
	CellsToRemoveListEUTRAN         *EUTRA_CellIndexList
	CellsToAddModListEUTRAN         []EUTRA_Cell
	ExcludedCellsToRemoveListEUTRAN *EUTRA_CellIndexList
	ExcludedCellsToAddModListEUTRAN []EUTRA_ExcludedCell
	Eutra_PresenceAntennaPort1      EUTRA_PresenceAntennaPort1
	Eutra_Q_OffsetRange             *EUTRA_Q_OffsetRange
	WidebandRSRQ_Meas               bool
	AssociatedMeasGap_r17           *MeasGapId_r17
	MeasSequence_r18                *MeasSequence_r18
}

func (ie *MeasObjectEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectEUTRAConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AssociatedMeasGap_r17 != nil
	hasExtGroup1 := ie.MeasSequence_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellsToRemoveListEUTRAN != nil, ie.CellsToAddModListEUTRAN != nil, ie.ExcludedCellsToRemoveListEUTRAN != nil, ie.ExcludedCellsToAddModListEUTRAN != nil, ie.Eutra_Q_OffsetRange != nil}); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 4. allowedMeasBandwidth: ref
	{
		if err := ie.AllowedMeasBandwidth.Encode(e); err != nil {
			return err
		}
	}

	// 5. cellsToRemoveListEUTRAN: ref
	{
		if ie.CellsToRemoveListEUTRAN != nil {
			if err := ie.CellsToRemoveListEUTRAN.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. cellsToAddModListEUTRAN: sequence-of
	{
		if ie.CellsToAddModListEUTRAN != nil {
			seqOf := e.NewSequenceOfEncoder(measObjectEUTRACellsToAddModListEUTRANConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.CellsToAddModListEUTRAN))); err != nil {
				return err
			}
			for i := range ie.CellsToAddModListEUTRAN {
				if err := ie.CellsToAddModListEUTRAN[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. excludedCellsToRemoveListEUTRAN: ref
	{
		if ie.ExcludedCellsToRemoveListEUTRAN != nil {
			if err := ie.ExcludedCellsToRemoveListEUTRAN.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. excludedCellsToAddModListEUTRAN: sequence-of
	{
		if ie.ExcludedCellsToAddModListEUTRAN != nil {
			seqOf := e.NewSequenceOfEncoder(measObjectEUTRAExcludedCellsToAddModListEUTRANConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ExcludedCellsToAddModListEUTRAN))); err != nil {
				return err
			}
			for i := range ie.ExcludedCellsToAddModListEUTRAN {
				if err := ie.ExcludedCellsToAddModListEUTRAN[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. eutra-PresenceAntennaPort1: ref
	{
		if err := ie.Eutra_PresenceAntennaPort1.Encode(e); err != nil {
			return err
		}
	}

	// 10. eutra-Q-OffsetRange: ref
	{
		if ie.Eutra_Q_OffsetRange != nil {
			if err := ie.Eutra_Q_OffsetRange.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. widebandRSRQ-Meas: boolean
	{
		if err := e.EncodeBoolean(ie.WidebandRSRQ_Meas); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "associatedMeasGap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AssociatedMeasGap_r17 != nil}); err != nil {
				return err
			}

			if ie.AssociatedMeasGap_r17 != nil {
				if err := ie.AssociatedMeasGap_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "measSequence-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasSequence_r18 != nil}); err != nil {
				return err
			}

			if ie.MeasSequence_r18 != nil {
				if err := ie.MeasSequence_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasObjectEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 4. allowedMeasBandwidth: ref
	{
		if err := ie.AllowedMeasBandwidth.Decode(d); err != nil {
			return err
		}
	}

	// 5. cellsToRemoveListEUTRAN: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
			if err := ie.CellsToRemoveListEUTRAN.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. cellsToAddModListEUTRAN: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(measObjectEUTRACellsToAddModListEUTRANConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellsToAddModListEUTRAN = make([]EUTRA_Cell, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellsToAddModListEUTRAN[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. excludedCellsToRemoveListEUTRAN: ref
	{
		if seq.IsComponentPresent(4) {
			ie.ExcludedCellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
			if err := ie.ExcludedCellsToRemoveListEUTRAN.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. excludedCellsToAddModListEUTRAN: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(measObjectEUTRAExcludedCellsToAddModListEUTRANConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ExcludedCellsToAddModListEUTRAN = make([]EUTRA_ExcludedCell, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ExcludedCellsToAddModListEUTRAN[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. eutra-PresenceAntennaPort1: ref
	{
		if err := ie.Eutra_PresenceAntennaPort1.Decode(d); err != nil {
			return err
		}
	}

	// 10. eutra-Q-OffsetRange: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Eutra_Q_OffsetRange = new(EUTRA_Q_OffsetRange)
			if err := ie.Eutra_Q_OffsetRange.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. widebandRSRQ-Meas: boolean
	{
		v8, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.WidebandRSRQ_Meas = v8
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "associatedMeasGap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AssociatedMeasGap_r17 = new(MeasGapId_r17)
			if err := ie.AssociatedMeasGap_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "measSequence-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasSequence_r18 = new(MeasSequence_r18)
			if err := ie.MeasSequence_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
