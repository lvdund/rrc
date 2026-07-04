// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DedicatedSIBRequest-r16-IEs (line 265).

var dedicatedSIBRequestR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "onDemandSIB-RequestList-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var dedicatedSIBRequestR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

var dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedSIB-List-r16", Optional: true},
		{Name: "requestedPosSIB-List-r16", Optional: true},
	},
}

var dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedSIBListR16Constraints = per.SizeRange(1, common.MaxOnDemandSIB_r16)

var dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedPosSIBListR16Constraints = per.SizeRange(1, common.MaxOnDemandPosSIB_r16)

type DedicatedSIBRequest_r16_IEs struct {
	OnDemandSIB_RequestList_r16 *struct {
		RequestedSIB_List_r16    []SIB_ReqInfo_r16
		RequestedPosSIB_List_r16 []PosSIB_ReqInfo_r16
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *DedicatedSIBRequest_v1900_IEs
}

func (ie *DedicatedSIBRequest_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dedicatedSIBRequestR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OnDemandSIB_RequestList_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. onDemandSIB-RequestList-r16: sequence
	{
		if ie.OnDemandSIB_RequestList_r16 != nil {
			c := ie.OnDemandSIB_RequestList_r16
			dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq := e.NewSequenceEncoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Constraints)
			if err := dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq.EncodePreamble([]bool{c.RequestedSIB_List_r16 != nil, c.RequestedPosSIB_List_r16 != nil}); err != nil {
				return err
			}
			if c.RequestedSIB_List_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedSIBListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.RequestedSIB_List_r16))); err != nil {
					return err
				}
				for i := range c.RequestedSIB_List_r16 {
					if err := c.RequestedSIB_List_r16[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if c.RequestedPosSIB_List_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedPosSIBListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.RequestedPosSIB_List_r16))); err != nil {
					return err
				}
				for i := range c.RequestedPosSIB_List_r16 {
					if err := c.RequestedPosSIB_List_r16[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, dedicatedSIBRequestR16IEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *DedicatedSIBRequest_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dedicatedSIBRequestR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. onDemandSIB-RequestList-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.OnDemandSIB_RequestList_r16 = &struct {
				RequestedSIB_List_r16    []SIB_ReqInfo_r16
				RequestedPosSIB_List_r16 []PosSIB_ReqInfo_r16
			}{}
			c := ie.OnDemandSIB_RequestList_r16
			dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq := d.NewSequenceDecoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Constraints)
			if err := dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedSIBListR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.RequestedSIB_List_r16 = make([]SIB_ReqInfo_r16, n)
				for i := int64(0); i < n; i++ {
					if err := c.RequestedSIB_List_r16[i].Decode(d); err != nil {
						return err
					}
				}
			}
			if dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16Seq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(dedicatedSIBRequestR16IEsOnDemandSIBRequestListR16RequestedPosSIBListR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.RequestedPosSIB_List_r16 = make([]PosSIB_ReqInfo_r16, n)
				for i := int64(0); i < n; i++ {
					if err := c.RequestedPosSIB_List_r16[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(dedicatedSIBRequestR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(DedicatedSIBRequest_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
