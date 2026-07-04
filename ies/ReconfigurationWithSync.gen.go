// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReconfigurationWithSync (line 5682).

var reconfigurationWithSyncConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "spCellConfigCommon", Optional: true},
		{Name: "newUE-Identity"},
		{Name: "t304"},
		{Name: "rach-ConfigDedicated", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	ReconfigurationWithSync_T304_Ms50    = 0
	ReconfigurationWithSync_T304_Ms100   = 1
	ReconfigurationWithSync_T304_Ms150   = 2
	ReconfigurationWithSync_T304_Ms200   = 3
	ReconfigurationWithSync_T304_Ms500   = 4
	ReconfigurationWithSync_T304_Ms1000  = 5
	ReconfigurationWithSync_T304_Ms2000  = 6
	ReconfigurationWithSync_T304_Ms10000 = 7
)

var reconfigurationWithSyncT304Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReconfigurationWithSync_T304_Ms50, ReconfigurationWithSync_T304_Ms100, ReconfigurationWithSync_T304_Ms150, ReconfigurationWithSync_T304_Ms200, ReconfigurationWithSync_T304_Ms500, ReconfigurationWithSync_T304_Ms1000, ReconfigurationWithSync_T304_Ms2000, ReconfigurationWithSync_T304_Ms10000},
}

var reconfigurationWithSyncRachConfigDedicatedConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "uplink"},
		{Name: "supplementaryUplink"},
	},
}

const (
	ReconfigurationWithSync_Rach_ConfigDedicated_Uplink              = 0
	ReconfigurationWithSync_Rach_ConfigDedicated_SupplementaryUplink = 1
)

const (
	ReconfigurationWithSync_Ext_Sl_IndirectPathMaintain_r18_True = 0
)

var reconfigurationWithSyncExtSlIndirectPathMaintainR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReconfigurationWithSync_Ext_Sl_IndirectPathMaintain_r18_True},
}

type ReconfigurationWithSync struct {
	SpCellConfigCommon   *ServingCellConfigCommon
	NewUE_Identity       RNTI_Value
	T304                 int64
	Rach_ConfigDedicated *struct {
		Choice              int
		Uplink              *RACH_ConfigDedicated
		SupplementaryUplink *RACH_ConfigDedicated
	}
	Smtc                          *SSB_MTC
	Daps_UplinkPowerConfig_r16    *DAPS_UplinkPowerConfig_r16
	Sl_PathSwitchConfig_r17       *SL_PathSwitchConfig_r17
	Rach_LessHO_r18               *RACH_LessHO_r18
	Sl_IndirectPathMaintain_r18   *int64
	Ltm_SchedulingRequestInfo_r19 *LTM_SchedulingRequestInfo_r19
	EarlyCSI_Acquisition_r19      *EarlyCSI_Acquisition_r19
}

