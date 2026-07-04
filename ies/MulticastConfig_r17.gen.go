// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MulticastConfig-r17 (line 11751).

var multicastConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-HARQ-ACK-CodebookListMulticast-r17", Optional: true},
		{Name: "type1CodebookGenerationMode-r17", Optional: true},
	},
}

var multicastConfig_r17PdschHARQACKCodebookListMulticastR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Release = 0
	MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Setup   = 1
)

const (
	MulticastConfig_r17_Type1CodebookGenerationMode_r17_Mode1 = 0
	MulticastConfig_r17_Type1CodebookGenerationMode_r17_Mode2 = 1
)

var multicastConfigR17Type1CodebookGenerationModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MulticastConfig_r17_Type1CodebookGenerationMode_r17_Mode1, MulticastConfig_r17_Type1CodebookGenerationMode_r17_Mode2},
}

type MulticastConfig_r17 struct {
	Pdsch_HARQ_ACK_CodebookListMulticast_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_HARQ_ACK_CodebookList_r16
	}
	Type1CodebookGenerationMode_r17 *int64
}

func (ie *MulticastConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multicastConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil, ie.Type1CodebookGenerationMode_r17 != nil}); err != nil {
		return err
	}

	// 2. pdsch-HARQ-ACK-CodebookListMulticast-r17: choice
	{
		if ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(multicastConfig_r17PdschHARQACKCodebookListMulticastR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Choice {
			case MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Setup:
				if err := (*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. type1CodebookGenerationMode-r17: enumerated
	{
		if ie.Type1CodebookGenerationMode_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Type1CodebookGenerationMode_r17, multicastConfigR17Type1CodebookGenerationModeR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MulticastConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multicastConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdsch-HARQ-ACK-CodebookListMulticast-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_HARQ_ACK_CodebookList_r16
			}{}
			choiceDec := d.NewChoiceDecoder(multicastConfig_r17PdschHARQACKCodebookListMulticastR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Release:
				(*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MulticastConfig_r17_Pdsch_HARQ_ACK_CodebookListMulticast_r17_Setup:
				(*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Setup = new(PDSCH_HARQ_ACK_CodebookList_r16)
				if err := (*ie.Pdsch_HARQ_ACK_CodebookListMulticast_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. type1CodebookGenerationMode-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(multicastConfigR17Type1CodebookGenerationModeR17Constraints)
			if err != nil {
				return err
			}
			ie.Type1CodebookGenerationMode_r17 = &idx
		}
	}

	return nil
}
