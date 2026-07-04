// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-HARQ-ACK-EnhType3-r17 (line 11710).

var pDSCHHARQACKEnhType3R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-HARQ-ACK-EnhType3Index-r17"},
		{Name: "applicable-r17"},
		{Name: "pdsch-HARQ-ACK-EnhType3NDI-r17", Optional: true},
		{Name: "pdsch-HARQ-ACK-EnhType3CBG-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var pDSCH_HARQ_ACK_EnhType3_r17ApplicableR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perCC"},
		{Name: "perHARQ"},
	},
}

const (
	PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerCC   = 0
	PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerHARQ = 1
)

const (
	PDSCH_HARQ_ACK_EnhType3_r17_Pdsch_HARQ_ACK_EnhType3NDI_r17_True = 0
)

var pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3NDIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_HARQ_ACK_EnhType3_r17_Pdsch_HARQ_ACK_EnhType3NDI_r17_True},
}

const (
	PDSCH_HARQ_ACK_EnhType3_r17_Pdsch_HARQ_ACK_EnhType3CBG_r17_True = 0
)

var pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3CBGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_HARQ_ACK_EnhType3_r17_Pdsch_HARQ_ACK_EnhType3CBG_r17_True},
}

var pDSCHHARQACKEnhType3R17ExtPerHARQExtR17Constraints = per.SizeRange(1, common.MaxNrofServingCells)

var pDSCHHARQACKEnhType3R17ApplicableR17PerCCConstraints = per.SizeRange(1, common.MaxNrofServingCells)

var pDSCHHARQACKEnhType3R17ApplicableR17PerHARQConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type PDSCH_HARQ_ACK_EnhType3_r17 struct {
	Pdsch_HARQ_ACK_EnhType3Index_r17 PDSCH_HARQ_ACK_EnhType3Index_r17
	Applicable_r17                   struct {
		Choice  int
		PerCC   *[]int64
		PerHARQ *[]per.BitString
	}
	Pdsch_HARQ_ACK_EnhType3NDI_r17 *int64
	Pdsch_HARQ_ACK_EnhType3CBG_r17 *int64
	PerHARQ_Ext_r17                []per.BitString
}

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHHARQACKEnhType3R17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PerHARQ_Ext_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 != nil}); err != nil {
		return err
	}

	// 3. pdsch-HARQ-ACK-EnhType3Index-r17: ref
	{
		if err := ie.Pdsch_HARQ_ACK_EnhType3Index_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. applicable-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(pDSCH_HARQ_ACK_EnhType3_r17ApplicableR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Applicable_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Applicable_r17.Choice {
		case PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerCC:
			seqOf := e.NewSequenceOfEncoder(pDSCHHARQACKEnhType3R17ApplicableR17PerCCConstraints)
			if err := seqOf.EncodeLength(int64(len((*ie.Applicable_r17.PerCC)))); err != nil {
				return err
			}
			for i := range *ie.Applicable_r17.PerCC {
				if err := e.EncodeInteger((*ie.Applicable_r17.PerCC)[i], per.Constrained(0, 1)); err != nil {
					return err
				}
			}
		case PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerHARQ:
			seqOf := e.NewSequenceOfEncoder(pDSCHHARQACKEnhType3R17ApplicableR17PerHARQConstraints)
			if err := seqOf.EncodeLength(int64(len((*ie.Applicable_r17.PerHARQ)))); err != nil {
				return err
			}
			for i := range *ie.Applicable_r17.PerHARQ {
				if err := e.EncodeBitString((*ie.Applicable_r17.PerHARQ)[i], per.FixedSize(16)); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Applicable_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 5. pdsch-HARQ-ACK-EnhType3NDI-r17: enumerated
	{
		if ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3NDI_r17, pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3NDIR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pdsch-HARQ-ACK-EnhType3CBG-r17: enumerated
	{
		if ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3CBG_r17, pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3CBGR17Constraints); err != nil {
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
					{Name: "perHARQ-Ext-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PerHARQ_Ext_r17 != nil}); err != nil {
				return err
			}

			if ie.PerHARQ_Ext_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDSCHHARQACKEnhType3R17ExtPerHARQExtR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.PerHARQ_Ext_r17))); err != nil {
					return err
				}
				for i := range ie.PerHARQ_Ext_r17 {
					if err := ex.EncodeBitString(ie.PerHARQ_Ext_r17[i], per.FixedSize(32)); err != nil {
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

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHHARQACKEnhType3R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdsch-HARQ-ACK-EnhType3Index-r17: ref
	{
		if err := ie.Pdsch_HARQ_ACK_EnhType3Index_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. applicable-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(pDSCH_HARQ_ACK_EnhType3_r17ApplicableR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Applicable_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerCC:
			ie.Applicable_r17.PerCC = new([]int64)
			seqOf := d.NewSequenceOfDecoder(pDSCHHARQACKEnhType3R17ApplicableR17PerCCConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.Applicable_r17.PerCC) = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*ie.Applicable_r17.PerCC)[i] = v
			}
		case PDSCH_HARQ_ACK_EnhType3_r17_Applicable_r17_PerHARQ:
			ie.Applicable_r17.PerHARQ = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(pDSCHHARQACKEnhType3R17ApplicableR17PerHARQConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.Applicable_r17.PerHARQ) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(16))
				if err != nil {
					return err
				}
				(*ie.Applicable_r17.PerHARQ)[i] = v
			}
		}
	}

	// 5. pdsch-HARQ-ACK-EnhType3NDI-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3NDIR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 = &idx
		}
	}

	// 6. pdsch-HARQ-ACK-EnhType3CBG-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pDSCHHARQACKEnhType3R17PdschHARQACKEnhType3CBGR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 = &idx
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
				{Name: "perHARQ-Ext-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pDSCHHARQACKEnhType3R17ExtPerHARQExtR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PerHARQ_Ext_r17 = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeBitString(per.FixedSize(32))
				if err != nil {
					return err
				}
				ie.PerHARQ_Ext_r17[i] = v
			}
		}
	}

	return nil
}
