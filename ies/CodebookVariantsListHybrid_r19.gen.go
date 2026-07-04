// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookVariantsListHybrid-r19 (line 19332).
// CodebookVariantsListHybrid-r19 ::=    SEQUENCE (SIZE (1..maxNrofCSI-RS-ResourcesAlt-r16)) OF SupportedCSI-RS-ResourceHybrid-r19

var codebookVariantsListHybridR19SizeConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesAlt_r16)

type CodebookVariantsListHybrid_r19 []SupportedCSI_RS_ResourceHybrid_r19

func (ie *CodebookVariantsListHybrid_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(codebookVariantsListHybridR19SizeConstraints)
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

func (ie *CodebookVariantsListHybrid_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(codebookVariantsListHybridR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CodebookVariantsListHybrid_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
