// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CounterCheck-IEs (line 215).

var counterCheckIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-CountMSB-InfoList"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var counterCheckIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type CounterCheck_IEs struct {
	Drb_CountMSB_InfoList    DRB_CountMSB_InfoList
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *CounterCheck_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(counterCheckIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. drb-CountMSB-InfoList: ref
	{
		if err := ie.Drb_CountMSB_InfoList.Encode(e); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, counterCheckIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *CounterCheck_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(counterCheckIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. drb-CountMSB-InfoList: ref
	{
		if err := ie.Drb_CountMSB_InfoList.Decode(d); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(counterCheckIEsLateNonCriticalExtensionConstraints)
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
