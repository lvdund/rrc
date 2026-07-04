// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FrequencyHoppingOffsetListsDCI-0-2-r16 (line 12474).
// FrequencyHoppingOffsetListsDCI-0-2-r16 ::=  SEQUENCE (SIZE (1..4)) OF INTEGER (1.. maxNrofPhysicalResourceBlocks-1)

var frequencyHoppingOffsetListsDCI02R16SizeConstraints = per.SizeRange(1, 4)

type FrequencyHoppingOffsetListsDCI_0_2_r16 []int64

func (ie *FrequencyHoppingOffsetListsDCI_0_2_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(frequencyHoppingOffsetListsDCI02R16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FrequencyHoppingOffsetListsDCI_0_2_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(frequencyHoppingOffsetListsDCI02R16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FrequencyHoppingOffsetListsDCI_0_2_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
