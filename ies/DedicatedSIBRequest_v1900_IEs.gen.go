// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DedicatedSIBRequest-v1900-IEs (line 274).

var dedicatedSIBRequestV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "onDemandPosSIB-RequestList-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedPeriodicAD-PosSIB-List-r19", Optional: true},
	},
}

var dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19RequestedPeriodicADPosSIBListR19Constraints = per.SizeRange(1, common.MaxOnDemandPosSIB_r16)

type DedicatedSIBRequest_v1900_IEs struct {
	OnDemandPosSIB_RequestList_r19 *struct {
		RequestedPeriodicAD_PosSIB_List_r19 []PeriodicAD_PosSIB_ReqInfo_r19
	}
	NonCriticalExtension *struct{}
}

func (ie *DedicatedSIBRequest_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dedicatedSIBRequestV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OnDemandPosSIB_RequestList_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. onDemandPosSIB-RequestList-r19: sequence
	{
		if ie.OnDemandPosSIB_RequestList_r19 != nil {
			c := ie.OnDemandPosSIB_RequestList_r19
			dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Seq := e.NewSequenceEncoder(dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Constraints)
			if err := dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Seq.EncodePreamble([]bool{c.RequestedPeriodicAD_PosSIB_List_r19 != nil}); err != nil {
				return err
			}
			if c.RequestedPeriodicAD_PosSIB_List_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19RequestedPeriodicADPosSIBListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.RequestedPeriodicAD_PosSIB_List_r19))); err != nil {
					return err
				}
				for i := range c.RequestedPeriodicAD_PosSIB_List_r19 {
					if err := c.RequestedPeriodicAD_PosSIB_List_r19[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *DedicatedSIBRequest_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dedicatedSIBRequestV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. onDemandPosSIB-RequestList-r19: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.OnDemandPosSIB_RequestList_r19 = &struct {
				RequestedPeriodicAD_PosSIB_List_r19 []PeriodicAD_PosSIB_ReqInfo_r19
			}{}
			c := ie.OnDemandPosSIB_RequestList_r19
			dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Seq := d.NewSequenceDecoder(dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Constraints)
			if err := dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(dedicatedSIBRequestV1900IEsOnDemandPosSIBRequestListR19RequestedPeriodicADPosSIBListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.RequestedPeriodicAD_PosSIB_List_r19 = make([]PeriodicAD_PosSIB_ReqInfo_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.RequestedPeriodicAD_PosSIB_List_r19[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
