// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-NS-PmaxList (line 10744).
// NR-NS-PmaxList ::=                      SEQUENCE (SIZE (1..maxNR-NS-Pmax)) OF NR-NS-PmaxValue

var nRNSPmaxListSizeConstraints = per.SizeRange(1, common.MaxNR_NS_Pmax)

type NR_NS_PmaxList []NR_NS_PmaxValue

func (ie *NR_NS_PmaxList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRNSPmaxListSizeConstraints)
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

func (ie *NR_NS_PmaxList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRNSPmaxListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NR_NS_PmaxList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
