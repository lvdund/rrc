// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicableDisasterInfo-r17 (line 4388).

var applicableDisasterInfoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "noDisasterRoaming-r17"},
		{Name: "disasterRelatedIndication-r17"},
		{Name: "commonPLMNs-r17"},
		{Name: "dedicatedPLMNs-r17"},
	},
}

const (
	ApplicableDisasterInfo_r17_NoDisasterRoaming_r17         = 0
	ApplicableDisasterInfo_r17_DisasterRelatedIndication_r17 = 1
	ApplicableDisasterInfo_r17_CommonPLMNs_r17               = 2
	ApplicableDisasterInfo_r17_DedicatedPLMNs_r17            = 3
)

var applicableDisasterInfoR17DedicatedPLMNsR17Constraints = per.SizeRange(1, common.MaxPLMN)

type ApplicableDisasterInfo_r17 struct {
	Choice                        int
	NoDisasterRoaming_r17         *struct{}
	DisasterRelatedIndication_r17 *struct{}
	CommonPLMNs_r17               *struct{}
	DedicatedPLMNs_r17            *[]PLMN_Identity
}

func (ie *ApplicableDisasterInfo_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(applicableDisasterInfoR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case ApplicableDisasterInfo_r17_NoDisasterRoaming_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_DisasterRelatedIndication_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_CommonPLMNs_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_DedicatedPLMNs_r17:
		seqOf := e.NewSequenceOfEncoder(applicableDisasterInfoR17DedicatedPLMNsR17Constraints)
		if err := seqOf.EncodeLength(int64(len((*ie.DedicatedPLMNs_r17)))); err != nil {
			return err
		}
		for i := range *ie.DedicatedPLMNs_r17 {
			if err := (*ie.DedicatedPLMNs_r17)[i].Encode(e); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "ApplicableDisasterInfo-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *ApplicableDisasterInfo_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(applicableDisasterInfoR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "ApplicableDisasterInfo-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case ApplicableDisasterInfo_r17_NoDisasterRoaming_r17:
		ie.NoDisasterRoaming_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_DisasterRelatedIndication_r17:
		ie.DisasterRelatedIndication_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_CommonPLMNs_r17:
		ie.CommonPLMNs_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case ApplicableDisasterInfo_r17_DedicatedPLMNs_r17:
		ie.DedicatedPLMNs_r17 = new([]PLMN_Identity)
		seqOf := d.NewSequenceOfDecoder(applicableDisasterInfoR17DedicatedPLMNsR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		(*ie.DedicatedPLMNs_r17) = make([]PLMN_Identity, n)
		for i := int64(0); i < n; i++ {
			if err := (*ie.DedicatedPLMNs_r17)[i].Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "ApplicableDisasterInfo-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
