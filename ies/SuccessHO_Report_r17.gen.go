// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SuccessHO-Report-r17 (line 3307).

var successHOReportR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sourceCellInfo-r17"},
		{Name: "targetCellInfo-r17"},
		{Name: "measResultNeighCells-r17", Optional: true},
		{Name: "locationInfo-r17", Optional: true},
		{Name: "timeSinceCHO-Reconfig-r17", Optional: true},
		{Name: "shr-Cause-r17", Optional: true},
		{Name: "ra-InformationCommon-r17", Optional: true},
		{Name: "upInterruptionTimeAtHO-r17", Optional: true},
		{Name: "c-RNTI-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	SuccessHO_Report_r17_Ext_Rach_Less_r19_True = 0
)

var successHOReportR17ExtRachLessR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Report_r17_Ext_Rach_Less_r19_True},
}

var successHOReportR17ExtTargetPSCellIDR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r19"},
		{Name: "pci-arfcn-r19"},
	},
}

const (
	SuccessHO_Report_r17_Ext_TargetPSCellID_r19_CellGlobalId_r19 = 0
	SuccessHO_Report_r17_Ext_TargetPSCellID_r19_Pci_Arfcn_r19    = 1
)

var successHOReportR17SourceCellInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sourcePCellId-r17"},
		{Name: "sourceCellMeas-r17", Optional: true},
		{Name: "rlf-InSourceDAPS-r17", Optional: true},
	},
}

const (
	SuccessHO_Report_r17_SourceCellInfo_r17_Rlf_InSourceDAPS_r17_True = 0
)

var successHOReportR17SourceCellInfoR17RlfInSourceDAPSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Report_r17_SourceCellInfo_r17_Rlf_InSourceDAPS_r17_True},
}

var successHOReportR17TargetCellInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "targetPCellId-r17"},
		{Name: "targetCellMeas-r17", Optional: true},
	},
}

var successHOReportR17MeasResultNeighCellsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListNR-r17", Optional: true},
		{Name: "measResultListEUTRA-r17", Optional: true},
	},
}

var successHOReportR17ExtEutraTargetCellInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "targetPCellId-r18"},
		{Name: "targetCellMeas-r18", Optional: true},
	},
}

var successHOReportR17ExtEutraTargetCellInfoR18TargetPCellIdR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r18"},
		{Name: "pci-arfcn-r18"},
	},
}

const (
	SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_CellGlobalId_r18 = 0
	SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_Pci_Arfcn_r18    = 1
)

var successHOReportR17ExtSourcePSCellInfoR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sourcePSCellId-r19"},
		{Name: "sourcePSCellMeas-r19", Optional: true},
	},
}

type SuccessHO_Report_r17 struct {
	SourceCellInfo_r17 struct {
		SourcePCellId_r17    CGI_Info_Logging_r16
		SourceCellMeas_r17   *MeasResultSuccessHONR_r17
		Rlf_InSourceDAPS_r17 *int64
	}
	TargetCellInfo_r17 struct {
		TargetPCellId_r17  CGI_Info_Logging_r16
		TargetCellMeas_r17 *MeasResultSuccessHONR_r17
	}
	MeasResultNeighCells_r17 *struct {
		MeasResultListNR_r17    *MeasResultList2NR_r16
		MeasResultListEUTRA_r17 *MeasResultList2EUTRA_r16
	}
	LocationInfo_r17           *LocationInfo_r16
	TimeSinceCHO_Reconfig_r17  *TimeSinceCHO_Reconfig_r17
	Shr_Cause_r17              *SHR_Cause_r17
	Ra_InformationCommon_r17   *RA_InformationCommon_r16
	UpInterruptionTimeAtHO_r17 *UPInterruptionTimeAtHO_r17
	C_RNTI_r17                 *RNTI_Value
	TargetCell_PCI_ARFCN_r17   *PCI_ARFCN_NR_r16
	Eutra_TargetCellInfo_r18   *struct {
		TargetPCellId_r18 struct {
			Choice           int
			CellGlobalId_r18 *CGI_Info_Logging_r16
			Pci_Arfcn_r18    *PCI_ARFCN_EUTRA_r16
		}
		TargetCellMeas_r18 *MeasQuantityResultsEUTRA
	}
	MeasResultServCellRSSI_r18      *RSSI_Range_r16
	MeasResultNeighFreqListRSSI_r18 *MeasResultNeighFreqListRSSI_r18
	Eutra_C_RNTI_r18                *EUTRA_C_RNTI
	TimeSinceSHR_r18                *TimeSinceSHR_r18
	SourceCellMeasL1_r19            *MeasResultL1_r19
	TargetCellMeasL1_r19            *MeasResultL1_r19
	NeighCellsMeasL1ListNR_r19      *MeasResultList3NR_r19
	Rach_Less_r19                   *int64
	SourcePSCellInfo_r19            *struct {
		SourcePSCellId_r19   CGI_Info_Logging_r16
		SourcePSCellMeas_r19 *MeasResultSuccessHONR_r17
	}
	Cho_WithCandidateSCGInfoList_r19 *CHO_WithCandidateSCGInfoList_r19
	TargetPSCellID_r19               *struct {
		Choice           int
		CellGlobalId_r19 *CGI_Info_Logging_r16
		Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
	}
}

