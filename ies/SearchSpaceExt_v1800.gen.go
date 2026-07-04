// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpaceExt-v1800 (line 14522).

var searchSpaceExtV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "searchSpaceType-r18", Optional: true},
	},
}

var searchSpaceExt_v1800SearchSpaceTypeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "common-r18"},
		{Name: "ue-Specific-r18"},
	},
}

const (
	SearchSpaceExt_v1800_SearchSpaceType_r18_Common_r18      = 0
	SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18 = 1
)

var searchSpaceExtV1800SearchSpaceTypeR18CommonR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-Format2-9-r18", Optional: true},
	},
}

var searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-FormatsMC-r18", Optional: true},
	},
}

const (
	SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats0_3         = 0
	SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats1_3         = 1
	SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats0_3_And_1_3 = 2
)

var searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18DciFormatsMCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats0_3, SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats1_3, SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18_Dci_FormatsMC_r18_Formats0_3_And_1_3},
}

type SearchSpaceExt_v1800 struct {
	SearchSpaceType_r18 *struct {
		Choice          int
		Common_r18      *struct{ Dci_Format2_9_r18 *struct{} }
		Ue_Specific_r18 *struct{ Dci_FormatsMC_r18 *int64 }
	}
}

func (ie *SearchSpaceExt_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceExtV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SearchSpaceType_r18 != nil}); err != nil {
		return err
	}

	// 2. searchSpaceType-r18: choice
	{
		if ie.SearchSpaceType_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(searchSpaceExt_v1800SearchSpaceTypeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SearchSpaceType_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SearchSpaceType_r18).Choice {
			case SearchSpaceExt_v1800_SearchSpaceType_r18_Common_r18:
				c := (*(*ie.SearchSpaceType_r18).Common_r18)
				searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq := e.NewSequenceEncoder(searchSpaceExtV1800SearchSpaceTypeR18CommonR18Constraints)
				if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq.EncodeExtensionBit(false); err != nil {
					return err
				}
				if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq.EncodePreamble([]bool{c.Dci_Format2_9_r18 != nil}); err != nil {
					return err
				}
				if c.Dci_Format2_9_r18 != nil {
					searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Seq := e.NewSequenceEncoder(searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Constraints)
					if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
			case SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18:
				c := (*(*ie.SearchSpaceType_r18).Ue_Specific_r18)
				searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq := e.NewSequenceEncoder(searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Constraints)
				if err := searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq.EncodeExtensionBit(false); err != nil {
					return err
				}
				if err := searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq.EncodePreamble([]bool{c.Dci_FormatsMC_r18 != nil}); err != nil {
					return err
				}
				if c.Dci_FormatsMC_r18 != nil {
					if err := e.EncodeEnumerated((*c.Dci_FormatsMC_r18), searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18DciFormatsMCR18Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SearchSpaceType_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SearchSpaceExt_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceExtV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. searchSpaceType-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.SearchSpaceType_r18 = &struct {
				Choice          int
				Common_r18      *struct{ Dci_Format2_9_r18 *struct{} }
				Ue_Specific_r18 *struct{ Dci_FormatsMC_r18 *int64 }
			}{}
			choiceDec := d.NewChoiceDecoder(searchSpaceExt_v1800SearchSpaceTypeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SearchSpaceType_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SearchSpaceExt_v1800_SearchSpaceType_r18_Common_r18:
				(*ie.SearchSpaceType_r18).Common_r18 = &struct{ Dci_Format2_9_r18 *struct{} }{}
				c := (*(*ie.SearchSpaceType_r18).Common_r18)
				searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq := d.NewSequenceDecoder(searchSpaceExtV1800SearchSpaceTypeR18CommonR18Constraints)
				if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if searchSpaceExtV1800SearchSpaceTypeR18CommonR18Seq.IsComponentPresent(0) {
					c.Dci_Format2_9_r18 = &struct{}{}
					searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Seq := d.NewSequenceDecoder(searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Constraints)
					if err := searchSpaceExtV1800SearchSpaceTypeR18CommonR18DciFormat29R18Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
			case SearchSpaceExt_v1800_SearchSpaceType_r18_Ue_Specific_r18:
				(*ie.SearchSpaceType_r18).Ue_Specific_r18 = &struct{ Dci_FormatsMC_r18 *int64 }{}
				c := (*(*ie.SearchSpaceType_r18).Ue_Specific_r18)
				searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq := d.NewSequenceDecoder(searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Constraints)
				if err := searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18Seq.IsComponentPresent(0) {
					c.Dci_FormatsMC_r18 = new(int64)
					v, err := d.DecodeEnumerated(searchSpaceExtV1800SearchSpaceTypeR18UeSpecificR18DciFormatsMCR18Constraints)
					if err != nil {
						return err
					}
					(*c.Dci_FormatsMC_r18) = v
				}
			}
		}
	}

	return nil
}
