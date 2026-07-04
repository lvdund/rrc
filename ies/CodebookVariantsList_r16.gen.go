// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookVariantsList-r16 (line 19329).
// CodebookVariantsList-r16 ::=          SEQUENCE (SIZE (1..maxNrofCSI-RS-ResourcesAlt-r16)) OF SupportedCSI-RS-Resource

var codebookVariantsListR16SizeConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesAlt_r16)

type CodebookVariantsList_r16 []SupportedCSI_RS_Resource

func (ie *CodebookVariantsList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(codebookVariantsListR16SizeConstraints)
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

func (ie *CodebookVariantsList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(codebookVariantsListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CodebookVariantsList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
