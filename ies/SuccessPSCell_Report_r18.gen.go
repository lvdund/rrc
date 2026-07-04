// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SuccessPSCell-Report-r18 (line 3361).

var successPSCellReportR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pCellId-r18"},
		{Name: "sourcePSCellInfo-r18", Optional: true},
		{Name: "targetPSCellInfo-r18"},
		{Name: "measResultNeighCells-r18", Optional: true},
		{Name: "spr-Cause-r18", Optional: true},
		{Name: "timeSinceCPAC-Reconfig-r18", Optional: true},
		{Name: "locationInfo-r18", Optional: true},
		{Name: "ra-InformationCommon-r18", Optional: true},
		{Name: "sn-InitiatedPSCellChange-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SuccessPSCell_Report_r18_Sn_InitiatedPSCellChange_r18_True = 0
)

var successPSCellReportR18SnInitiatedPSCellChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessPSCell_Report_r18_Sn_InitiatedPSCellChange_r18_True},
}

var successPSCellReportR18ExtTargetPCellIDR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r19"},
		{Name: "pci-arfcn-r19"},
	},
}

const (
	SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_CellGlobalId_r19 = 0
	SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_Pci_Arfcn_r19    = 1
)

var successPSCellReportR18SourcePSCellInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sourcePSCellId-r18"},
		{Name: "sourcePSCellMeas-r18", Optional: true},
	},
}

var successPSCellReportR18SourcePSCellInfoR18SourcePSCellIdR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r18"},
		{Name: "pci-arfcn-r18"},
	},
}

const (
	SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_CellGlobalId_r18 = 0
	SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_Pci_Arfcn_r18    = 1
)

var successPSCellReportR18TargetPSCellInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "targetPSCellId-r18"},
		{Name: "targetPSCellMeas-r18", Optional: true},
	},
}

var successPSCellReportR18TargetPSCellInfoR18TargetPSCellIdR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r18"},
		{Name: "pci-arfcn-r18"},
	},
}

const (
	SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_CellGlobalId_r18 = 0
	SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_Pci_Arfcn_r18    = 1
)

var successPSCellReportR18MeasResultNeighCellsR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListNR-r18", Optional: true},
		{Name: "measResultListEUTRA-r18", Optional: true},
	},
}

type SuccessPSCell_Report_r18 struct {
	PCellId_r18          CGI_Info_Logging_r16
	SourcePSCellInfo_r18 *struct {
		SourcePSCellId_r18 struct {
			Choice           int
			CellGlobalId_r18 *CGI_Info_Logging_r16
			Pci_Arfcn_r18    *PCI_ARFCN_EUTRA_r16
		}
		SourcePSCellMeas_r18 *MeasResultSuccessHONR_r17
	}
	TargetPSCellInfo_r18 struct {
		TargetPSCellId_r18 struct {
			Choice           int
			CellGlobalId_r18 *CGI_Info_Logging_r16
			Pci_Arfcn_r18    *PCI_ARFCN_NR_r16
		}
		TargetPSCellMeas_r18 *MeasResultSuccessHONR_r17
	}
	MeasResultNeighCells_r18 *struct {
		MeasResultListNR_r18    *MeasResultList2NR_r16
		MeasResultListEUTRA_r18 *MeasResultList2EUTRA_r16
	}
	Spr_Cause_r18                    *SPR_Cause_r18
	TimeSinceCPAC_Reconfig_r18       *TimeSinceCPAC_Reconfig_r18
	LocationInfo_r18                 *LocationInfo_r16
	Ra_InformationCommon_r18         *RA_InformationCommon_r16
	Sn_InitiatedPSCellChange_r18     *int64
	Cho_WithCandidateSCGInfoList_r19 *CHO_WithCandidateSCGInfoList_r19
	TargetPCellID_r19                *struct {
		Choice           int
		CellGlobalId_r19 *CGI_Info_Logging_r16
		Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
	}
	C_RNTI_r19 *RNTI_Value
}

