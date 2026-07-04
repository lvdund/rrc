// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-NS-PmaxList-v1760 (line 10751).
// NR-NS-PmaxList-v1760 ::=                SEQUENCE (SIZE (1.. maxNR-NS-Pmax)) OF NR-NS-PmaxValue-v1760

var nRNSPmaxListV1760SizeConstraints = per.SizeRange(1, common.MaxNR_NS_Pmax)

type NR_NS_PmaxList_v1760 []NR_NS_PmaxValue_v1760

func (ie *NR_NS_PmaxList_v1760) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRNSPmaxListV1760SizeConstraints)
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

func (ie *NR_NS_PmaxList_v1760) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRNSPmaxListV1760SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NR_NS_PmaxList_v1760, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