func (ie *SuccessHO_Report_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(successHOReportR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.TargetCell_PCI_ARFCN_r17 != nil
	hasExtGroup1 := ie.Eutra_TargetCellInfo_r18 != nil || ie.MeasResultServCellRSSI_r18 != nil || ie.MeasResultNeighFreqListRSSI_r18 != nil || ie.Eutra_C_RNTI_r18 != nil || ie.TimeSinceSHR_r18 != nil
	hasExtGroup2 := ie.SourceCellMeasL1_r19 != nil || ie.TargetCellMeasL1_r19 != nil || ie.NeighCellsMeasL1ListNR_r19 != nil || ie.Rach_Less_r19 != nil || ie.SourcePSCellInfo_r19 != nil || ie.Cho_WithCandidateSCGInfoList_r19 != nil || ie.TargetPSCellID_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultNeighCells_r17 != nil, ie.LocationInfo_r17 != nil, ie.TimeSinceCHO_Reconfig_r17 != nil, ie.Shr_Cause_r17 != nil, ie.Ra_InformationCommon_r17 != nil, ie.UpInterruptionTimeAtHO_r17 != nil, ie.C_RNTI_r17 != nil}); err != nil {
		return err
	}

	// 3. sourceCellInfo-r17: sequence
	{
		{
			c := &ie.SourceCellInfo_r17
			successHOReportR17SourceCellInfoR17Seq := e.NewSequenceEncoder(successHOReportR17SourceCellInfoR17Constraints)
			if err := successHOReportR17SourceCellInfoR17Seq.EncodePreamble([]bool{c.SourceCellMeas_r17 != nil, c.Rlf_InSourceDAPS_r17 != nil}); err != nil {
				return err
			}
			if err := c.SourcePCellId_r17.Encode(e); err != nil {
				return err
			}
			if c.SourceCellMeas_r17 != nil {
				if err := c.SourceCellMeas_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Rlf_InSourceDAPS_r17 != nil {
				if err := e.EncodeEnumerated((*c.Rlf_InSourceDAPS_r17), successHOReportR17SourceCellInfoR17RlfInSourceDAPSR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. targetCellInfo-r17: sequence
	{
		{
			c := &ie.TargetCellInfo_r17
			successHOReportR17TargetCellInfoR17Seq := e.NewSequenceEncoder(successHOReportR17TargetCellInfoR17Constraints)
			if err := successHOReportR17TargetCellInfoR17Seq.EncodePreamble([]bool{c.TargetCellMeas_r17 != nil}); err != nil {
				return err
			}
			if err := c.TargetPCellId_r17.Encode(e); err != nil {
				return err
			}
			if c.TargetCellMeas_r17 != nil {
				if err := c.TargetCellMeas_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. measResultNeighCells-r17: sequence
	{
		if ie.MeasResultNeighCells_r17 != nil {
			c := ie.MeasResultNeighCells_r17
			successHOReportR17MeasResultNeighCellsR17Seq := e.NewSequenceEncoder(successHOReportR17MeasResultNeighCellsR17Constraints)
			if err := successHOReportR17MeasResultNeighCellsR17Seq.EncodePreamble([]bool{c.MeasResultListNR_r17 != nil, c.MeasResultListEUTRA_r17 != nil}); err != nil {
				return err
			}
			if c.MeasResultListNR_r17 != nil {
				if err := c.MeasResultListNR_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.MeasResultListEUTRA_r17 != nil {
				if err := c.MeasResultListEUTRA_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. locationInfo-r17: ref
	{
		if ie.LocationInfo_r17 != nil {
			if err := ie.LocationInfo_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. timeSinceCHO-Reconfig-r17: ref
	{
		if ie.TimeSinceCHO_Reconfig_r17 != nil {
			if err := ie.TimeSinceCHO_Reconfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. shr-Cause-r17: ref
	{
		if ie.Shr_Cause_r17 != nil {
			if err := ie.Shr_Cause_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. ra-InformationCommon-r17: ref
	{
		if ie.Ra_InformationCommon_r17 != nil {
			if err := ie.Ra_InformationCommon_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. upInterruptionTimeAtHO-r17: ref
	{
		if ie.UpInterruptionTimeAtHO_r17 != nil {
			if err := ie.UpInterruptionTimeAtHO_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. c-RNTI-r17: ref
	{
		if ie.C_RNTI_r17 != nil {
			if err := ie.C_RNTI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "targetCell-PCI-ARFCN-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TargetCell_PCI_ARFCN_r17 != nil}); err != nil {
				return err
			}

			if ie.TargetCell_PCI_ARFCN_r17 != nil {
				if err := ie.TargetCell_PCI_ARFCN_r17.Encode(ex); err != nil {
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
					{Name: "eutra-TargetCellInfo-r18", Optional: true},
					{Name: "measResultServCellRSSI-r18", Optional: true},
					{Name: "measResultNeighFreqListRSSI-r18", Optional: true},
					{Name: "eutra-C-RNTI-r18", Optional: true},
					{Name: "timeSinceSHR-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Eutra_TargetCellInfo_r18 != nil, ie.MeasResultServCellRSSI_r18 != nil, ie.MeasResultNeighFreqListRSSI_r18 != nil, ie.Eutra_C_RNTI_r18 != nil, ie.TimeSinceSHR_r18 != nil}); err != nil {
				return err
			}

			if ie.Eutra_TargetCellInfo_r18 != nil {
				c := ie.Eutra_TargetCellInfo_r18
				successHOReportR17ExtEutraTargetCellInfoR18Seq := ex.NewSequenceEncoder(successHOReportR17ExtEutraTargetCellInfoR18Constraints)
				if err := successHOReportR17ExtEutraTargetCellInfoR18Seq.EncodePreamble([]bool{c.TargetCellMeas_r18 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(successHOReportR17ExtEutraTargetCellInfoR18TargetPCellIdR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.TargetPCellId_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.TargetPCellId_r18.Choice {
					case SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_CellGlobalId_r18:
						if err := c.TargetPCellId_r18.CellGlobalId_r18.Encode(ex); err != nil {
							return err
						}
					case SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_Pci_Arfcn_r18:
						if err := c.TargetPCellId_r18.Pci_Arfcn_r18.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.TargetCellMeas_r18 != nil {
					if err := c.TargetCellMeas_r18.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MeasResultServCellRSSI_r18 != nil {
				if err := ie.MeasResultServCellRSSI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultNeighFreqListRSSI_r18 != nil {
				if err := ie.MeasResultNeighFreqListRSSI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Eutra_C_RNTI_r18 != nil {
				if err := ie.Eutra_C_RNTI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TimeSinceSHR_r18 != nil {
				if err := ie.TimeSinceSHR_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sourceCellMeasL1-r19", Optional: true},
					{Name: "targetCellMeasL1-r19", Optional: true},
					{Name: "neighCellsMeasL1ListNR-r19", Optional: true},
					{Name: "rach-Less-r19", Optional: true},
					{Name: "sourcePSCellInfo-r19", Optional: true},
					{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
					{Name: "targetPSCellID-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SourceCellMeasL1_r19 != nil, ie.TargetCellMeasL1_r19 != nil, ie.NeighCellsMeasL1ListNR_r19 != nil, ie.Rach_Less_r19 != nil, ie.SourcePSCellInfo_r19 != nil, ie.Cho_WithCandidateSCGInfoList_r19 != nil, ie.TargetPSCellID_r19 != nil}); err != nil {
				return err
			}

			if ie.SourceCellMeasL1_r19 != nil {
				if err := ie.SourceCellMeasL1_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TargetCellMeasL1_r19 != nil {
				if err := ie.TargetCellMeasL1_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.NeighCellsMeasL1ListNR_r19 != nil {
				if err := ie.NeighCellsMeasL1ListNR_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Rach_Less_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Rach_Less_r19, successHOReportR17ExtRachLessR19Constraints); err != nil {
					return err
				}
			}

			if ie.SourcePSCellInfo_r19 != nil {
				c := ie.SourcePSCellInfo_r19
				successHOReportR17ExtSourcePSCellInfoR19Seq := ex.NewSequenceEncoder(successHOReportR17ExtSourcePSCellInfoR19Constraints)
				if err := successHOReportR17ExtSourcePSCellInfoR19Seq.EncodePreamble([]bool{c.SourcePSCellMeas_r19 != nil}); err != nil {
					return err
				}
				if err := c.SourcePSCellId_r19.Encode(ex); err != nil {
					return err
				}
				if c.SourcePSCellMeas_r19 != nil {
					if err := c.SourcePSCellMeas_r19.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cho_WithCandidateSCGInfoList_r19 != nil {
				if err := ie.Cho_WithCandidateSCGInfoList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TargetPSCellID_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(successHOReportR17ExtTargetPSCellIDR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.TargetPSCellID_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.TargetPSCellID_r19).Choice {
				case SuccessHO_Report_r17_Ext_TargetPSCellID_r19_CellGlobalId_r19:
					if err := (*ie.TargetPSCellID_r19).CellGlobalId_r19.Encode(ex); err != nil {
						return err
					}
				case SuccessHO_Report_r17_Ext_TargetPSCellID_r19_Pci_Arfcn_r19:
					if err := (*ie.TargetPSCellID_r19).Pci_Arfcn_r19.Encode(ex); err != nil {
						return err
					}
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

func (ie *SuccessHO_Report_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(successHOReportR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sourceCellInfo-r17: sequence
	{
		{
			c := &ie.SourceCellInfo_r17
			successHOReportR17SourceCellInfoR17Seq := d.NewSequenceDecoder(successHOReportR17SourceCellInfoR17Constraints)
			if err := successHOReportR17SourceCellInfoR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.SourcePCellId_r17.Decode(d); err != nil {
					return err
				}
			}
			if successHOReportR17SourceCellInfoR17Seq.IsComponentPresent(1) {
				c.SourceCellMeas_r17 = new(MeasResultSuccessHONR_r17)
				if err := (*c.SourceCellMeas_r17).Decode(d); err != nil {
					return err
				}
			}
			if successHOReportR17SourceCellInfoR17Seq.IsComponentPresent(2) {
				c.Rlf_InSourceDAPS_r17 = new(int64)
				v, err := d.DecodeEnumerated(successHOReportR17SourceCellInfoR17RlfInSourceDAPSR17Constraints)
				if err != nil {
					return err
				}
				(*c.Rlf_InSourceDAPS_r17) = v
			}
		}
	}

	// 4. targetCellInfo-r17: sequence
	{
		{
			c := &ie.TargetCellInfo_r17
			successHOReportR17TargetCellInfoR17Seq := d.NewSequenceDecoder(successHOReportR17TargetCellInfoR17Constraints)
			if err := successHOReportR17TargetCellInfoR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.TargetPCellId_r17.Decode(d); err != nil {
					return err
				}
			}
			if successHOReportR17TargetCellInfoR17Seq.IsComponentPresent(1) {
				c.TargetCellMeas_r17 = new(MeasResultSuccessHONR_r17)
				if err := (*c.TargetCellMeas_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. measResultNeighCells-r17: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultNeighCells_r17 = &struct {
				MeasResultListNR_r17    *MeasResultList2NR_r16
				MeasResultListEUTRA_r17 *MeasResultList2EUTRA_r16
			}{}
			c := ie.MeasResultNeighCells_r17
			successHOReportR17MeasResultNeighCellsR17Seq := d.NewSequenceDecoder(successHOReportR17MeasResultNeighCellsR17Constraints)
			if err := successHOReportR17MeasResultNeighCellsR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if successHOReportR17MeasResultNeighCellsR17Seq.IsComponentPresent(0) {
				c.MeasResultListNR_r17 = new(MeasResultList2NR_r16)
				if err := (*c.MeasResultListNR_r17).Decode(d); err != nil {
					return err
				}
			}
			if successHOReportR17MeasResultNeighCellsR17Seq.IsComponentPresent(1) {
				c.MeasResultListEUTRA_r17 = new(MeasResultList2EUTRA_r16)
				if err := (*c.MeasResultListEUTRA_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. locationInfo-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.LocationInfo_r17 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. timeSinceCHO-Reconfig-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.TimeSinceCHO_Reconfig_r17 = new(TimeSinceCHO_Reconfig_r17)
			if err := ie.TimeSinceCHO_Reconfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. shr-Cause-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Shr_Cause_r17 = new(SHR_Cause_r17)
			if err := ie.Shr_Cause_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. ra-InformationCommon-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Ra_InformationCommon_r17 = new(RA_InformationCommon_r16)
			if err := ie.Ra_InformationCommon_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. upInterruptionTimeAtHO-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.UpInterruptionTimeAtHO_r17 = new(UPInterruptionTimeAtHO_r17)
			if err := ie.UpInterruptionTimeAtHO_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. c-RNTI-r17: ref
	{
		if seq.IsComponentPresent(8) {
			ie.C_RNTI_r17 = new(RNTI_Value)
			if err := ie.C_RNTI_r17.Decode(d); err != nil {
				return err
			}
		}
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
				{Name: "targetCell-PCI-ARFCN-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.TargetCell_PCI_ARFCN_r17 = new(PCI_ARFCN_NR_r16)
			if err := ie.TargetCell_PCI_ARFCN_r17.Decode(dx); err != nil {
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
				{Name: "eutra-TargetCellInfo-r18", Optional: true},
				{Name: "measResultServCellRSSI-r18", Optional: true},
				{Name: "measResultNeighFreqListRSSI-r18", Optional: true},
				{Name: "eutra-C-RNTI-r18", Optional: true},
				{Name: "timeSinceSHR-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Eutra_TargetCellInfo_r18 = &struct {
				TargetPCellId_r18 struct {
					Choice           int
					CellGlobalId_r18 *CGI_Info_Logging_r16
					Pci_Arfcn_r18    *PCI_ARFCN_EUTRA_r16
				}
				TargetCellMeas_r18 *MeasQuantityResultsEUTRA
			}{}
			c := ie.Eutra_TargetCellInfo_r18
			successHOReportR17ExtEutraTargetCellInfoR18Seq := dx.NewSequenceDecoder(successHOReportR17ExtEutraTargetCellInfoR18Constraints)
			if err := successHOReportR17ExtEutraTargetCellInfoR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(successHOReportR17ExtEutraTargetCellInfoR18TargetPCellIdR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.TargetPCellId_r18.Choice = int(index)
				switch index {
				case SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_CellGlobalId_r18:
					c.TargetPCellId_r18.CellGlobalId_r18 = new(CGI_Info_Logging_r16)
					if err := c.TargetPCellId_r18.CellGlobalId_r18.Decode(dx); err != nil {
						return err
					}
				case SuccessHO_Report_r17_Ext_Eutra_TargetCellInfo_r18_TargetPCellId_r18_Pci_Arfcn_r18:
					c.TargetPCellId_r18.Pci_Arfcn_r18 = new(PCI_ARFCN_EUTRA_r16)
					if err := c.TargetPCellId_r18.Pci_Arfcn_r18.Decode(dx); err != nil {
						return err
					}
				}
			}
			if successHOReportR17ExtEutraTargetCellInfoR18Seq.IsComponentPresent(1) {
				c.TargetCellMeas_r18 = new(MeasQuantityResultsEUTRA)
				if err := (*c.TargetCellMeas_r18).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MeasResultServCellRSSI_r18 = new(RSSI_Range_r16)
			if err := ie.MeasResultServCellRSSI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MeasResultNeighFreqListRSSI_r18 = new(MeasResultNeighFreqListRSSI_r18)
			if err := ie.MeasResultNeighFreqListRSSI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Eutra_C_RNTI_r18 = new(EUTRA_C_RNTI)
			if err := ie.Eutra_C_RNTI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.TimeSinceSHR_r18 = new(TimeSinceSHR_r18)
			if err := ie.TimeSinceSHR_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sourceCellMeasL1-r19", Optional: true},
				{Name: "targetCellMeasL1-r19", Optional: true},
				{Name: "neighCellsMeasL1ListNR-r19", Optional: true},
				{Name: "rach-Less-r19", Optional: true},
				{Name: "sourcePSCellInfo-r19", Optional: true},
				{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
				{Name: "targetPSCellID-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SourceCellMeasL1_r19 = new(MeasResultL1_r19)
			if err := ie.SourceCellMeasL1_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.TargetCellMeasL1_r19 = new(MeasResultL1_r19)
			if err := ie.TargetCellMeasL1_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.NeighCellsMeasL1ListNR_r19 = new(MeasResultList3NR_r19)
			if err := ie.NeighCellsMeasL1ListNR_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(successHOReportR17ExtRachLessR19Constraints)
			if err != nil {
				return err
			}
			ie.Rach_Less_r19 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.SourcePSCellInfo_r19 = &struct {
				SourcePSCellId_r19   CGI_Info_Logging_r16
				SourcePSCellMeas_r19 *MeasResultSuccessHONR_r17
			}{}
			c := ie.SourcePSCellInfo_r19
			successHOReportR17ExtSourcePSCellInfoR19Seq := dx.NewSequenceDecoder(successHOReportR17ExtSourcePSCellInfoR19Constraints)
			if err := successHOReportR17ExtSourcePSCellInfoR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.SourcePSCellId_r19.Decode(dx); err != nil {
					return err
				}
			}
			if successHOReportR17ExtSourcePSCellInfoR19Seq.IsComponentPresent(1) {
				c.SourcePSCellMeas_r19 = new(MeasResultSuccessHONR_r17)
				if err := (*c.SourcePSCellMeas_r19).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Cho_WithCandidateSCGInfoList_r19 = new(CHO_WithCandidateSCGInfoList_r19)
			if err := ie.Cho_WithCandidateSCGInfoList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.TargetPSCellID_r19 = &struct {
				Choice           int
				CellGlobalId_r19 *CGI_Info_Logging_r16
				Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(successHOReportR17ExtTargetPSCellIDR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.TargetPSCellID_r19).Choice = int(index)
			switch index {
			case SuccessHO_Report_r17_Ext_TargetPSCellID_r19_CellGlobalId_r19:
				(*ie.TargetPSCellID_r19).CellGlobalId_r19 = new(CGI_Info_Logging_r16)
				if err := (*ie.TargetPSCellID_r19).CellGlobalId_r19.Decode(dx); err != nil {
					return err
				}
			case SuccessHO_Report_r17_Ext_TargetPSCellID_r19_Pci_Arfcn_r19:
				(*ie.TargetPSCellID_r19).Pci_Arfcn_r19 = new(PCI_ARFCN_NR_r16)
				if err := (*ie.TargetPSCellID_r19).Pci_Arfcn_r19.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
