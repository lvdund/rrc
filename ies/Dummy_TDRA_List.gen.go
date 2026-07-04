// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Dummy-TDRA-List (line 11537).
// Dummy-TDRA-List ::= SEQUENCE (SIZE(1.. maxNrofDL-Allocations)) OF MultiPDSCH-TDRA-r17

var dummyTDRAListSizeConstraints = per.SizeRange(1, common.MaxNrofDL_Allocations)

type Dummy_TDRA_List []MultiPDSCH_TDRA_r17

func (ie *Dummy_TDRA_List) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dummyTDRAListSizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *Dummy_TDRA_List) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dummyTDRAListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(Dummy_TDRA_List, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
