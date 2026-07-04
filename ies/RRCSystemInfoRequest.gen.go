// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSystemInfoRequest (line 1808).

var rRCSystemInfoRequestConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var rRCSystemInfoRequestCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcSystemInfoRequest"},
		{Name: "criticalExtensionsFuture-r16"},
	},
}

const (
	RRCSystemInfoRequest_CriticalExtensions_RrcSystemInfoRequest         = 0
	RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16 = 1
)

var rRCSystemInfoRequestCriticalExtensionsCriticalExtensionsFutureR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcPosSystemInfoRequest-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_RrcPosSystemInfoRequest_r16 = 0
	RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture    = 1
)

type RRCSystemInfoRequest struct {
	CriticalExtensions struct {
		Choice                       int
		RrcSystemInfoRequest         *RRCSystemInfoRequest_IEs
		CriticalExtensionsFuture_r16 *struct {
			Choice                      int
			RrcPosSystemInfoRequest_r16 *RRC_PosSystemInfoRequest_r16_IEs
			CriticalExtensionsFuture    *struct{}
		}
	}
}

func (ie *RRCSystemInfoRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSystemInfoRequestConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCSystemInfoRequestCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCSystemInfoRequest_CriticalExtensions_RrcSystemInfoRequest:
			if err := ie.CriticalExtensions.RrcSystemInfoRequest.Encode(e); err != nil {
				return err
			}
		case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16:
			choiceEnc := e.NewChoiceEncoder(rRCSystemInfoRequestCriticalExtensionsCriticalExtensionsFutureR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice {
			case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_RrcPosSystemInfoRequest_r16:
				if err := (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).RrcPosSystemInfoRequest_r16.Encode(e); err != nil {
					return err
				}
			case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture:
				// empty SEQUENCE alternative
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCSystemInfoRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSystemInfoRequestConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCSystemInfoRequestCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCSystemInfoRequest_CriticalExtensions_RrcSystemInfoRequest:
			ie.CriticalExtensions.RrcSystemInfoRequest = new(RRCSystemInfoRequest_IEs)
			if err := ie.CriticalExtensions.RrcSystemInfoRequest.Decode(d); err != nil {
				return err
			}
		case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16:
			ie.CriticalExtensions.CriticalExtensionsFuture_r16 = &struct {
				Choice                      int
				RrcPosSystemInfoRequest_r16 *RRC_PosSystemInfoRequest_r16_IEs
				CriticalExtensionsFuture    *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(rRCSystemInfoRequestCriticalExtensionsCriticalExtensionsFutureR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice = int(index)
			switch index {
			case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_RrcPosSystemInfoRequest_r16:
				(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).RrcPosSystemInfoRequest_r16 = new(RRC_PosSystemInfoRequest_r16_IEs)
				if err := (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).RrcPosSystemInfoRequest_r16.Decode(d); err != nil {
					return err
				}
			case RRCSystemInfoRequest_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture:
				(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).CriticalExtensionsFuture = &struct{}{}
				// empty SEQUENCE alternative
			}
		}
	}

	return nil
}
