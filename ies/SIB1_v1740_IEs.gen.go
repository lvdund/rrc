// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB1-v1740-IEs (line 2045).

var sIB1V1740IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "si-SchedulingInfo-v1740", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type SIB1_v1740_IEs struct {
	Si_SchedulingInfo_v1740 *SI_SchedulingInfo_v1740
	NonCriticalExtension    *SIB1_v1800_IEs
}

func (ie *SIB1_v1740_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1V1740IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Si_SchedulingInfo_v1740 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. si-SchedulingInfo-v1740: ref
	{
		if ie.Si_SchedulingInfo_v1740 != nil {
			if err := ie.Si_SchedulingInfo_v1740.Encode(e); err != nil {
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

func (ie *SIB1_v1740_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1V1740IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. si-SchedulingInfo-v1740: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Si_SchedulingInfo_v1740 = new(SI_SchedulingInfo_v1740)
			if err := ie.Si_SchedulingInfo_v1740.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(SIB1_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
