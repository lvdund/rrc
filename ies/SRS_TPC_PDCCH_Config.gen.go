// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-TPC-PDCCH-Config (line 15298).

var sRSTPCPDCCHConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-CC-SetIndexlist", Optional: true},
	},
}

var sRSTPCPDCCHConfigSrsCCSetIndexlistConstraints = per.SizeRange(1, 4)

type SRS_TPC_PDCCH_Config struct {
	Srs_CC_SetIndexlist []SRS_CC_SetIndex
}

func (ie *SRS_TPC_PDCCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSTPCPDCCHConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_CC_SetIndexlist != nil}); err != nil {
		return err
	}

	// 2. srs-CC-SetIndexlist: sequence-of
	{
		if ie.Srs_CC_SetIndexlist != nil {
			seqOf := e.NewSequenceOfEncoder(sRSTPCPDCCHConfigSrsCCSetIndexlistConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_CC_SetIndexlist))); err != nil {
				return err
			}
			for i := range ie.Srs_CC_SetIndexlist {
				if err := ie.Srs_CC_SetIndexlist[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SRS_TPC_PDCCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSTPCPDCCHConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-CC-SetIndexlist: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sRSTPCPDCCHConfigSrsCCSetIndexlistConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_CC_SetIndexlist = make([]SRS_CC_SetIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_CC_SetIndexlist[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
