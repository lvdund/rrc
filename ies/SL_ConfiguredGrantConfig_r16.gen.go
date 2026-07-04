// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ConfiguredGrantConfig-r16 (line 27052).

var sLConfiguredGrantConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigIndexCG-r16"},
		{Name: "sl-PeriodCG-r16", Optional: true},
		{Name: "sl-NrOfHARQ-Processes-r16", Optional: true},
		{Name: "sl-HARQ-ProcID-offset-r16", Optional: true},
		{Name: "sl-CG-MaxTransNumList-r16", Optional: true},
		{Name: "rrc-ConfiguredSidelinkGrant-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sLConfiguredGrantConfigR16SlNrOfHARQProcessesR16Constraints = per.Constrained(1, 16)

var sLConfiguredGrantConfigR16SlHARQProcIDOffsetR16Constraints = per.Constrained(0, 15)

var sLConfiguredGrantConfigR16SlStartRBsetCGType1R18Constraints = per.Constrained(0, 4)

var sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TimeResourceCG-Type1-r16", Optional: true},
		{Name: "sl-StartSubchannelCG-Type1-r16", Optional: true},
		{Name: "sl-FreqResourceCG-Type1-r16", Optional: true},
		{Name: "sl-TimeOffsetCG-Type1-r16", Optional: true},
		{Name: "sl-N1PUCCH-AN-r16", Optional: true},
		{Name: "sl-PSFCH-ToPUCCH-CG-Type1-r16", Optional: true},
		{Name: "sl-ResourcePoolID-r16", Optional: true},
		{Name: "sl-TimeReferenceSFN-Type1-r16", Optional: true},
	},
}

const (
	SL_ConfiguredGrantConfig_r16_Rrc_ConfiguredSidelinkGrant_r16_Sl_TimeReferenceSFN_Type1_r16_Sfn512 = 0
)

var sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16SlTimeReferenceSFNType1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfiguredGrantConfig_r16_Rrc_ConfiguredSidelinkGrant_r16_Sl_TimeReferenceSFN_Type1_r16_Sfn512},
}

type SL_ConfiguredGrantConfig_r16 struct {
	Sl_ConfigIndexCG_r16            SL_ConfigIndexCG_r16
	Sl_PeriodCG_r16                 *SL_PeriodCG_r16
	Sl_NrOfHARQ_Processes_r16       *int64
	Sl_HARQ_ProcID_Offset_r16       *int64
	Sl_CG_MaxTransNumList_r16       *SL_CG_MaxTransNumList_r16
	Rrc_ConfiguredSidelinkGrant_r16 *struct {
		Sl_TimeResourceCG_Type1_r16    *int64
		Sl_StartSubchannelCG_Type1_r16 *int64
		Sl_FreqResourceCG_Type1_r16    *int64
		Sl_TimeOffsetCG_Type1_r16      *int64
		Sl_N1PUCCH_AN_r16              *PUCCH_ResourceId
		Sl_PSFCH_ToPUCCH_CG_Type1_r16  *int64
		Sl_ResourcePoolID_r16          *SL_ResourcePoolID_r16
		Sl_TimeReferenceSFN_Type1_r16  *int64
	}
	Sl_N1PUCCH_AN_Type2_r16   *PUCCH_ResourceId
	Sl_StartRBsetCG_Type1_r18 *int64
}

