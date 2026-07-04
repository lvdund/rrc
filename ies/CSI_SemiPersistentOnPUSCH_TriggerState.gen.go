// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-SemiPersistentOnPUSCH-TriggerState (line 7599).

var cSISemiPersistentOnPUSCHTriggerStateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "associatedReportConfigInfo"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	CSI_SemiPersistentOnPUSCH_TriggerState_Ext_Sp_CSI_MultiplexingMode_r17_Enabled = 0
)

var cSISemiPersistentOnPUSCHTriggerStateExtSpCSIMultiplexingModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_SemiPersistentOnPUSCH_TriggerState_Ext_Sp_CSI_MultiplexingMode_r17_Enabled},
}

type CSI_SemiPersistentOnPUSCH_TriggerState struct {
	AssociatedReportConfigInfo         CSI_ReportConfigId
	Sp_CSI_MultiplexingMode_r17        *int64
	Csi_ReportSubConfigTriggerList_r18 *CSI_ReportSubConfigTriggerList_r18
	Ltm_AssociatedReportConfigInfo_r18 *LTM_CSI_ReportConfigId_r18
}

func (ie *CSI_SemiPersistentOnPUSCH_TriggerState) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSISemiPersistentOnPUSCHTriggerStateConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sp_CSI_MultiplexingMode_r17 != nil
	hasExtGroup1 := ie.Csi_ReportSubConfigTriggerList_r18 != nil || ie.Ltm_AssociatedReportConfigInfo_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. associatedReportConfigInfo: ref
	{
		if err := ie.AssociatedReportConfigInfo.Encode(e); err != nil {
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
					{Name: "sp-CSI-MultiplexingMode-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sp_CSI_MultiplexingMode_r17 != nil}); err != nil {
				return err
			}

			if ie.Sp_CSI_MultiplexingMode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sp_CSI_MultiplexingMode_r17, cSISemiPersistentOnPUSCHTriggerStateExtSpCSIMultiplexingModeR17Constraints); err != nil {
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
					{Name: "csi-ReportSubConfigTriggerList-r18", Optional: true},
					{Name: "ltm-AssociatedReportConfigInfo-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Csi_ReportSubConfigTriggerList_r18 != nil, ie.Ltm_AssociatedReportConfigInfo_r18 != nil}); err != nil {
				return err
			}

			if ie.Csi_ReportSubConfigTriggerList_r18 != nil {
				if err := ie.Csi_ReportSubConfigTriggerList_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_AssociatedReportConfigInfo_r18 != nil {
				if err := ie.Ltm_AssociatedReportConfigInfo_r18.Encode(ex); err != nil {
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

func (ie *CSI_SemiPersistentOnPUSCH_TriggerState) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSISemiPersistentOnPUSCHTriggerStateConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. associatedReportConfigInfo: ref
	{
		if err := ie.AssociatedReportConfigInfo.Decode(d); err != nil {
			return err
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
				{Name: "sp-CSI-MultiplexingMode-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cSISemiPersistentOnPUSCHTriggerStateExtSpCSIMultiplexingModeR17Constraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_MultiplexingMode_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "csi-ReportSubConfigTriggerList-r18", Optional: true},
				{Name: "ltm-AssociatedReportConfigInfo-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Csi_ReportSubConfigTriggerList_r18 = new(CSI_ReportSubConfigTriggerList_r18)
			if err := ie.Csi_ReportSubConfigTriggerList_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_AssociatedReportConfigInfo_r18 = new(LTM_CSI_ReportConfigId_r18)
			if err := ie.Ltm_AssociatedReportConfigInfo_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
