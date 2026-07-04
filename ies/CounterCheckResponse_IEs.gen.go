// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CounterCheckResponse-IEs (line 240).

var counterCheckResponseIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-CountInfoList"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var counterCheckResponseIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type CounterCheckResponse_IEs struct {
	Drb_CountInfoList        DRB_CountInfoList
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *CounterCheckResponse_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(counterCheckResponseIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. drb-CountInfoList: ref
	{
		if err := ie.Drb_CountInfoList.Encode(e); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, counterCheckResponseIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *CounterCheckResponse_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(counterCheckResponseIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. drb-CountInfoList: ref
	{
		if err := ie.Drb_CountInfoList.Decode(d); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(counterCheckResponseIEsLateNonCriticalExtensionConstraints)
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
