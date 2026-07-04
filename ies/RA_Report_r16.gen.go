// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RA-Report-r16 (line 3061).

var rAReportR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellId-r16"},
		{Name: "ra-InformationCommon-r16", Optional: true},
		{Name: "raPurpose-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var rA_Report_r16CellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r16"},
		{Name: "pci-arfcn-r16"},
	},
}

const (
	RA_Report_r16_CellId_r16_CellGlobalId_r16 = 0
	RA_Report_r16_CellId_r16_Pci_Arfcn_r16    = 1
)

const (
	RA_Report_r16_RaPurpose_r16_AccessRelated             = 0
	RA_Report_r16_RaPurpose_r16_BeamFailureRecovery       = 1
	RA_Report_r16_RaPurpose_r16_ReconfigurationWithSync   = 2
	RA_Report_r16_RaPurpose_r16_UlUnSynchronized          = 3
	RA_Report_r16_RaPurpose_r16_SchedulingRequestFailure  = 4
	RA_Report_r16_RaPurpose_r16_NoPUCCHResourceAvailable  = 5
	RA_Report_r16_RaPurpose_r16_RequestForOtherSI         = 6
	RA_Report_r16_RaPurpose_r16_Msg3RequestForOtherSI_r17 = 7
	RA_Report_r16_RaPurpose_r16_Lbt_Failure_r18           = 8
	RA_Report_r16_RaPurpose_r16_Ltm_r19                   = 9
	RA_Report_r16_RaPurpose_r16_Spare6                    = 10
	RA_Report_r16_RaPurpose_r16_Spare5                    = 11
	RA_Report_r16_RaPurpose_r16_Spare4                    = 12
	RA_Report_r16_RaPurpose_r16_Spare3                    = 13
	RA_Report_r16_RaPurpose_r16_Spare2                    = 14
	RA_Report_r16_RaPurpose_r16_Spare1                    = 15
)

var rAReportR16RaPurposeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_Report_r16_RaPurpose_r16_AccessRelated, RA_Report_r16_RaPurpose_r16_BeamFailureRecovery, RA_Report_r16_RaPurpose_r16_ReconfigurationWithSync, RA_Report_r16_RaPurpose_r16_UlUnSynchronized, RA_Report_r16_RaPurpose_r16_SchedulingRequestFailure, RA_Report_r16_RaPurpose_r16_NoPUCCHResourceAvailable, RA_Report_r16_RaPurpose_r16_RequestForOtherSI, RA_Report_r16_RaPurpose_r16_Msg3RequestForOtherSI_r17, RA_Report_r16_RaPurpose_r16_Lbt_Failure_r18, RA_Report_r16_RaPurpose_r16_Ltm_r19, RA_Report_r16_RaPurpose_r16_Spare6, RA_Report_r16_RaPurpose_r16_Spare5, RA_Report_r16_RaPurpose_r16_Spare4, RA_Report_r16_RaPurpose_r16_Spare3, RA_Report_r16_RaPurpose_r16_Spare2, RA_Report_r16_RaPurpose_r16_Spare1},
}

const (
	RA_Report_r16_Ext_Sdt_Failed_r18_True = 0
)

var rAReportR16ExtSdtFailedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_Report_r16_Ext_Sdt_Failed_r18_True},
}

const (
	RA_Report_r16_Ext_Sdt_FailureCause_r19_T319a_Expiry              = 0
	RA_Report_r16_Ext_Sdt_FailureCause_r19_MaxRetxThreshold          = 1
	RA_Report_r16_Ext_Sdt_FailureCause_r19_PreambleTransMax          = 2
	RA_Report_r16_Ext_Sdt_FailureCause_r19_ConfiguredGrantTimer      = 3
	RA_Report_r16_Ext_Sdt_FailureCause_r19_Cg_SDT_TimeAlignmentTimer = 4
	RA_Report_r16_Ext_Sdt_FailureCause_r19_CellReselection           = 5
	RA_Report_r16_Ext_Sdt_FailureCause_r19_Spare2                    = 6
	RA_Report_r16_Ext_Sdt_FailureCause_r19_Spare1                    = 7
)

var rAReportR16ExtSdtFailureCauseR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_Report_r16_Ext_Sdt_FailureCause_r19_T319a_Expiry, RA_Report_r16_Ext_Sdt_FailureCause_r19_MaxRetxThreshold, RA_Report_r16_Ext_Sdt_FailureCause_r19_PreambleTransMax, RA_Report_r16_Ext_Sdt_FailureCause_r19_ConfiguredGrantTimer, RA_Report_r16_Ext_Sdt_FailureCause_r19_Cg_SDT_TimeAlignmentTimer, RA_Report_r16_Ext_Sdt_FailureCause_r19_CellReselection, RA_Report_r16_Ext_Sdt_FailureCause_r19_Spare2, RA_Report_r16_Ext_Sdt_FailureCause_r19_Spare1},
}

var rAReportR16SdtULDataVolumeR19Constraints = per.Constrained(0, 96000)

