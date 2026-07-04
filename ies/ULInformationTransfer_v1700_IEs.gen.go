// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransfer-v1700-IEs (line 3646).

var uLInformationTransferV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dedicatedInfoF1c-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type ULInformationTransfer_v1700_IEs struct {
	DedicatedInfoF1c_r17 *DedicatedInfoF1c_r17
	NonCriticalExtension *struct{}
}

func (ie *ULInformationTransfer_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DedicatedInfoF1c_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. dedicatedInfoF1c-r17: ref
	{
		if ie.DedicatedInfoF1c_r17 != nil {
			if err := ie.DedicatedInfoF1c_r17.Encode(e); err != nil {
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

func (ie *ULInformationTransfer_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dedicatedInfoF1c-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
			if err := ie.DedicatedInfoF1c_r17.Decode(d); err != nil {
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