func (ie *ReconfigurationWithSync) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reconfigurationWithSyncConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Smtc != nil
	hasExtGroup1 := ie.Daps_UplinkPowerConfig_r16 != nil
	hasExtGroup2 := ie.Sl_PathSwitchConfig_r17 != nil
	hasExtGroup3 := ie.Rach_LessHO_r18 != nil || ie.Sl_IndirectPathMaintain_r18 != nil
	hasExtGroup4 := ie.Ltm_SchedulingRequestInfo_r19 != nil || ie.EarlyCSI_Acquisition_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SpCellConfigCommon != nil, ie.Rach_ConfigDedicated != nil}); err != nil {
		return err
	}

	// 3. spCellConfigCommon: ref
	{
		if ie.SpCellConfigCommon != nil {
			if err := ie.SpCellConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. newUE-Identity: ref
	{
		if err := ie.NewUE_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 5. t304: enumerated
	{
		if err := e.EncodeEnumerated(ie.T304, reconfigurationWithSyncT304Constraints); err != nil {
			return err
		}
	}

	// 6. rach-ConfigDedicated: choice
	{
		if ie.Rach_ConfigDedicated != nil {
			choiceEnc := e.NewChoiceEncoder(reconfigurationWithSyncRachConfigDedicatedConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rach_ConfigDedicated).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rach_ConfigDedicated).Choice {
			case ReconfigurationWithSync_Rach_ConfigDedicated_Uplink:
				if err := (*ie.Rach_ConfigDedicated).Uplink.Encode(e); err != nil {
					return err
				}
			case ReconfigurationWithSync_Rach_ConfigDedicated_SupplementaryUplink:
				if err := (*ie.Rach_ConfigDedicated).SupplementaryUplink.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rach_ConfigDedicated).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "smtc", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Smtc != nil}); err != nil {
				return err
			}

			if ie.Smtc != nil {
				if err := ie.Smtc.Encode(ex); err != nil {
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
					{Name: "daps-UplinkPowerConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Daps_UplinkPowerConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.Daps_UplinkPowerConfig_r16 != nil {
				if err := ie.Daps_UplinkPowerConfig_r16.Encode(ex); err != nil {
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
					{Name: "sl-PathSwitchConfig-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PathSwitchConfig_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_PathSwitchConfig_r17 != nil {
				if err := ie.Sl_PathSwitchConfig_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rach-LessHO-r18", Optional: true},
					{Name: "sl-IndirectPathMaintain-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rach_LessHO_r18 != nil, ie.Sl_IndirectPathMaintain_r18 != nil}); err != nil {
				return err
			}

			if ie.Rach_LessHO_r18 != nil {
				if err := ie.Rach_LessHO_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_IndirectPathMaintain_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_IndirectPathMaintain_r18, reconfigurationWithSyncExtSlIndirectPathMaintainR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-SchedulingRequestInfo-r19", Optional: true},
					{Name: "earlyCSI-Acquisition-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_SchedulingRequestInfo_r19 != nil, ie.EarlyCSI_Acquisition_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_SchedulingRequestInfo_r19 != nil {
				if err := ie.Ltm_SchedulingRequestInfo_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.EarlyCSI_Acquisition_r19 != nil {
				if err := ie.EarlyCSI_Acquisition_r19.Encode(ex); err != nil {
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

func (ie *ReconfigurationWithSync) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reconfigurationWithSyncConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. spCellConfigCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SpCellConfigCommon = new(ServingCellConfigCommon)
			if err := ie.SpCellConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. newUE-Identity: ref
	{
		if err := ie.NewUE_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 5. t304: enumerated
	{
		v2, err := d.DecodeEnumerated(reconfigurationWithSyncT304Constraints)
		if err != nil {
			return err
		}
		ie.T304 = v2
	}

	// 6. rach-ConfigDedicated: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Rach_ConfigDedicated = &struct {
				Choice              int
				Uplink              *RACH_ConfigDedicated
				SupplementaryUplink *RACH_ConfigDedicated
			}{}
			choiceDec := d.NewChoiceDecoder(reconfigurationWithSyncRachConfigDedicatedConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rach_ConfigDedicated).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ReconfigurationWithSync_Rach_ConfigDedicated_Uplink:
				(*ie.Rach_ConfigDedicated).Uplink = new(RACH_ConfigDedicated)
				if err := (*ie.Rach_ConfigDedicated).Uplink.Decode(d); err != nil {
					return err
				}
			case ReconfigurationWithSync_Rach_ConfigDedicated_SupplementaryUplink:
				(*ie.Rach_ConfigDedicated).SupplementaryUplink = new(RACH_ConfigDedicated)
				if err := (*ie.Rach_ConfigDedicated).SupplementaryUplink.Decode(d); err != nil {
					return err
				}
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
				{Name: "smtc", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Smtc = new(SSB_MTC)
			if err := ie.Smtc.Decode(dx); err != nil {
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
				{Name: "daps-UplinkPowerConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Daps_UplinkPowerConfig_r16 = new(DAPS_UplinkPowerConfig_r16)
			if err := ie.Daps_UplinkPowerConfig_r16.Decode(dx); err != nil {
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
				{Name: "sl-PathSwitchConfig-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_PathSwitchConfig_r17 = new(SL_PathSwitchConfig_r17)
			if err := ie.Sl_PathSwitchConfig_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "rach-LessHO-r18", Optional: true},
				{Name: "sl-IndirectPathMaintain-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rach_LessHO_r18 = new(RACH_LessHO_r18)
			if err := ie.Rach_LessHO_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(reconfigurationWithSyncExtSlIndirectPathMaintainR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_IndirectPathMaintain_r18 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-SchedulingRequestInfo-r19", Optional: true},
				{Name: "earlyCSI-Acquisition-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_SchedulingRequestInfo_r19 = new(LTM_SchedulingRequestInfo_r19)
			if err := ie.Ltm_SchedulingRequestInfo_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.EarlyCSI_Acquisition_r19 = new(EarlyCSI_Acquisition_r19)
			if err := ie.EarlyCSI_Acquisition_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