type RA_Report_r16 struct {
	CellId_r16 struct {
		Choice           int
		CellGlobalId_r16 *CGI_Info_Logging_r16
		Pci_Arfcn_r16    *PCI_ARFCN_NR_r16
	}
	Ra_InformationCommon_r16   *RA_InformationCommon_r16
	RaPurpose_r16              int64
	SpCellID_r17               *CGI_Info_Logging_r16
	Sdt_Failed_r18             *int64
	Sdt_FailureCause_r19       *int64
	Sdt_DL_Rsrp_Info_r19       *RSRP_Range
	Sdt_UL_DataVolume_r19      *int64
	TimeSinceSdt_Execution_r19 *TimeSinceSdt_Execution_r19
}

func (ie *RA_Report_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAReportR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SpCellID_r17 != nil
	hasExtGroup1 := ie.Sdt_Failed_r18 != nil
	hasExtGroup2 := ie.Sdt_FailureCause_r19 != nil || ie.Sdt_DL_Rsrp_Info_r19 != nil || ie.Sdt_UL_DataVolume_r19 != nil || ie.TimeSinceSdt_Execution_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ra_InformationCommon_r16 != nil}); err != nil {
		return err
	}

	// 3. cellId-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(rA_Report_r16CellIdR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CellId_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CellId_r16.Choice {
		case RA_Report_r16_CellId_r16_CellGlobalId_r16:
			if err := ie.CellId_r16.CellGlobalId_r16.Encode(e); err != nil {
				return err
			}
		case RA_Report_r16_CellId_r16_Pci_Arfcn_r16:
			if err := ie.CellId_r16.Pci_Arfcn_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CellId_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 4. ra-InformationCommon-r16: ref
	{
		if ie.Ra_InformationCommon_r16 != nil {
			if err := ie.Ra_InformationCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. raPurpose-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.RaPurpose_r16, rAReportR16RaPurposeR16Constraints); err != nil {
			return err
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
					{Name: "spCellID-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpCellID_r17 != nil}); err != nil {
				return err
			}

			if ie.SpCellID_r17 != nil {
				if err := ie.SpCellID_r17.Encode(ex); err != nil {
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
					{Name: "sdt-Failed-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sdt_Failed_r18 != nil}); err != nil {
				return err
			}

			if ie.Sdt_Failed_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sdt_Failed_r18, rAReportR16ExtSdtFailedR18Constraints); err != nil {
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
					{Name: "sdt-FailureCause-r19", Optional: true},
					{Name: "sdt-DL-Rsrp-Info-r19", Optional: true},
					{Name: "sdt-UL-DataVolume-r19", Optional: true},
					{Name: "timeSinceSdt-Execution-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sdt_FailureCause_r19 != nil, ie.Sdt_DL_Rsrp_Info_r19 != nil, ie.Sdt_UL_DataVolume_r19 != nil, ie.TimeSinceSdt_Execution_r19 != nil}); err != nil {
				return err
			}

			if ie.Sdt_FailureCause_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sdt_FailureCause_r19, rAReportR16ExtSdtFailureCauseR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sdt_DL_Rsrp_Info_r19 != nil {
				if err := ie.Sdt_DL_Rsrp_Info_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sdt_UL_DataVolume_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sdt_UL_DataVolume_r19, rAReportR16SdtULDataVolumeR19Constraints); err != nil {
					return err
				}
			}

			if ie.TimeSinceSdt_Execution_r19 != nil {
				if err := ie.TimeSinceSdt_Execution_r19.Encode(ex); err != nil {
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

func (ie *RA_Report_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAReportR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cellId-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(rA_Report_r16CellIdR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CellId_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RA_Report_r16_CellId_r16_CellGlobalId_r16:
			ie.CellId_r16.CellGlobalId_r16 = new(CGI_Info_Logging_r16)
			if err := ie.CellId_r16.CellGlobalId_r16.Decode(d); err != nil {
				return err
			}
		case RA_Report_r16_CellId_r16_Pci_Arfcn_r16:
			ie.CellId_r16.Pci_Arfcn_r16 = new(PCI_ARFCN_NR_r16)
			if err := ie.CellId_r16.Pci_Arfcn_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ra-InformationCommon-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ra_InformationCommon_r16 = new(RA_InformationCommon_r16)
			if err := ie.Ra_InformationCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. raPurpose-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(rAReportR16RaPurposeR16Constraints)
		if err != nil {
			return err
		}
		ie.RaPurpose_r16 = v2
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
				{Name: "spCellID-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SpCellID_r17 = new(CGI_Info_Logging_r16)
			if err := ie.SpCellID_r17.Decode(dx); err != nil {
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
				{Name: "sdt-Failed-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rAReportR16ExtSdtFailedR18Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_Failed_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sdt-FailureCause-r19", Optional: true},
				{Name: "sdt-DL-Rsrp-Info-r19", Optional: true},
				{Name: "sdt-UL-DataVolume-r19", Optional: true},
				{Name: "timeSinceSdt-Execution-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rAReportR16ExtSdtFailureCauseR19Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_FailureCause_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sdt_DL_Rsrp_Info_r19 = new(RSRP_Range)
			if err := ie.Sdt_DL_Rsrp_Info_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(rAReportR16SdtULDataVolumeR19Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_UL_DataVolume_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.TimeSinceSdt_Execution_r19 = new(TimeSinceSdt_Execution_r19)
			if err := ie.TimeSinceSdt_Execution_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