func (ie *SuccessPSCell_Report_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(successPSCellReportR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Cho_WithCandidateSCGInfoList_r19 != nil || ie.TargetPCellID_r19 != nil || ie.C_RNTI_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SourcePSCellInfo_r18 != nil, ie.MeasResultNeighCells_r18 != nil, ie.Spr_Cause_r18 != nil, ie.TimeSinceCPAC_Reconfig_r18 != nil, ie.LocationInfo_r18 != nil, ie.Ra_InformationCommon_r18 != nil, ie.Sn_InitiatedPSCellChange_r18 != nil}); err != nil {
		return err
	}

	// 3. pCellId-r18: ref
	{
		if err := ie.PCellId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sourcePSCellInfo-r18: sequence
	{
		if ie.SourcePSCellInfo_r18 != nil {
			c := ie.SourcePSCellInfo_r18
			successPSCellReportR18SourcePSCellInfoR18Seq := e.NewSequenceEncoder(successPSCellReportR18SourcePSCellInfoR18Constraints)
			if err := successPSCellReportR18SourcePSCellInfoR18Seq.EncodePreamble([]bool{c.SourcePSCellMeas_r18 != nil}); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(successPSCellReportR18SourcePSCellInfoR18SourcePSCellIdR18Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.SourcePSCellId_r18.Choice), false, nil); err != nil {
					return err
				}
				switch c.SourcePSCellId_r18.Choice {
				case SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_CellGlobalId_r18:
					if err := c.SourcePSCellId_r18.CellGlobalId_r18.Encode(e); err != nil {
						return err
					}
				case SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_Pci_Arfcn_r18:
					if err := c.SourcePSCellId_r18.Pci_Arfcn_r18.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.SourcePSCellMeas_r18 != nil {
				if err := c.SourcePSCellMeas_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. targetPSCellInfo-r18: sequence
	{
		{
			c := &ie.TargetPSCellInfo_r18
			successPSCellReportR18TargetPSCellInfoR18Seq := e.NewSequenceEncoder(successPSCellReportR18TargetPSCellInfoR18Constraints)
			if err := successPSCellReportR18TargetPSCellInfoR18Seq.EncodePreamble([]bool{c.TargetPSCellMeas_r18 != nil}); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(successPSCellReportR18TargetPSCellInfoR18TargetPSCellIdR18Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.TargetPSCellId_r18.Choice), false, nil); err != nil {
					return err
				}
				switch c.TargetPSCellId_r18.Choice {
				case SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_CellGlobalId_r18:
					if err := c.TargetPSCellId_r18.CellGlobalId_r18.Encode(e); err != nil {
						return err
					}
				case SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_Pci_Arfcn_r18:
					if err := c.TargetPSCellId_r18.Pci_Arfcn_r18.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.TargetPSCellMeas_r18 != nil {
				if err := c.TargetPSCellMeas_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. measResultNeighCells-r18: sequence
	{
		if ie.MeasResultNeighCells_r18 != nil {
			c := ie.MeasResultNeighCells_r18
			successPSCellReportR18MeasResultNeighCellsR18Seq := e.NewSequenceEncoder(successPSCellReportR18MeasResultNeighCellsR18Constraints)
			if err := successPSCellReportR18MeasResultNeighCellsR18Seq.EncodePreamble([]bool{c.MeasResultListNR_r18 != nil, c.MeasResultListEUTRA_r18 != nil}); err != nil {
				return err
			}
			if c.MeasResultListNR_r18 != nil {
				if err := c.MeasResultListNR_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.MeasResultListEUTRA_r18 != nil {
				if err := c.MeasResultListEUTRA_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. spr-Cause-r18: ref
	{
		if ie.Spr_Cause_r18 != nil {
			if err := ie.Spr_Cause_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. timeSinceCPAC-Reconfig-r18: ref
	{
		if ie.TimeSinceCPAC_Reconfig_r18 != nil {
			if err := ie.TimeSinceCPAC_Reconfig_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. locationInfo-r18: ref
	{
		if ie.LocationInfo_r18 != nil {
			if err := ie.LocationInfo_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ra-InformationCommon-r18: ref
	{
		if ie.Ra_InformationCommon_r18 != nil {
			if err := ie.Ra_InformationCommon_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. sn-InitiatedPSCellChange-r18: enumerated
	{
		if ie.Sn_InitiatedPSCellChange_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sn_InitiatedPSCellChange_r18, successPSCellReportR18SnInitiatedPSCellChangeR18Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
					{Name: "targetPCellID-r19", Optional: true},
					{Name: "c-RNTI-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cho_WithCandidateSCGInfoList_r19 != nil, ie.TargetPCellID_r19 != nil, ie.C_RNTI_r19 != nil}); err != nil {
				return err
			}

			if ie.Cho_WithCandidateSCGInfoList_r19 != nil {
				if err := ie.Cho_WithCandidateSCGInfoList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TargetPCellID_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(successPSCellReportR18ExtTargetPCellIDR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.TargetPCellID_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.TargetPCellID_r19).Choice {
				case SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_CellGlobalId_r19:
					if err := (*ie.TargetPCellID_r19).CellGlobalId_r19.Encode(ex); err != nil {
						return err
					}
				case SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_Pci_Arfcn_r19:
					if err := (*ie.TargetPCellID_r19).Pci_Arfcn_r19.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.C_RNTI_r19 != nil {
				if err := ie.C_RNTI_r19.Encode(ex); err != nil {
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

func (ie *SuccessPSCell_Report_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(successPSCellReportR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pCellId-r18: ref
	{
		if err := ie.PCellId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sourcePSCellInfo-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.SourcePSCellInfo_r18 = &struct {
				SourcePSCellId_r18 struct {
					Choice           int
					CellGlobalId_r18 *CGI_Info_Logging_r16
					Pci_Arfcn_r18    *PCI_ARFCN_EUTRA_r16
				}
				SourcePSCellMeas_r18 *MeasResultSuccessHONR_r17
			}{}
			c := ie.SourcePSCellInfo_r18
			successPSCellReportR18SourcePSCellInfoR18Seq := d.NewSequenceDecoder(successPSCellReportR18SourcePSCellInfoR18Constraints)
			if err := successPSCellReportR18SourcePSCellInfoR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := d.NewChoiceDecoder(successPSCellReportR18SourcePSCellInfoR18SourcePSCellIdR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.SourcePSCellId_r18.Choice = int(index)
				switch index {
				case SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_CellGlobalId_r18:
					c.SourcePSCellId_r18.CellGlobalId_r18 = new(CGI_Info_Logging_r16)
					if err := c.SourcePSCellId_r18.CellGlobalId_r18.Decode(d); err != nil {
						return err
					}
				case SuccessPSCell_Report_r18_SourcePSCellInfo_r18_SourcePSCellId_r18_Pci_Arfcn_r18:
					c.SourcePSCellId_r18.Pci_Arfcn_r18 = new(PCI_ARFCN_EUTRA_r16)
					if err := c.SourcePSCellId_r18.Pci_Arfcn_r18.Decode(d); err != nil {
						return err
					}
				}
			}
			if successPSCellReportR18SourcePSCellInfoR18Seq.IsComponentPresent(1) {
				c.SourcePSCellMeas_r18 = new(MeasResultSuccessHONR_r17)
				if err := (*c.SourcePSCellMeas_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. targetPSCellInfo-r18: sequence
	{
		{
			c := &ie.TargetPSCellInfo_r18
			successPSCellReportR18TargetPSCellInfoR18Seq := d.NewSequenceDecoder(successPSCellReportR18TargetPSCellInfoR18Constraints)
			if err := successPSCellReportR18TargetPSCellInfoR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := d.NewChoiceDecoder(successPSCellReportR18TargetPSCellInfoR18TargetPSCellIdR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.TargetPSCellId_r18.Choice = int(index)
				switch index {
				case SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_CellGlobalId_r18:
					c.TargetPSCellId_r18.CellGlobalId_r18 = new(CGI_Info_Logging_r16)
					if err := c.TargetPSCellId_r18.CellGlobalId_r18.Decode(d); err != nil {
						return err
					}
				case SuccessPSCell_Report_r18_TargetPSCellInfo_r18_TargetPSCellId_r18_Pci_Arfcn_r18:
					c.TargetPSCellId_r18.Pci_Arfcn_r18 = new(PCI_ARFCN_NR_r16)
					if err := c.TargetPSCellId_r18.Pci_Arfcn_r18.Decode(d); err != nil {
						return err
					}
				}
			}
			if successPSCellReportR18TargetPSCellInfoR18Seq.IsComponentPresent(1) {
				c.TargetPSCellMeas_r18 = new(MeasResultSuccessHONR_r17)
				if err := (*c.TargetPSCellMeas_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. measResultNeighCells-r18: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.MeasResultNeighCells_r18 = &struct {
				MeasResultListNR_r18    *MeasResultList2NR_r16
				MeasResultListEUTRA_r18 *MeasResultList2EUTRA_r16
			}{}
			c := ie.MeasResultNeighCells_r18
			successPSCellReportR18MeasResultNeighCellsR18Seq := d.NewSequenceDecoder(successPSCellReportR18MeasResultNeighCellsR18Constraints)
			if err := successPSCellReportR18MeasResultNeighCellsR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if successPSCellReportR18MeasResultNeighCellsR18Seq.IsComponentPresent(0) {
				c.MeasResultListNR_r18 = new(MeasResultList2NR_r16)
				if err := (*c.MeasResultListNR_r18).Decode(d); err != nil {
					return err
				}
			}
			if successPSCellReportR18MeasResultNeighCellsR18Seq.IsComponentPresent(1) {
				c.MeasResultListEUTRA_r18 = new(MeasResultList2EUTRA_r16)
				if err := (*c.MeasResultListEUTRA_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. spr-Cause-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Spr_Cause_r18 = new(SPR_Cause_r18)
			if err := ie.Spr_Cause_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. timeSinceCPAC-Reconfig-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.TimeSinceCPAC_Reconfig_r18 = new(TimeSinceCPAC_Reconfig_r18)
			if err := ie.TimeSinceCPAC_Reconfig_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. locationInfo-r18: ref
	{
		if seq.IsComponentPresent(6) {
			ie.LocationInfo_r18 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ra-InformationCommon-r18: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ra_InformationCommon_r18 = new(RA_InformationCommon_r16)
			if err := ie.Ra_InformationCommon_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. sn-InitiatedPSCellChange-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(successPSCellReportR18SnInitiatedPSCellChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.Sn_InitiatedPSCellChange_r18 = &idx
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
				{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
				{Name: "targetPCellID-r19", Optional: true},
				{Name: "c-RNTI-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cho_WithCandidateSCGInfoList_r19 = new(CHO_WithCandidateSCGInfoList_r19)
			if err := ie.Cho_WithCandidateSCGInfoList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.TargetPCellID_r19 = &struct {
				Choice           int
				CellGlobalId_r19 *CGI_Info_Logging_r16
				Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(successPSCellReportR18ExtTargetPCellIDR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.TargetPCellID_r19).Choice = int(index)
			switch index {
			case SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_CellGlobalId_r19:
				(*ie.TargetPCellID_r19).CellGlobalId_r19 = new(CGI_Info_Logging_r16)
				if err := (*ie.TargetPCellID_r19).CellGlobalId_r19.Decode(dx); err != nil {
					return err
				}
			case SuccessPSCell_Report_r18_Ext_TargetPCellID_r19_Pci_Arfcn_r19:
				(*ie.TargetPCellID_r19).Pci_Arfcn_r19 = new(PCI_ARFCN_NR_r16)
				if err := (*ie.TargetPCellID_r19).Pci_Arfcn_r19.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.C_RNTI_r19 = new(RNTI_Value)
			if err := ie.C_RNTI_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
