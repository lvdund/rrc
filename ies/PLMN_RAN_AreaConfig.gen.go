// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PLMN-RAN-AreaConfig (line 1359).

var pLMNRANAreaConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity", Optional: true},
		{Name: "ran-Area"},
	},
}

var pLMNRANAreaConfigRanAreaConstraints = per.SizeRange(1, 16)

type PLMN_RAN_AreaConfig struct {
	Plmn_Identity *PLMN_Identity
	Ran_Area      []RAN_AreaConfig
}

func (ie *PLMN_RAN_AreaConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pLMNRANAreaConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Plmn_Identity != nil}); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if ie.Plmn_Identity != nil {
			if err := ie.Plmn_Identity.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ran-Area: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pLMNRANAreaConfigRanAreaConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ran_Area))); err != nil {
			return err
		}
		for i := range ie.Ran_Area {
			if err := ie.Ran_Area[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PLMN_RAN_AreaConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pLMNRANAreaConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Plmn_Identity = new(PLMN_Identity)
			if err := ie.Plmn_Identity.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ran-Area: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pLMNRANAreaConfigRanAreaConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ran_Area = make([]RAN_AreaConfig, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ran_Area[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