func (ie *SL_ConfiguredGrantConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfiguredGrantConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_N1PUCCH_AN_Type2_r16 != nil
	hasExtGroup1 := ie.Sl_StartRBsetCG_Type1_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PeriodCG_r16 != nil, ie.Sl_NrOfHARQ_Processes_r16 != nil, ie.Sl_HARQ_ProcID_Offset_r16 != nil, ie.Sl_CG_MaxTransNumList_r16 != nil, ie.Rrc_ConfiguredSidelinkGrant_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-ConfigIndexCG-r16: ref
	{
		if err := ie.Sl_ConfigIndexCG_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-PeriodCG-r16: ref
	{
		if ie.Sl_PeriodCG_r16 != nil {
			if err := ie.Sl_PeriodCG_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-NrOfHARQ-Processes-r16: integer
	{
		if ie.Sl_NrOfHARQ_Processes_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_NrOfHARQ_Processes_r16, sLConfiguredGrantConfigR16SlNrOfHARQProcessesR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-HARQ-ProcID-offset-r16: integer
	{
		if ie.Sl_HARQ_ProcID_Offset_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_HARQ_ProcID_Offset_r16, sLConfiguredGrantConfigR16SlHARQProcIDOffsetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-CG-MaxTransNumList-r16: ref
	{
		if ie.Sl_CG_MaxTransNumList_r16 != nil {
			if err := ie.Sl_CG_MaxTransNumList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. rrc-ConfiguredSidelinkGrant-r16: sequence
	{
		if ie.Rrc_ConfiguredSidelinkGrant_r16 != nil {
			c := ie.Rrc_ConfiguredSidelinkGrant_r16
			sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq := e.NewSequenceEncoder(sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Constraints)
			if err := sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.EncodePreamble([]bool{c.Sl_TimeResourceCG_Type1_r16 != nil, c.Sl_StartSubchannelCG_Type1_r16 != nil, c.Sl_FreqResourceCG_Type1_r16 != nil, c.Sl_TimeOffsetCG_Type1_r16 != nil, c.Sl_N1PUCCH_AN_r16 != nil, c.Sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil, c.Sl_ResourcePoolID_r16 != nil, c.Sl_TimeReferenceSFN_Type1_r16 != nil}); err != nil {
				return err
			}
			if c.Sl_TimeResourceCG_Type1_r16 != nil {
				if err := e.EncodeInteger((*c.Sl_TimeResourceCG_Type1_r16), per.Constrained(0, 496)); err != nil {
					return err
				}
			}
			if c.Sl_StartSubchannelCG_Type1_r16 != nil {
				if err := e.EncodeInteger((*c.Sl_StartSubchannelCG_Type1_r16), per.Constrained(0, 26)); err != nil {
					return err
				}
			}
			if c.Sl_FreqResourceCG_Type1_r16 != nil {
				if err := e.EncodeInteger((*c.Sl_FreqResourceCG_Type1_r16), per.Constrained(0, 6929)); err != nil {
					return err
				}
			}
			if c.Sl_TimeOffsetCG_Type1_r16 != nil {
				if err := e.EncodeInteger((*c.Sl_TimeOffsetCG_Type1_r16), per.Constrained(0, 7999)); err != nil {
					return err
				}
			}
			if c.Sl_N1PUCCH_AN_r16 != nil {
				if err := c.Sl_N1PUCCH_AN_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil {
				if err := e.EncodeInteger((*c.Sl_PSFCH_ToPUCCH_CG_Type1_r16), per.Constrained(0, 15)); err != nil {
					return err
				}
			}
			if c.Sl_ResourcePoolID_r16 != nil {
				if err := c.Sl_ResourcePoolID_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Sl_TimeReferenceSFN_Type1_r16 != nil {
				if err := e.EncodeEnumerated((*c.Sl_TimeReferenceSFN_Type1_r16), sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16SlTimeReferenceSFNType1R16Constraints); err != nil {
					return err
				}
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
					{Name: "sl-N1PUCCH-AN-Type2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_N1PUCCH_AN_Type2_r16 != nil}); err != nil {
				return err
			}

			if ie.Sl_N1PUCCH_AN_Type2_r16 != nil {
				if err := ie.Sl_N1PUCCH_AN_Type2_r16.Encode(ex); err != nil {
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
					{Name: "sl-StartRBsetCG-Type1-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_StartRBsetCG_Type1_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_StartRBsetCG_Type1_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_StartRBsetCG_Type1_r18, sLConfiguredGrantConfigR16SlStartRBsetCGType1R18Constraints); err != nil {
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

func (ie *SL_ConfiguredGrantConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfiguredGrantConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-ConfigIndexCG-r16: ref
	{
		if err := ie.Sl_ConfigIndexCG_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-PeriodCG-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PeriodCG_r16 = new(SL_PeriodCG_r16)
			if err := ie.Sl_PeriodCG_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-NrOfHARQ-Processes-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLConfiguredGrantConfigR16SlNrOfHARQProcessesR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NrOfHARQ_Processes_r16 = &v
		}
	}

	// 6. sl-HARQ-ProcID-offset-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLConfiguredGrantConfigR16SlHARQProcIDOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_HARQ_ProcID_Offset_r16 = &v
		}
	}

	// 7. sl-CG-MaxTransNumList-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_CG_MaxTransNumList_r16 = new(SL_CG_MaxTransNumList_r16)
			if err := ie.Sl_CG_MaxTransNumList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. rrc-ConfiguredSidelinkGrant-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Rrc_ConfiguredSidelinkGrant_r16 = &struct {
				Sl_TimeResourceCG_Type1_r16    *int64
				Sl_StartSubchannelCG_Type1_r16 *int64
				Sl_FreqResourceCG_Type1_r16    *int64
				Sl_TimeOffsetCG_Type1_r16      *int64
				Sl_N1PUCCH_AN_r16              *PUCCH_ResourceId
				Sl_PSFCH_ToPUCCH_CG_Type1_r16  *int64
				Sl_ResourcePoolID_r16          *SL_ResourcePoolID_r16
				Sl_TimeReferenceSFN_Type1_r16  *int64
			}{}
			c := ie.Rrc_ConfiguredSidelinkGrant_r16
			sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq := d.NewSequenceDecoder(sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Constraints)
			if err := sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(0) {
				c.Sl_TimeResourceCG_Type1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 496))
				if err != nil {
					return err
				}
				(*c.Sl_TimeResourceCG_Type1_r16) = v
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(1) {
				c.Sl_StartSubchannelCG_Type1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 26))
				if err != nil {
					return err
				}
				(*c.Sl_StartSubchannelCG_Type1_r16) = v
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(2) {
				c.Sl_FreqResourceCG_Type1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 6929))
				if err != nil {
					return err
				}
				(*c.Sl_FreqResourceCG_Type1_r16) = v
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(3) {
				c.Sl_TimeOffsetCG_Type1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7999))
				if err != nil {
					return err
				}
				(*c.Sl_TimeOffsetCG_Type1_r16) = v
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(4) {
				c.Sl_N1PUCCH_AN_r16 = new(PUCCH_ResourceId)
				if err := (*c.Sl_N1PUCCH_AN_r16).Decode(d); err != nil {
					return err
				}
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(5) {
				c.Sl_PSFCH_ToPUCCH_CG_Type1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*c.Sl_PSFCH_ToPUCCH_CG_Type1_r16) = v
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(6) {
				c.Sl_ResourcePoolID_r16 = new(SL_ResourcePoolID_r16)
				if err := (*c.Sl_ResourcePoolID_r16).Decode(d); err != nil {
					return err
				}
			}
			if sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16Seq.IsComponentPresent(7) {
				c.Sl_TimeReferenceSFN_Type1_r16 = new(int64)
				v, err := d.DecodeEnumerated(sLConfiguredGrantConfigR16RrcConfiguredSidelinkGrantR16SlTimeReferenceSFNType1R16Constraints)
				if err != nil {
					return err
				}
				(*c.Sl_TimeReferenceSFN_Type1_r16) = v
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
				{Name: "sl-N1PUCCH-AN-Type2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_N1PUCCH_AN_Type2_r16 = new(PUCCH_ResourceId)
			if err := ie.Sl_N1PUCCH_AN_Type2_r16.Decode(dx); err != nil {
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
				{Name: "sl-StartRBsetCG-Type1-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLConfiguredGrantConfigR16SlStartRBsetCGType1R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartRBsetCG_Type1_r18 = &v
		}
	}

	return nil
}
