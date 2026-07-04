// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TCI-ActivatedConfig-r17 (line 15982).

var tCIActivatedConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-TCI-r17"},
		{Name: "pdsch-TCI-r17"},
	},
}

var tCIActivatedConfigR17PdcchTCIR17Constraints = per.SizeRange(1, 5)

var tCIActivatedConfigR17PdschTCIR17Constraints = per.SizeRange(1, common.MaxNrofTCI_States)

type TCI_ActivatedConfig_r17 struct {
	Pdcch_TCI_r17 []TCI_StateId
	Pdsch_TCI_r17 per.BitString
}

func (ie *TCI_ActivatedConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tCIActivatedConfigR17Constraints)
	_ = seq

	// 1. pdcch-TCI-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(tCIActivatedConfigR17PdcchTCIR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Pdcch_TCI_r17))); err != nil {
			return err
		}
		for i := range ie.Pdcch_TCI_r17 {
			if err := ie.Pdcch_TCI_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. pdsch-TCI-r17: bit-string
	{
		if err := e.EncodeBitString(ie.Pdsch_TCI_r17, tCIActivatedConfigR17PdschTCIR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TCI_ActivatedConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tCIActivatedConfigR17Constraints)
	_ = seq

	// 1. pdcch-TCI-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(tCIActivatedConfigR17PdcchTCIR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Pdcch_TCI_r17 = make([]TCI_StateId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Pdcch_TCI_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. pdsch-TCI-r17: bit-string
	{
		v1, err := d.DecodeBitString(tCIActivatedConfigR17PdschTCIR17Constraints)
		if err != nil {
			return err
		}
		ie.Pdsch_TCI_r17 = v1
	}

	return nil
}
