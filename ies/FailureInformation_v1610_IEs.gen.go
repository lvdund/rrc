// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureInformation-v1610-IEs (line 416).

var failureInformationV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureInfoDAPS-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type FailureInformation_v1610_IEs struct {
	FailureInfoDAPS_r16  *FailureInfoDAPS_r16
	NonCriticalExtension *struct{}
}

func (ie *FailureInformation_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureInformationV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureInfoDAPS_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureInfoDAPS-r16: ref
	{
		if ie.FailureInfoDAPS_r16 != nil {
			if err := ie.FailureInfoDAPS_r16.Encode(e); err != nil {
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

func (ie *FailureInformation_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureInformationV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureInfoDAPS-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureInfoDAPS_r16 = new(FailureInfoDAPS_r16)
			if err := ie.FailureInfoDAPS_r16.Decode(d); err != nil {
				return err
			}
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
