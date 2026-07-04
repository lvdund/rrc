// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-AperiodicTriggerState (line 6846).

var cSIAperiodicTriggerStateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "associatedReportConfigInfoList"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var cSIAperiodicTriggerStateAssociatedReportConfigInfoListConstraints = per.SizeRange(1, common.MaxNrofReportConfigPerAperiodicTrigger)

const (
	CSI_AperiodicTriggerState_Ext_Ap_CSI_MultiplexingMode_r17_Enabled = 0
)

var cSIAperiodicTriggerStateExtApCSIMultiplexingModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AperiodicTriggerState_Ext_Ap_CSI_MultiplexingMode_r17_Enabled},
}

const (
	CSI_AperiodicTriggerState_Ext_DelayOffsetCompensation_r19_Enabled = 0
)

var cSIAperiodicTriggerStateExtDelayOffsetCompensationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AperiodicTriggerState_Ext_DelayOffsetCompensation_r19_Enabled},
}

type CSI_AperiodicTriggerState struct {
	AssociatedReportConfigInfoList     []CSI_AssociatedReportConfigInfo
	Ap_CSI_MultiplexingMode_r17        *int64
	Ltm_AssociatedReportConfigInfo_r18 *LTM_CSI_ReportConfigId_r18
	DelayOffsetCompensation_r19        *int64
}

func (ie *CSI_AperiodicTriggerState) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIAperiodicTriggerStateConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ap_CSI_MultiplexingMode_r17 != nil
	hasExtGroup1 := ie.Ltm_AssociatedReportConfigInfo_r18 != nil
	hasExtGroup2 := ie.DelayOffsetCompensation_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. associatedReportConfigInfoList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSIAperiodicTriggerStateAssociatedReportConfigInfoListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.AssociatedReportConfigInfoList))); err != nil {
			return err
		}
		for i := range ie.AssociatedReportConfigInfoList {
			if err := ie.AssociatedReportConfigInfoList[i].Encode(e); err != nil {
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
					{Name: "ap-CSI-MultiplexingMode-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ap_CSI_MultiplexingMode_r17 != nil}); err != nil {
				return err
			}

			if ie.Ap_CSI_MultiplexingMode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ap_CSI_MultiplexingMode_r17, cSIAperiodicTriggerStateExtApCSIMultiplexingModeR17Constraints); err != nil {
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
					{Name: "ltm-AssociatedReportConfigInfo-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_AssociatedReportConfigInfo_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_AssociatedReportConfigInfo_r18 != nil {
				if err := ie.Ltm_AssociatedReportConfigInfo_r18.Encode(ex); err != nil {
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
					{Name: "delayOffsetCompensation-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DelayOffsetCompensation_r19 != nil}); err != nil {
				return err
			}

			if ie.DelayOffsetCompensation_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.DelayOffsetCompensation_r19, cSIAperiodicTriggerStateExtDelayOffsetCompensationR19Constraints); err != nil {
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

func (ie *CSI_AperiodicTriggerState) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIAperiodicTriggerStateConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. associatedReportConfigInfoList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSIAperiodicTriggerStateAssociatedReportConfigInfoListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.AssociatedReportConfigInfoList = make([]CSI_AssociatedReportConfigInfo, n)
		for i := int64(0); i < n; i++ {
			if err := ie.AssociatedReportConfigInfoList[i].Decode(d); err != nil {
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
				{Name: "ap-CSI-MultiplexingMode-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cSIAperiodicTriggerStateExtApCSIMultiplexingModeR17Constraints)
			if err != nil {
				return err
			}
			ie.Ap_CSI_MultiplexingMode_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-AssociatedReportConfigInfo-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_AssociatedReportConfigInfo_r18 = new(LTM_CSI_ReportConfigId_r18)
			if err := ie.Ltm_AssociatedReportConfigInfo_r18.Decode(dx); err != nil {
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
				{Name: "delayOffsetCompensation-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cSIAperiodicTriggerStateExtDelayOffsetCompensationR19Constraints)
			if err != nil {
				return err
			}
			ie.DelayOffsetCompensation_r19 = &v
		}
	}

	return nil
}
