// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Tx-PoolMeasList-r16 (line 9627).
// Tx-PoolMeasList-r16 ::= SEQUENCE (SIZE (1..maxNrofSL-PoolToMeasureNR-r16)) OF SL-ResourcePoolID-r16

var txPoolMeasListR16SizeConstraints = per.SizeRange(1, common.MaxNrofSL_PoolToMeasureNR_r16)

type Tx_PoolMeasList_r16 []SL_ResourcePoolID_r16

func (ie *Tx_PoolMeasList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(txPoolMeasListR16SizeConstraints)
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

func (ie *Tx_PoolMeasList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(txPoolMeasListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(Tx_PoolMeasList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
