// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-LogMeasReport-r19 (line 3485).

var cSILogMeasReportR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-LogMeasInfoCellList-r19"},
		{Name: "csi-MoreLogMeasAvailable-r19", Optional: true},
	},
}

const (
	CSI_LogMeasReport_r19_Csi_MoreLogMeasAvailable_r19_True = 0
)

var cSILogMeasReportR19CsiMoreLogMeasAvailableR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_LogMeasReport_r19_Csi_MoreLogMeasAvailable_r19_True},
}

type CSI_LogMeasReport_r19 struct {
	Csi_LogMeasInfoCellList_r19  CSI_LogMeasInfoCellList_r19
	Csi_MoreLogMeasAvailable_r19 *int64
}

func (ie *CSI_LogMeasReport_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILogMeasReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_MoreLogMeasAvailable_r19 != nil}); err != nil {
		return err
	}

	// 3. csi-LogMeasInfoCellList-r19: ref
	{
		if err := ie.Csi_LogMeasInfoCellList_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. csi-MoreLogMeasAvailable-r19: enumerated
	{
		if ie.Csi_MoreLogMeasAvailable_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_MoreLogMeasAvailable_r19, cSILogMeasReportR19CsiMoreLogMeasAvailableR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_LogMeasReport_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILogMeasReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-LogMeasInfoCellList-r19: ref
	{
		if err := ie.Csi_LogMeasInfoCellList_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. csi-MoreLogMeasAvailable-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cSILogMeasReportR19CsiMoreLogMeasAvailableR19Constraints)
			if err != nil {
				return err
			}
			ie.Csi_MoreLogMeasAvailable_r19 = &idx
		}
	}

	return nil
}
