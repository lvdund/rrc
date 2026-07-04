// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ResultsPerCSI-RS-IndexList (line 9828).
// ResultsPerCSI-RS-IndexList::=           SEQUENCE (SIZE (1..maxNrofIndexesToReport2)) OF ResultsPerCSI-RS-Index

var resultsPerCSIRSIndexListSizeConstraints = per.SizeRange(1, common.MaxNrofIndexesToReport2)

type ResultsPerCSI_RS_IndexList []ResultsPerCSI_RS_Index

func (ie *ResultsPerCSI_RS_IndexList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(resultsPerCSIRSIndexListSizeConstraints)
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

func (ie *ResultsPerCSI_RS_IndexList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(resultsPerCSIRSIndexListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ResultsPerCSI_RS_IndexList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
