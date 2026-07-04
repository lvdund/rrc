// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-FSAI-InterFreq-r17 (line 4602).

var mBSFSAIInterFreqR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-CarrierFreq-r17"},
		{Name: "mbs-FSAI-List-r17"},
	},
}

type MBS_FSAI_InterFreq_r17 struct {
	Dl_CarrierFreq_r17 ARFCN_ValueNR
	Mbs_FSAI_List_r17  MBS_FSAI_List_r17
}

func (ie *MBS_FSAI_InterFreq_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSFSAIInterFreqR17Constraints)
	_ = seq

	// 1. dl-CarrierFreq-r17: ref
	{
		if err := ie.Dl_CarrierFreq_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. mbs-FSAI-List-r17: ref
	{
		if err := ie.Mbs_FSAI_List_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MBS_FSAI_InterFreq_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSFSAIInterFreqR17Constraints)
	_ = seq

	// 1. dl-CarrierFreq-r17: ref
	{
		if err := ie.Dl_CarrierFreq_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. mbs-FSAI-List-r17: ref
	{
		if err := ie.Mbs_FSAI_List_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
