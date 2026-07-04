// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSBroadcastConfiguration-v1900-IEs (line 631).

var mBSBroadcastConfigurationV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionInfoList-v1900", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type MBSBroadcastConfiguration_v1900_IEs struct {
	Mbs_SessionInfoList_v1900 *MBS_SessionInfoList_v1900
	NonCriticalExtension      *struct{}
}

func (ie *MBSBroadcastConfiguration_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSBroadcastConfigurationV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_SessionInfoList_v1900 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mbs-SessionInfoList-v1900: ref
	{
		if ie.Mbs_SessionInfoList_v1900 != nil {
			if err := ie.Mbs_SessionInfoList_v1900.Encode(e); err != nil {
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

func (ie *MBSBroadcastConfiguration_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSBroadcastConfigurationV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-SessionInfoList-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_SessionInfoList_v1900 = new(MBS_SessionInfoList_v1900)
			if err := ie.Mbs_SessionInfoList_v1900.Decode(d); err != nil {
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
