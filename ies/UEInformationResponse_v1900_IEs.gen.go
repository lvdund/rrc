// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationResponse-v1900-IEs (line 2977).

var uEInformationResponseV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-LogMeasReport-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UEInformationResponse_v1900_IEs struct {
	Csi_LogMeasReport_r19 *CSI_LogMeasReport_r19
	NonCriticalExtension  *struct{}
}

func (ie *UEInformationResponse_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationResponseV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_LogMeasReport_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. csi-LogMeasReport-r19: ref
	{
		if ie.Csi_LogMeasReport_r19 != nil {
			if err := ie.Csi_LogMeasReport_r19.Encode(e); err != nil {
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

func (ie *UEInformationResponse_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationResponseV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-LogMeasReport-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Csi_LogMeasReport_r19 = new(CSI_LogMeasReport_r19)
			if err := ie.Csi_LogMeasReport_r19.Decode(d); err != nil {
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
