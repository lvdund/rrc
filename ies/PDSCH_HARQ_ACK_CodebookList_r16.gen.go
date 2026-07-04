// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-HARQ-ACK-CodebookList-r16 (line 11738).
// PDSCH-HARQ-ACK-CodebookList-r16 ::=     SEQUENCE (SIZE (1..2)) OF ENUMERATED {semiStatic, dynamic}

var pDSCHHARQACKCodebookListR16SizeConstraints = per.SizeRange(1, 2)

const (
	PDSCH_HARQ_ACK_CodebookList_r16_PDSCH_HARQ_ACK_CodebookList_r16_Elem_SemiStatic = 0
	PDSCH_HARQ_ACK_CodebookList_r16_PDSCH_HARQ_ACK_CodebookList_r16_Elem_Dynamic    = 1
)

var pDSCHHARQACKCodebookListR16PDSCHHARQACKCodebookListR16ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_HARQ_ACK_CodebookList_r16_PDSCH_HARQ_ACK_CodebookList_r16_Elem_SemiStatic, PDSCH_HARQ_ACK_CodebookList_r16_PDSCH_HARQ_ACK_CodebookList_r16_Elem_Dynamic},
}

type PDSCH_HARQ_ACK_CodebookList_r16 []int64

func (ie *PDSCH_HARQ_ACK_CodebookList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDSCHHARQACKCodebookListR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeEnumerated((*ie)[i], pDSCHHARQACKCodebookListR16PDSCHHARQACKCodebookListR16ElemConstraints); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDSCH_HARQ_ACK_CodebookList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDSCHHARQACKCodebookListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PDSCH_HARQ_ACK_CodebookList_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeEnumerated(pDSCHHARQACKCodebookListR16PDSCHHARQACKCodebookListR16ElemConstraints)
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
