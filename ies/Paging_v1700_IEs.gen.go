// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Paging-v1700-IEs (line 842).

var pagingV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pagingRecordList-v1700", Optional: true},
		{Name: "pagingGroupList-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type Paging_v1700_IEs struct {
	PagingRecordList_v1700 *PagingRecordList_v1700
	PagingGroupList_r17    *PagingGroupList_r17
	NonCriticalExtension   *Paging_v1800_IEs
}

func (ie *Paging_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pagingV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PagingRecordList_v1700 != nil, ie.PagingGroupList_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. pagingRecordList-v1700: ref
	{
		if ie.PagingRecordList_v1700 != nil {
			if err := ie.PagingRecordList_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pagingGroupList-r17: ref
	{
		if ie.PagingGroupList_r17 != nil {
			if err := ie.PagingGroupList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Paging_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pagingV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pagingRecordList-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PagingRecordList_v1700 = new(PagingRecordList_v1700)
			if err := ie.PagingRecordList_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pagingGroupList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.PagingGroupList_r17 = new(PagingGroupList_r17)
			if err := ie.PagingGroupList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(Paging_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
