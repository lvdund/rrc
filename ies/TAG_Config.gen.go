// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TAG-Config (line 15950).

var tAGConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tag-ToReleaseList", Optional: true},
		{Name: "tag-ToAddModList", Optional: true},
	},
}

var tAGConfigTagToReleaseListConstraints = per.SizeRange(1, common.MaxNrofTAGs)

var tAGConfigTagToAddModListConstraints = per.SizeRange(1, common.MaxNrofTAGs)

type TAG_Config struct {
	Tag_ToReleaseList []TAG_Id
	Tag_ToAddModList  []TAG
}

func (ie *TAG_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAGConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tag_ToReleaseList != nil, ie.Tag_ToAddModList != nil}); err != nil {
		return err
	}

	// 2. tag-ToReleaseList: sequence-of
	{
		if ie.Tag_ToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(tAGConfigTagToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tag_ToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Tag_ToReleaseList {
				if err := ie.Tag_ToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. tag-ToAddModList: sequence-of
	{
		if ie.Tag_ToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(tAGConfigTagToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tag_ToAddModList))); err != nil {
				return err
			}
			for i := range ie.Tag_ToAddModList {
				if err := ie.Tag_ToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *TAG_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAGConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tag-ToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(tAGConfigTagToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tag_ToReleaseList = make([]TAG_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tag_ToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. tag-ToAddModList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(tAGConfigTagToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tag_ToAddModList = make([]TAG, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tag_ToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
