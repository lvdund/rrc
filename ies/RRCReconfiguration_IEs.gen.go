// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-IEs (line 971).

var rRCReconfigurationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "radioBearerConfig", Optional: true},
		{Name: "secondaryCellGroup", Optional: true},
		{Name: "measConfig", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfigurationIEsSecondaryCellGroupConstraints = per.SizeConstraints{}

var rRCReconfigurationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCReconfiguration_IEs struct {
	RadioBearerConfig        *RadioBearerConfig
	SecondaryCellGroup       []byte
	MeasConfig               *MeasConfig
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCReconfiguration_v1530_IEs
}

func (ie *RRCReconfiguration_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RadioBearerConfig != nil, ie.SecondaryCellGroup != nil, ie.MeasConfig != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if ie.RadioBearerConfig != nil {
			if err := ie.RadioBearerConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. secondaryCellGroup: octet-string
	{
		if ie.SecondaryCellGroup != nil {
			if err := e.EncodeOctetString(ie.SecondaryCellGroup, rRCReconfigurationIEsSecondaryCellGroupConstraints); err != nil {
				return err
			}
		}
	}

	// 4. measConfig: ref
	{
		if ie.MeasConfig != nil {
			if err := ie.MeasConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCReconfigurationIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RadioBearerConfig = new(RadioBearerConfig)
			if err := ie.RadioBearerConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. secondaryCellGroup: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(rRCReconfigurationIEsSecondaryCellGroupConstraints)
			if err != nil {
				return err
			}
			ie.SecondaryCellGroup = v
		}
	}

	// 4. measConfig: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasConfig = new(MeasConfig)
			if err := ie.MeasConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(rRCReconfigurationIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1530_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
