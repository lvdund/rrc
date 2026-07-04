// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiPDSCH-TDRA-List-r17 (line 11539).
// MultiPDSCH-TDRA-List-r17 ::= SEQUENCE (SIZE(1.. maxNrofDL-AllocationsExt-r17)) OF MultiPDSCH-TDRA-r17

var multiPDSCHTDRAListR17SizeConstraints = per.SizeRange(1, common.MaxNrofDL_AllocationsExt_r17)

type MultiPDSCH_TDRA_List_r17 []MultiPDSCH_TDRA_r17

func (ie *MultiPDSCH_TDRA_List_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiPDSCHTDRAListR17SizeConstraints)
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

func (ie *MultiPDSCH_TDRA_List_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiPDSCHTDRAListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiPDSCH_TDRA_List_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
