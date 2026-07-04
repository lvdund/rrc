// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PagingRecord (line 864).

var pagingRecordConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-Identity"},
		{Name: "accessType", Optional: true},
	},
}

const (
	PagingRecord_AccessType_Non3GPP = 0
)

var pagingRecordAccessTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PagingRecord_AccessType_Non3GPP},
}

type PagingRecord struct {
	Ue_Identity PagingUE_Identity
	AccessType  *int64
}

func (ie *PagingRecord) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pagingRecordConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AccessType != nil}); err != nil {
		return err
	}

	// 3. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 4. accessType: enumerated
	{
		if ie.AccessType != nil {
			if err := e.EncodeEnumerated(*ie.AccessType, pagingRecordAccessTypeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PagingRecord) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pagingRecordConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ue-Identity: ref
	{
		if err := ie.Ue_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 4. accessType: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pagingRecordAccessTypeConstraints)
			if err != nil {
				return err
			}
			ie.AccessType = &idx
		}
	}

	return nil
}
