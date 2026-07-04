// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosUplinkTransmissionWindowConfig-r18 (line 15763).

var sRSPosUplinkTransmissionWindowConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy"},
		{Name: "windowPeriodicityAndOffset-r18"},
		{Name: "duration-r18"},
	},
}

var sRSPosUplinkTransmissionWindowConfigR18DummyConstraints = per.Constrained(0, 1023)

var sRS_PosUplinkTransmissionWindowConfig_r18WindowPeriodicityAndOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodicityAndOffset-r18"},
		{Name: "periodicityAndOffsetExt-r18"},
	},
}

const (
	SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffset_r18    = 0
	SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffsetExt_r18 = 1
)

const (
	SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl1 = 0
	SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl2 = 1
	SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl4 = 2
	SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl6 = 3
)

var sRSPosUplinkTransmissionWindowConfigR18DurationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl1, SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl2, SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl4, SRS_PosUplinkTransmissionWindowConfig_r18_Duration_r18_Sl6},
}

type SRS_PosUplinkTransmissionWindowConfig_r18 struct {
	Dummy                          int64
	WindowPeriodicityAndOffset_r18 struct {
		Choice                      int
		PeriodicityAndOffset_r18    *SRS_PeriodicityAndOffset_r16
		PeriodicityAndOffsetExt_r18 *SRS_PeriodicityAndOffsetExt_r16
	}
	Duration_r18 int64
}

func (ie *SRS_PosUplinkTransmissionWindowConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosUplinkTransmissionWindowConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. dummy: integer
	{
		if err := e.EncodeInteger(ie.Dummy, sRSPosUplinkTransmissionWindowConfigR18DummyConstraints); err != nil {
			return err
		}
	}

	// 3. windowPeriodicityAndOffset-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_PosUplinkTransmissionWindowConfig_r18WindowPeriodicityAndOffsetR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.WindowPeriodicityAndOffset_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.WindowPeriodicityAndOffset_r18.Choice {
		case SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffset_r18:
			if err := ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffset_r18.Encode(e); err != nil {
				return err
			}
		case SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffsetExt_r18:
			if err := ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffsetExt_r18.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.WindowPeriodicityAndOffset_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 4. duration-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Duration_r18, sRSPosUplinkTransmissionWindowConfigR18DurationR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_PosUplinkTransmissionWindowConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosUplinkTransmissionWindowConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. dummy: integer
	{
		v0, err := d.DecodeInteger(sRSPosUplinkTransmissionWindowConfigR18DummyConstraints)
		if err != nil {
			return err
		}
		ie.Dummy = v0
	}

	// 3. windowPeriodicityAndOffset-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_PosUplinkTransmissionWindowConfig_r18WindowPeriodicityAndOffsetR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.WindowPeriodicityAndOffset_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffset_r18:
			ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffset_r18 = new(SRS_PeriodicityAndOffset_r16)
			if err := ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffset_r18.Decode(d); err != nil {
				return err
			}
		case SRS_PosUplinkTransmissionWindowConfig_r18_WindowPeriodicityAndOffset_r18_PeriodicityAndOffsetExt_r18:
			ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffsetExt_r18 = new(SRS_PeriodicityAndOffsetExt_r16)
			if err := ie.WindowPeriodicityAndOffset_r18.PeriodicityAndOffsetExt_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. duration-r18: enumerated
	{
		v2, err := d.DecodeEnumerated(sRSPosUplinkTransmissionWindowConfigR18DurationR18Constraints)
		if err != nil {
			return err
		}
		ie.Duration_r18 = v2
	}

	return nil
}
