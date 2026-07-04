// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Paging-v1800-IEs (line 848).

var pagingV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pagingRecordList-v1800", Optional: true},
		{Name: "pagingGroupList-v1800", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type Paging_v1800_IEs struct {
	PagingRecordList_v1800 *PagingRecordList_v1800
	PagingGroupList_v1800  *PagingGroupList_v1800
	NonCriticalExtension   *struct{}
}

func (ie *Paging_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pagingV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PagingRecordList_v1800 != nil, ie.PagingGroupList_v1800 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. pagingRecordList-v1800: ref
	{
		if ie.PagingRecordList_v1800 != nil {
			if err := ie.PagingRecordList_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pagingGroupList-v1800: ref
	{
		if ie.PagingGroupList_v1800 != nil {
			if err := ie.PagingGroupList_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *Paging_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pagingV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pagingRecordList-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PagingRecordList_v1800 = new(PagingRecordList_v1800)
			if err := ie.PagingRecordList_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pagingGroupList-v1800: ref
	{
		if seq.IsComponentPresent(1) {
			ie.PagingGroupList_v1800 = new(PagingGroupList_v1800)
			if err := ie.PagingGroupList_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
