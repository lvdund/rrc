// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation-v1540-IEs (line 2410).

var uEAssistanceInformationV1540IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "overheatingAssistance", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UEAssistanceInformation_v1540_IEs struct {
	OverheatingAssistance *OverheatingAssistance
	NonCriticalExtension  *UEAssistanceInformation_v1610_IEs
}

func (ie *UEAssistanceInformation_v1540_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationV1540IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OverheatingAssistance != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. overheatingAssistance: ref
	{
		if ie.OverheatingAssistance != nil {
			if err := ie.OverheatingAssistance.Encode(e); err != nil {
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

func (ie *UEAssistanceInformation_v1540_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationV1540IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. overheatingAssistance: ref
	{
		if seq.IsComponentPresent(0) {
			ie.OverheatingAssistance = new(OverheatingAssistance)
			if err := ie.OverheatingAssistance.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UEAssistanceInformation_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
