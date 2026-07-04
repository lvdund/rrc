// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-IEs (line 1144).

var rRCReconfigurationCompleteIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfigurationCompleteIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCReconfigurationComplete_IEs struct {
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCReconfigurationComplete_v1530_IEs
}

func (ie *RRCReconfigurationComplete_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCReconfigurationCompleteIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(rRCReconfigurationCompleteIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1530_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
