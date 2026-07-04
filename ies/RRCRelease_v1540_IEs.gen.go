// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease-v1540-IEs (line 1241).

var rRCReleaseV1540IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "waitTime", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCRelease_v1540_IEs struct {
	WaitTime             *RejectWaitTime
	NonCriticalExtension *RRCRelease_v1610_IEs
}

func (ie *RRCRelease_v1540_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseV1540IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.WaitTime != nil, ie.NonCriticalExtension != nil}); err != nil {
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

func (ie *RRCRelease_v1540_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseV1540IEsConstraints)

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

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCRelease_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
