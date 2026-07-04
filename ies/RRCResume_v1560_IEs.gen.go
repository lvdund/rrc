// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-v1560-IEs (line 1546).

var rRCResumeV1560IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "radioBearerConfig2", Optional: true},
		{Name: "sk-Counter", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCResumeV1560IEsRadioBearerConfig2Constraints = per.SizeConstraints{}

type RRCResume_v1560_IEs struct {
	RadioBearerConfig2   []byte
	Sk_Counter           *SK_Counter
	NonCriticalExtension *RRCResume_v1610_IEs
}

func (ie *RRCResume_v1560_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeV1560IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RadioBearerConfig2 != nil, ie.Sk_Counter != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. radioBearerConfig2: octet-string
	{
		if ie.RadioBearerConfig2 != nil {
			if err := e.EncodeOctetString(ie.RadioBearerConfig2, rRCResumeV1560IEsRadioBearerConfig2Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sk-Counter: ref
	{
		if ie.Sk_Counter != nil {
			if err := ie.Sk_Counter.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResume_v1560_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeV1560IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. radioBearerConfig2: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(rRCResumeV1560IEsRadioBearerConfig2Constraints)
			if err != nil {
				return err
			}
			ie.RadioBearerConfig2 = v
		}
	}

	// 3. sk-Counter: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sk_Counter = new(SK_Counter)
			if err := ie.Sk_Counter.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(RRCResume_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
