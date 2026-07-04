// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersAdditionPerBC-r16 (line 18718).

var codebookParametersAdditionPerBCR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "etype2R1-r16", Optional: true},
		{Name: "etype2R2-r16", Optional: true},
		{Name: "etype2R1-PortSelection-r16", Optional: true},
		{Name: "etype2R2-PortSelection-r16", Optional: true},
	},
}

var codebookParametersAdditionPerBCR16Etype2R1R16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersAdditionPerBCR16Etype2R2R16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersAdditionPerBCR16Etype2R1PortSelectionR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersAdditionPerBCR16Etype2R2PortSelectionR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParametersAdditionPerBC_r16 struct {
	Etype2R1_r16               []int64
	Etype2R2_r16               []int64
	Etype2R1_PortSelection_r16 []int64
	Etype2R2_PortSelection_r16 []int64
}

func (ie *CodebookParametersAdditionPerBC_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersAdditionPerBCR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Etype2R1_r16 != nil, ie.Etype2R2_r16 != nil, ie.Etype2R1_PortSelection_r16 != nil, ie.Etype2R2_PortSelection_r16 != nil}); err != nil {
		return err
	}

	// 2. etype2R1-r16: sequence-of
	{
		if ie.Etype2R1_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionPerBCR16Etype2R1R16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Etype2R1_r16))); err != nil {
				return err
			}
			for i := range ie.Etype2R1_r16 {
				if err := e.EncodeInteger(ie.Etype2R1_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 3. etype2R2-r16: sequence-of
	{
		if ie.Etype2R2_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionPerBCR16Etype2R2R16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Etype2R2_r16))); err != nil {
				return err
			}
			for i := range ie.Etype2R2_r16 {
				if err := e.EncodeInteger(ie.Etype2R2_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. etype2R1-PortSelection-r16: sequence-of
	{
		if ie.Etype2R1_PortSelection_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionPerBCR16Etype2R1PortSelectionR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Etype2R1_PortSelection_r16))); err != nil {
				return err
			}
			for i := range ie.Etype2R1_PortSelection_r16 {
				if err := e.EncodeInteger(ie.Etype2R1_PortSelection_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. etype2R2-PortSelection-r16: sequence-of
	{
		if ie.Etype2R2_PortSelection_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionPerBCR16Etype2R2PortSelectionR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Etype2R2_PortSelection_r16))); err != nil {
				return err
			}
			for i := range ie.Etype2R2_PortSelection_r16 {
				if err := e.EncodeInteger(ie.Etype2R2_PortSelection_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParametersAdditionPerBC_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersAdditionPerBCR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. etype2R1-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionPerBCR16Etype2R1R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Etype2R1_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Etype2R1_r16[i] = v
			}
		}
	}

	// 3. etype2R2-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionPerBCR16Etype2R2R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Etype2R2_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Etype2R2_r16[i] = v
			}
		}
	}

	// 4. etype2R1-PortSelection-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionPerBCR16Etype2R1PortSelectionR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Etype2R1_PortSelection_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Etype2R1_PortSelection_r16[i] = v
			}
		}
	}

	// 5. etype2R2-PortSelection-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionPerBCR16Etype2R2PortSelectionR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Etype2R2_PortSelection_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Etype2R2_PortSelection_r16[i] = v
			}
		}
	}

	return nil
}
