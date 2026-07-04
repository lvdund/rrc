// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB1-v1630-IEs (line 2016).

var sIB1V1630IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uac-BarringInfo-v1630", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var sIB1V1630IEsUacBarringInfoV1630UacAC1SelectAssistInfoR16Constraints = per.SizeRange(2, common.MaxPLMN)

type SIB1_v1630_IEs struct {
	Uac_BarringInfo_v1630 *struct {
		Uac_AC1_SelectAssistInfo_r16 []UAC_AC1_SelectAssistInfo_r16
	}
	NonCriticalExtension *SIB1_v1700_IEs
}

func (ie *SIB1_v1630_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1V1630IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uac_BarringInfo_v1630 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. uac-BarringInfo-v1630: sequence
	{
		if ie.Uac_BarringInfo_v1630 != nil {
			c := ie.Uac_BarringInfo_v1630
			{
				seqOf := e.NewSequenceOfEncoder(sIB1V1630IEsUacBarringInfoV1630UacAC1SelectAssistInfoR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Uac_AC1_SelectAssistInfo_r16))); err != nil {
					return err
				}
				for i := range c.Uac_AC1_SelectAssistInfo_r16 {
					if err := c.Uac_AC1_SelectAssistInfo_r16[i].Encode(e); err != nil {
						return err
					}
				}
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

func (ie *SIB1_v1630_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1V1630IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uac-BarringInfo-v1630: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Uac_BarringInfo_v1630 = &struct {
				Uac_AC1_SelectAssistInfo_r16 []UAC_AC1_SelectAssistInfo_r16
			}{}
			c := ie.Uac_BarringInfo_v1630
			{
				seqOf := d.NewSequenceOfDecoder(sIB1V1630IEsUacBarringInfoV1630UacAC1SelectAssistInfoR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Uac_AC1_SelectAssistInfo_r16 = make([]UAC_AC1_SelectAssistInfo_r16, n)
				for i := int64(0); i < n; i++ {
					if err := c.Uac_AC1_SelectAssistInfo_r16[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(SIB1_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
