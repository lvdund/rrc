// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSet (line 19428).

var featureSetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra"},
		{Name: "nr"},
	},
}

const (
	FeatureSet_Eutra = 0
	FeatureSet_Nr    = 1
)

type FeatureSet struct {
	Choice int
	Eutra  *struct {
		DownlinkSetEUTRA FeatureSetEUTRA_DownlinkId
		UplinkSetEUTRA   FeatureSetEUTRA_UplinkId
	}
	Nr *struct {
		DownlinkSetNR FeatureSetDownlinkId
		UplinkSetNR   FeatureSetUplinkId
	}
}

func (ie *FeatureSet) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(featureSetConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case FeatureSet_Eutra:
		c := (*ie.Eutra)
		if err := c.DownlinkSetEUTRA.Encode(e); err != nil {
			return err
		}
		if err := c.UplinkSetEUTRA.Encode(e); err != nil {
			return err
		}
	case FeatureSet_Nr:
		c := (*ie.Nr)
		if err := c.DownlinkSetNR.Encode(e); err != nil {
			return err
		}
		if err := c.UplinkSetNR.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "FeatureSet",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *FeatureSet) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(featureSetConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "FeatureSet", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case FeatureSet_Eutra:
		ie.Eutra = &struct {
			DownlinkSetEUTRA FeatureSetEUTRA_DownlinkId
			UplinkSetEUTRA   FeatureSetEUTRA_UplinkId
		}{}
		c := (*ie.Eutra)
		{
			if err := c.DownlinkSetEUTRA.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.UplinkSetEUTRA.Decode(d); err != nil {
				return err
			}
		}
	case FeatureSet_Nr:
		ie.Nr = &struct {
			DownlinkSetNR FeatureSetDownlinkId
			UplinkSetNR   FeatureSetUplinkId
		}{}
		c := (*ie.Nr)
		{
			if err := c.DownlinkSetNR.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.UplinkSetNR.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "FeatureSet", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
