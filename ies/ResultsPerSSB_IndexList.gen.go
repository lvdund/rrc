// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ResultsPerSSB-IndexList (line 9821).
// ResultsPerSSB-IndexList::=              SEQUENCE (SIZE (1..maxNrofIndexesToReport2)) OF ResultsPerSSB-Index

var resultsPerSSBIndexListSizeConstraints = per.SizeRange(1, common.MaxNrofIndexesToReport2)

type ResultsPerSSB_IndexList []ResultsPerSSB_Index

func (ie *ResultsPerSSB_IndexList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(resultsPerSSBIndexListSizeConstraints)
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

func (ie *ResultsPerSSB_IndexList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(resultsPerSSBIndexListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ResultsPerSSB_IndexList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
