// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParameters-v1610 (line 18516).

var codebookParametersV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedCSI-RS-ResourceListAlt-r16", Optional: true},
	},
}

var codebookParametersV1610SupportedCSIRSResourceListAltR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1-SinglePanel-r16", Optional: true},
		{Name: "type1-MultiPanel-r16", Optional: true},
		{Name: "type2-r16", Optional: true},
		{Name: "type2-PortSelection-r16", Optional: true},
	},
}

var codebookParametersV1610SupportedCSIRSResourceListAltR16Type1SinglePanelR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

var codebookParametersV1610SupportedCSIRSResourceListAltR16Type1MultiPanelR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

var codebookParametersV1610SupportedCSIRSResourceListAltR16Type2R16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

var codebookParametersV1610SupportedCSIRSResourceListAltR16Type2PortSelectionR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

type CodebookParameters_v1610 struct {
	SupportedCSI_RS_ResourceListAlt_r16 *struct {
		Type1_SinglePanel_r16   []int64
		Type1_MultiPanel_r16    []int64
		Type2_r16               []int64
		Type2_PortSelection_r16 []int64
	}
}

func (ie *CodebookParameters_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedCSI_RS_ResourceListAlt_r16 != nil}); err != nil {
		return err
	}

	// 2. supportedCSI-RS-ResourceListAlt-r16: sequence
	{
		if ie.SupportedCSI_RS_ResourceListAlt_r16 != nil {
			c := ie.SupportedCSI_RS_ResourceListAlt_r16
			codebookParametersV1610SupportedCSIRSResourceListAltR16Seq := e.NewSequenceEncoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Constraints)
			if err := codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.EncodePreamble([]bool{c.Type1_SinglePanel_r16 != nil, c.Type1_MultiPanel_r16 != nil, c.Type2_r16 != nil, c.Type2_PortSelection_r16 != nil}); err != nil {
				return err
			}
			if c.Type1_SinglePanel_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type1SinglePanelR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1_SinglePanel_r16))); err != nil {
					return err
				}
				for i := range c.Type1_SinglePanel_r16 {
					if err := e.EncodeInteger(c.Type1_SinglePanel_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type1_MultiPanel_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type1MultiPanelR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1_MultiPanel_r16))); err != nil {
					return err
				}
				for i := range c.Type1_MultiPanel_r16 {
					if err := e.EncodeInteger(c.Type1_MultiPanel_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type2_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type2R16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type2_r16))); err != nil {
					return err
				}
				for i := range c.Type2_r16 {
					if err := e.EncodeInteger(c.Type2_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type2_PortSelection_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type2PortSelectionR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type2_PortSelection_r16))); err != nil {
					return err
				}
				for i := range c.Type2_PortSelection_r16 {
					if err := e.EncodeInteger(c.Type2_PortSelection_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParameters_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedCSI-RS-ResourceListAlt-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedCSI_RS_ResourceListAlt_r16 = &struct {
				Type1_SinglePanel_r16   []int64
				Type1_MultiPanel_r16    []int64
				Type2_r16               []int64
				Type2_PortSelection_r16 []int64
			}{}
			c := ie.SupportedCSI_RS_ResourceListAlt_r16
			codebookParametersV1610SupportedCSIRSResourceListAltR16Seq := d.NewSequenceDecoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Constraints)
			if err := codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type1SinglePanelR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1_SinglePanel_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1_SinglePanel_r16[i] = v
				}
			}
			if codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type1MultiPanelR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1_MultiPanel_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1_MultiPanel_r16[i] = v
				}
			}
			if codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.IsComponentPresent(2) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type2R16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type2_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type2_r16[i] = v
				}
			}
			if codebookParametersV1610SupportedCSIRSResourceListAltR16Seq.IsComponentPresent(3) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersV1610SupportedCSIRSResourceListAltR16Type2PortSelectionR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type2_PortSelection_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type2_PortSelection_r16[i] = v
				}
			}
		}
	}

	return nil
}
