// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransfer-v1610-IEs (line 352).

var dLInformationTransferV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceTimeInfo-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type DLInformationTransfer_v1610_IEs struct {
	ReferenceTimeInfo_r16 *ReferenceTimeInfo_r16
	NonCriticalExtension  *DLInformationTransfer_v1700_IEs
}

func (ie *DLInformationTransfer_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReferenceTimeInfo_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. referenceTimeInfo-r16: ref
	{
		if ie.ReferenceTimeInfo_r16 != nil {
			if err := ie.ReferenceTimeInfo_r16.Encode(e); err != nil {
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

func (ie *DLInformationTransfer_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. referenceTimeInfo-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ReferenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
			if err := ie.ReferenceTimeInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(DLInformationTransfer_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
