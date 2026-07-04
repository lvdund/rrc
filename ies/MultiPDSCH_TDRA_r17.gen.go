// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiPDSCH-TDRA-r17 (line 11541).

var multiPDSCHTDRAR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-TDRA-List-r17"},
	},
}

var multiPDSCHTDRAR17PdschTDRAListR17Constraints = per.SizeRange(1, common.MaxNrofMultiplePDSCHs_r17)

type MultiPDSCH_TDRA_r17 struct {
	Pdsch_TDRA_List_r17 []PDSCH_TimeDomainResourceAllocation_r16
}

func (ie *MultiPDSCH_TDRA_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multiPDSCHTDRAR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. pdsch-TDRA-List-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(multiPDSCHTDRAR17PdschTDRAListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Pdsch_TDRA_List_r17))); err != nil {
			return err
		}
		for i := range ie.Pdsch_TDRA_List_r17 {
			if err := ie.Pdsch_TDRA_List_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MultiPDSCH_TDRA_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multiPDSCHTDRAR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. pdsch-TDRA-List-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(multiPDSCHTDRAR17PdschTDRAListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Pdsch_TDRA_List_r17 = make([]PDSCH_TimeDomainResourceAllocation_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Pdsch_TDRA_List_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
