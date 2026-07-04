// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LocationMeasurementIndication-IEs (line 525).

var locationMeasurementIndicationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementIndication"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var locationMeasurementIndication_IEsMeasurementIndicationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LocationMeasurementIndication_IEs_MeasurementIndication_Release = 0
	LocationMeasurementIndication_IEs_MeasurementIndication_Setup   = 1
)

var locationMeasurementIndicationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type LocationMeasurementIndication_IEs struct {
	MeasurementIndication struct {
		Choice  int
		Release *struct{}
		Setup   *LocationMeasurementInfo
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *LocationMeasurementIndication_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(locationMeasurementIndicationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measurementIndication: choice
	{
		choiceEnc := e.NewChoiceEncoder(locationMeasurementIndication_IEsMeasurementIndicationConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.MeasurementIndication.Choice), false, nil); err != nil {
			return err
		}
		switch ie.MeasurementIndication.Choice {
		case LocationMeasurementIndication_IEs_MeasurementIndication_Release:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case LocationMeasurementIndication_IEs_MeasurementIndication_Setup:
			if err := ie.MeasurementIndication.Setup.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.MeasurementIndication.Choice), Constraint: "index out of range"}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, locationMeasurementIndicationIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *LocationMeasurementIndication_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(locationMeasurementIndicationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measurementIndication: choice
	{
		choiceDec := d.NewChoiceDecoder(locationMeasurementIndication_IEsMeasurementIndicationConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.MeasurementIndication.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LocationMeasurementIndication_IEs_MeasurementIndication_Release:
			ie.MeasurementIndication.Release = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case LocationMeasurementIndication_IEs_MeasurementIndication_Setup:
			ie.MeasurementIndication.Setup = new(LocationMeasurementInfo)
			if err := ie.MeasurementIndication.Setup.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(locationMeasurementIndicationIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
