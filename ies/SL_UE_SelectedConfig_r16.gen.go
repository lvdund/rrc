// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-UE-SelectedConfig-r16 (line 28358).

var sLUESelectedConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PSSCH-TxConfigList-r16", Optional: true},
		{Name: "sl-ProbResourceKeep-r16", Optional: true},
		{Name: "sl-ReselectAfter-r16", Optional: true},
		{Name: "sl-CBR-CommonTxConfigList-r16", Optional: true},
		{Name: "ul-PrioritizationThres-r16", Optional: true},
		{Name: "sl-PrioritizationThres-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_v0     = 0
	SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot2 = 1
	SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot4 = 2
	SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot6 = 3
	SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot8 = 4
)

var sLUESelectedConfigR16SlProbResourceKeepR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_v0, SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot2, SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot4, SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot6, SL_UE_SelectedConfig_r16_Sl_ProbResourceKeep_r16_V0dot8},
}

const (
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N1 = 0
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N2 = 1
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N3 = 2
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N4 = 3
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N5 = 4
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N6 = 5
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N7 = 6
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N8 = 7
	SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N9 = 8
)

var sLUESelectedConfigR16SlReselectAfterR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N1, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N2, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N3, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N4, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N5, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N6, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N7, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N8, SL_UE_SelectedConfig_r16_Sl_ReselectAfter_r16_N9},
}

var sLUESelectedConfigR16UlPrioritizationThresR16Constraints = per.Constrained(1, 16)

var sLUESelectedConfigR16SlPrioritizationThresR16Constraints = per.Constrained(1, 8)

type SL_UE_SelectedConfig_r16 struct {
	Sl_PSSCH_TxConfigList_r16                  *SL_PSSCH_TxConfigList_r16
	Sl_ProbResourceKeep_r16                    *int64
	Sl_ReselectAfter_r16                       *int64
	Sl_CBR_CommonTxConfigList_r16              *SL_CBR_CommonTxConfigList_r16
	Ul_PrioritizationThres_r16                 *int64
	Sl_PrioritizationThres_r16                 *int64
	Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 *SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18
}

func (ie *SL_UE_SelectedConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLUESelectedConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PSSCH_TxConfigList_r16 != nil, ie.Sl_ProbResourceKeep_r16 != nil, ie.Sl_ReselectAfter_r16 != nil, ie.Sl_CBR_CommonTxConfigList_r16 != nil, ie.Ul_PrioritizationThres_r16 != nil, ie.Sl_PrioritizationThres_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PSSCH-TxConfigList-r16: ref
	{
		if ie.Sl_PSSCH_TxConfigList_r16 != nil {
			if err := ie.Sl_PSSCH_TxConfigList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-ProbResourceKeep-r16: enumerated
	{
		if ie.Sl_ProbResourceKeep_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ProbResourceKeep_r16, sLUESelectedConfigR16SlProbResourceKeepR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-ReselectAfter-r16: enumerated
	{
		if ie.Sl_ReselectAfter_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ReselectAfter_r16, sLUESelectedConfigR16SlReselectAfterR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-CBR-CommonTxConfigList-r16: ref
	{
		if ie.Sl_CBR_CommonTxConfigList_r16 != nil {
			if err := ie.Sl_CBR_CommonTxConfigList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. ul-PrioritizationThres-r16: integer
	{
		if ie.Ul_PrioritizationThres_r16 != nil {
			if err := e.EncodeInteger(*ie.Ul_PrioritizationThres_r16, sLUESelectedConfigR16UlPrioritizationThresR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-PrioritizationThres-r16: integer
	{
		if ie.Sl_PrioritizationThres_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PrioritizationThres_r16, sLUESelectedConfigR16SlPrioritizationThresR16Constraints); err != nil {
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
					{Name: "sl-CBR-CommonTxDedicatedSL-PRS-RP-List-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 != nil {
				if err := ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18.Encode(ex); err != nil {
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

func (ie *SL_UE_SelectedConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLUESelectedConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PSSCH-TxConfigList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PSSCH_TxConfigList_r16 = new(SL_PSSCH_TxConfigList_r16)
			if err := ie.Sl_PSSCH_TxConfigList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-ProbResourceKeep-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLUESelectedConfigR16SlProbResourceKeepR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ProbResourceKeep_r16 = &idx
		}
	}

	// 5. sl-ReselectAfter-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLUESelectedConfigR16SlReselectAfterR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ReselectAfter_r16 = &idx
		}
	}

	// 6. sl-CBR-CommonTxConfigList-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_CBR_CommonTxConfigList_r16 = new(SL_CBR_CommonTxConfigList_r16)
			if err := ie.Sl_CBR_CommonTxConfigList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. ul-PrioritizationThres-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLUESelectedConfigR16UlPrioritizationThresR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_PrioritizationThres_r16 = &v
		}
	}

	// 8. sl-PrioritizationThres-r16: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLUESelectedConfigR16SlPrioritizationThresR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PrioritizationThres_r16 = &v
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
				{Name: "sl-CBR-CommonTxDedicatedSL-PRS-RP-List-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 = new(SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18)
			if err := ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
