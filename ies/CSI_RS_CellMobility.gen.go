// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-RS-CellMobility (line 7527).

var cSIRSCellMobilityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellId"},
		{Name: "csi-rs-MeasurementBW"},
		{Name: "density", Optional: true},
		{Name: "csi-rs-ResourceList-Mobility"},
	},
}

const (
	CSI_RS_CellMobility_Density_D1 = 0
	CSI_RS_CellMobility_Density_D3 = 1
)

var cSIRSCellMobilityDensityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_CellMobility_Density_D1, CSI_RS_CellMobility_Density_D3},
}

var cSIRSCellMobilityCsiRsResourceListMobilityConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesRRM)

const (
	CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size24  = 0
	CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size48  = 1
	CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size96  = 2
	CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size192 = 3
	CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size264 = 4
)

var cSIRSCellMobilityCsiRsMeasurementBWNrofPRBsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size24, CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size48, CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size96, CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size192, CSI_RS_CellMobility_Csi_Rs_MeasurementBW_NrofPRBs_Size264},
}

type CSI_RS_CellMobility struct {
	CellId               PhysCellId
	Csi_Rs_MeasurementBW struct {
		NrofPRBs int64
		StartPRB int64
	}
	Density                      *int64
	Csi_Rs_ResourceList_Mobility []CSI_RS_Resource_Mobility
}

func (ie *CSI_RS_CellMobility) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSCellMobilityConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Density != nil}); err != nil {
		return err
	}

	// 2. cellId: ref
	{
		if err := ie.CellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-rs-MeasurementBW: sequence
	{
		{
			c := &ie.Csi_Rs_MeasurementBW
			if err := e.EncodeEnumerated(c.NrofPRBs, cSIRSCellMobilityCsiRsMeasurementBWNrofPRBsConstraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.StartPRB, per.Constrained(0, 2169)); err != nil {
				return err
			}
		}
	}

	// 4. density: enumerated
	{
		if ie.Density != nil {
			if err := e.EncodeEnumerated(*ie.Density, cSIRSCellMobilityDensityConstraints); err != nil {
				return err
			}
		}
	}

	// 5. csi-rs-ResourceList-Mobility: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSIRSCellMobilityCsiRsResourceListMobilityConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_Rs_ResourceList_Mobility))); err != nil {
			return err
		}
		for i := range ie.Csi_Rs_ResourceList_Mobility {
			if err := ie.Csi_Rs_ResourceList_Mobility[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_RS_CellMobility) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSCellMobilityConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cellId: ref
	{
		if err := ie.CellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-rs-MeasurementBW: sequence
	{
		{
			c := &ie.Csi_Rs_MeasurementBW
			{
				v, err := d.DecodeEnumerated(cSIRSCellMobilityCsiRsMeasurementBWNrofPRBsConstraints)
				if err != nil {
					return err
				}
				c.NrofPRBs = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 2169))
				if err != nil {
					return err
				}
				c.StartPRB = v
			}
		}
	}

	// 4. density: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cSIRSCellMobilityDensityConstraints)
			if err != nil {
				return err
			}
			ie.Density = &idx
		}
	}

	// 5. csi-rs-ResourceList-Mobility: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSIRSCellMobilityCsiRsResourceListMobilityConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_Rs_ResourceList_Mobility = make([]CSI_RS_Resource_Mobility, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_Rs_ResourceList_Mobility[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
