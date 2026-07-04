// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCGFailureInformationEUTRA-v1590-IEs (line 1902).

var sCGFailureInformationEUTRAV1590IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var sCGFailureInformationEUTRAV1590IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SCGFailureInformationEUTRA_v1590_IEs struct {
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *SCGFailureInformationEUTRA_v1590_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGFailureInformationEUTRAV1590IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sCGFailureInformationEUTRAV1590IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *SCGFailureInformationEUTRA_v1590_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGFailureInformationEUTRAV1590IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(sCGFailureInformationEUTRAV1590IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
