// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSInterestIndication-r17-IEs (line 646).

var mBSInterestIndicationR17IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-FreqList-r17", Optional: true},
		{Name: "mbs-Priority-r17", Optional: true},
		{Name: "mbs-ServiceList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	MBSInterestIndication_r17_IEs_Mbs_Priority_r17_True = 0
)

var mBSInterestIndicationR17IEsMbsPriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MBSInterestIndication_r17_IEs_Mbs_Priority_r17_True},
}

var mBSInterestIndicationR17IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MBSInterestIndication_r17_IEs struct {
	Mbs_FreqList_r17         *CarrierFreqListMBS_r17
	Mbs_Priority_r17         *int64
	Mbs_ServiceList_r17      *MBS_ServiceList_r17
	LateNonCriticalExtension []byte
	NonCriticalExtension     *MBSInterestIndication_v1800
}

func (ie *MBSInterestIndication_r17_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSInterestIndicationR17IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_FreqList_r17 != nil, ie.Mbs_Priority_r17 != nil, ie.Mbs_ServiceList_r17 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mbs-FreqList-r17: ref
	{
		if ie.Mbs_FreqList_r17 != nil {
			if err := ie.Mbs_FreqList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mbs-Priority-r17: enumerated
	{
		if ie.Mbs_Priority_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mbs_Priority_r17, mBSInterestIndicationR17IEsMbsPriorityR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. mbs-ServiceList-r17: ref
	{
		if ie.Mbs_ServiceList_r17 != nil {
			if err := ie.Mbs_ServiceList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, mBSInterestIndicationR17IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBSInterestIndication_r17_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSInterestIndicationR17IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-FreqList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_FreqList_r17 = new(CarrierFreqListMBS_r17)
			if err := ie.Mbs_FreqList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mbs-Priority-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mBSInterestIndicationR17IEsMbsPriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.Mbs_Priority_r17 = &idx
		}
	}

	// 4. mbs-ServiceList-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mbs_ServiceList_r17 = new(MBS_ServiceList_r17)
			if err := ie.Mbs_ServiceList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(mBSInterestIndicationR17IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(MBSInterestIndication_v1800)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
