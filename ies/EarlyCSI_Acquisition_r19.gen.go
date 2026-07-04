// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EarlyCSI-Acquisition-r19 (line 5710).

var earlyCSIAcquisitionR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "early-NZP-CSI-RS-ResourceSet-r19"},
		{Name: "early-CSI-IM-ResourceSet-r19", Optional: true},
		{Name: "earlyCSI-ReportQuantity-r19"},
		{Name: "earlyCSI-CQI-Table-r19", Optional: true},
		{Name: "earlyCSI-CodeBookConfig-r19", Optional: true},
	},
}

const (
	EarlyCSI_Acquisition_r19_EarlyCSI_ReportQuantity_r19_Cri_RI_PMI_CQI = 0
	EarlyCSI_Acquisition_r19_EarlyCSI_ReportQuantity_r19_Spare          = 1
)

var earlyCSIAcquisitionR19EarlyCSIReportQuantityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EarlyCSI_Acquisition_r19_EarlyCSI_ReportQuantity_r19_Cri_RI_PMI_CQI, EarlyCSI_Acquisition_r19_EarlyCSI_ReportQuantity_r19_Spare},
}

type EarlyCSI_Acquisition_r19 struct {
	Early_NZP_CSI_RS_ResourceSet_r19 NZP_CSI_RS_ResourceSetId
	Early_CSI_IM_ResourceSet_r19     *CSI_IM_ResourceSetId
	EarlyCSI_ReportQuantity_r19      int64
	EarlyCSI_CQI_Table_r19           *CQI_Table
	EarlyCSI_CodeBookConfig_r19      *EarlyCSI_CodeBookConfig_r19
}

func (ie *EarlyCSI_Acquisition_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(earlyCSIAcquisitionR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Early_CSI_IM_ResourceSet_r19 != nil, ie.EarlyCSI_CQI_Table_r19 != nil, ie.EarlyCSI_CodeBookConfig_r19 != nil}); err != nil {
		return err
	}

	// 2. early-NZP-CSI-RS-ResourceSet-r19: ref
	{
		if err := ie.Early_NZP_CSI_RS_ResourceSet_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. early-CSI-IM-ResourceSet-r19: ref
	{
		if ie.Early_CSI_IM_ResourceSet_r19 != nil {
			if err := ie.Early_CSI_IM_ResourceSet_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. earlyCSI-ReportQuantity-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.EarlyCSI_ReportQuantity_r19, earlyCSIAcquisitionR19EarlyCSIReportQuantityR19Constraints); err != nil {
			return err
		}
	}

	// 5. earlyCSI-CQI-Table-r19: ref
	{
		if ie.EarlyCSI_CQI_Table_r19 != nil {
			if err := ie.EarlyCSI_CQI_Table_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. earlyCSI-CodeBookConfig-r19: ref
	{
		if ie.EarlyCSI_CodeBookConfig_r19 != nil {
			if err := ie.EarlyCSI_CodeBookConfig_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EarlyCSI_Acquisition_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(earlyCSIAcquisitionR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. early-NZP-CSI-RS-ResourceSet-r19: ref
	{
		if err := ie.Early_NZP_CSI_RS_ResourceSet_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. early-CSI-IM-ResourceSet-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Early_CSI_IM_ResourceSet_r19 = new(CSI_IM_ResourceSetId)
			if err := ie.Early_CSI_IM_ResourceSet_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. earlyCSI-ReportQuantity-r19: enumerated
	{
		v2, err := d.DecodeEnumerated(earlyCSIAcquisitionR19EarlyCSIReportQuantityR19Constraints)
		if err != nil {
			return err
		}
		ie.EarlyCSI_ReportQuantity_r19 = v2
	}

	// 5. earlyCSI-CQI-Table-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.EarlyCSI_CQI_Table_r19 = new(CQI_Table)
			if err := ie.EarlyCSI_CQI_Table_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. earlyCSI-CodeBookConfig-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.EarlyCSI_CodeBookConfig_r19 = new(EarlyCSI_CodeBookConfig_r19)
			if err := ie.EarlyCSI_CodeBookConfig_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
