// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ScheduledConfig-r16 (line 28191).

var sLScheduledConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RNTI-r16"},
		{Name: "mac-MainConfigSL-r16", Optional: true},
		{Name: "sl-CS-RNTI-r16", Optional: true},
		{Name: "sl-PSFCH-ToPUCCH-r16", Optional: true},
		{Name: "sl-ConfiguredGrantConfigList-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sLScheduledConfigR16SlPSFCHToPUCCHR16Constraints = per.SizeRange(1, 8)

var sLScheduledConfigR16ExtSlDCIToSLTransR16Constraints = per.SizeRange(1, 8)

type SL_ScheduledConfig_r16 struct {
	Sl_RNTI_r16                                         RNTI_Value
	Mac_MainConfigSL_r16                                *MAC_MainConfigSL_r16
	Sl_CS_RNTI_r16                                      *RNTI_Value
	Sl_PSFCH_ToPUCCH_r16                                []int64
	Sl_ConfiguredGrantConfigList_r16                    *SL_ConfiguredGrantConfigList_r16
	Sl_DCI_ToSL_Trans_r16                               []int64
	Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 *SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18
	Sl_PRS_RNTI_r18                                     *RNTI_Value
	Sl_PRS_CS_RNTI_r18                                  *RNTI_Value
}

func (ie *SL_ScheduledConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLScheduledConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_DCI_ToSL_Trans_r16 != nil
	hasExtGroup1 := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 != nil || ie.Sl_PRS_RNTI_r18 != nil || ie.Sl_PRS_CS_RNTI_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_MainConfigSL_r16 != nil, ie.Sl_CS_RNTI_r16 != nil, ie.Sl_PSFCH_ToPUCCH_r16 != nil, ie.Sl_ConfiguredGrantConfigList_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-RNTI-r16: ref
	{
		if err := ie.Sl_RNTI_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. mac-MainConfigSL-r16: ref
	{
		if ie.Mac_MainConfigSL_r16 != nil {
			if err := ie.Mac_MainConfigSL_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-CS-RNTI-r16: ref
	{
		if ie.Sl_CS_RNTI_r16 != nil {
			if err := ie.Sl_CS_RNTI_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-PSFCH-ToPUCCH-r16: sequence-of
	{
		if ie.Sl_PSFCH_ToPUCCH_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLScheduledConfigR16SlPSFCHToPUCCHR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PSFCH_ToPUCCH_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PSFCH_ToPUCCH_r16 {
				if err := e.EncodeInteger(ie.Sl_PSFCH_ToPUCCH_r16[i], per.Constrained(0, 15)); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-ConfiguredGrantConfigList-r16: ref
	{
		if ie.Sl_ConfiguredGrantConfigList_r16 != nil {
			if err := ie.Sl_ConfiguredGrantConfigList_r16.Encode(e); err != nil {
				return err
			}
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
					{Name: "sl-DCI-ToSL-Trans-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DCI_ToSL_Trans_r16 != nil}); err != nil {
				return err
			}

			if ie.Sl_DCI_ToSL_Trans_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLScheduledConfigR16ExtSlDCIToSLTransR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_DCI_ToSL_Trans_r16))); err != nil {
					return err
				}
				for i := range ie.Sl_DCI_ToSL_Trans_r16 {
					if err := ex.EncodeInteger(ie.Sl_DCI_ToSL_Trans_r16[i], per.Constrained(1, 32)); err != nil {
						return err
					}
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
					{Name: "sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-List-r18", Optional: true},
					{Name: "sl-PRS-RNTI-r18", Optional: true},
					{Name: "sl-PRS-CS-RNTI-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 != nil, ie.Sl_PRS_RNTI_r18 != nil, ie.Sl_PRS_CS_RNTI_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 != nil {
				if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_RNTI_r18 != nil {
				if err := ie.Sl_PRS_RNTI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_CS_RNTI_r18 != nil {
				if err := ie.Sl_PRS_CS_RNTI_r18.Encode(ex); err != nil {
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

func (ie *SL_ScheduledConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLScheduledConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-RNTI-r16: ref
	{
		if err := ie.Sl_RNTI_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. mac-MainConfigSL-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_MainConfigSL_r16 = new(MAC_MainConfigSL_r16)
			if err := ie.Mac_MainConfigSL_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-CS-RNTI-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_CS_RNTI_r16 = new(RNTI_Value)
			if err := ie.Sl_CS_RNTI_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-PSFCH-ToPUCCH-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLScheduledConfigR16SlPSFCHToPUCCHR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_ToPUCCH_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				ie.Sl_PSFCH_ToPUCCH_r16[i] = v
			}
		}
	}

	// 7. sl-ConfiguredGrantConfigList-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_ConfiguredGrantConfigList_r16 = new(SL_ConfiguredGrantConfigList_r16)
			if err := ie.Sl_ConfiguredGrantConfigList_r16.Decode(d); err != nil {
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
				{Name: "sl-DCI-ToSL-Trans-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLScheduledConfigR16ExtSlDCIToSLTransR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DCI_ToSL_Trans_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				ie.Sl_DCI_ToSL_Trans_r16[i] = v
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
				{Name: "sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-List-r18", Optional: true},
				{Name: "sl-PRS-RNTI-r18", Optional: true},
				{Name: "sl-PRS-CS-RNTI-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 = new(SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18)
			if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_PRS_RNTI_r18 = new(RNTI_Value)
			if err := ie.Sl_PRS_RNTI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Sl_PRS_CS_RNTI_r18 = new(RNTI_Value)
			if err := ie.Sl_PRS_CS_RNTI_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
