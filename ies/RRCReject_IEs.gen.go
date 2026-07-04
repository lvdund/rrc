// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReject-IEs (line 1212).

var rRCRejectIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "waitTime", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCRejectIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCReject_IEs struct {
	WaitTime                 *RejectWaitTime
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *RRCReject_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCRejectIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.WaitTime != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. waitTime: ref
	{
		if ie.WaitTime != nil {
			if err := ie.WaitTime.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCRejectIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *RRCReject_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCRejectIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. waitTime: ref
	{
		if seq.IsComponentPresent(0) {
			ie.WaitTime = new(RejectWaitTime)
			if err := ie.WaitTime.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(rRCRejectIEsLateNonCriticalExtensionConstraints)
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
