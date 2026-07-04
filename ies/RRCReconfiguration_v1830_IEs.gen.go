// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1830-IEs (line 1045).

var rRCReconfigurationV1830IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "otherConfig-v1830", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReconfiguration_v1830_IEs struct {
	OtherConfig_v1830    *OtherConfig_v1830
	NonCriticalExtension *RRCReconfiguration_v1900_IEs
}

func (ie *RRCReconfiguration_v1830_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1830IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OtherConfig_v1830 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. otherConfig-v1830: ref
	{
		if ie.OtherConfig_v1830 != nil {
			if err := ie.OtherConfig_v1830.Encode(e); err != nil {
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

func (ie *RRCReconfiguration_v1830_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1830IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. otherConfig-v1830: ref
	{
		if seq.IsComponentPresent(0) {
			ie.OtherConfig_v1830 = new(OtherConfig_v1830)
			if err := ie.OtherConfig_v1830.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
