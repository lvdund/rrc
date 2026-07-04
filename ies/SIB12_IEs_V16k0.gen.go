// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB12-IEs-v16k0 (line 4311).

var sIB12IEsV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigCommonNR-v16k0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type SIB12_IEs_V16k0 struct {
	Sl_ConfigCommonNR_V16k0 *SL_ConfigCommonNR_V16k0
	NonCriticalExtension    *struct{}
}

func (ie *SIB12_IEs_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB12IEsV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfigCommonNR_V16k0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-ConfigCommonNR-v16k0: ref
	{
		if ie.Sl_ConfigCommonNR_V16k0 != nil {
			if err := ie.Sl_ConfigCommonNR_V16k0.Encode(e); err != nil {
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

func (ie *SIB12_IEs_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB12IEsV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfigCommonNR-v16k0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_ConfigCommonNR_V16k0 = new(SL_ConfigCommonNR_V16k0)
			if err := ie.Sl_ConfigCommonNR_V16k0.Decode(d); err != nil {
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
