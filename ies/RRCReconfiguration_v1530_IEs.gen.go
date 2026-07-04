// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RRCReconfiguration-v1530-IEs (line 980).

var rRCReconfigurationV1530IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "masterCellGroup", Optional: true},
		{Name: "fullConfig", Optional: true},
		{Name: "dedicatedNAS-MessageList", Optional: true},
		{Name: "masterKeyUpdate", Optional: true},
		{Name: "dedicatedSIB1-Delivery", Optional: true},
		{Name: "dedicatedSystemInformationDelivery", Optional: true},
		{Name: "otherConfig", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfigurationV1530IEsMasterCellGroupConstraints = per.SizeConstraints{}

const (
	RRCReconfiguration_v1530_IEs_FullConfig_True = 0
)

var rRCReconfigurationV1530IEsFullConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1530_IEs_FullConfig_True},
}

var rRCReconfigurationV1530IEsDedicatedNASMessageListConstraints = per.SizeRange(1, common.MaxDRB)

var rRCReconfigurationV1530IEsDedicatedSIB1DeliveryConstraints = per.SizeConstraints{}

var rRCReconfigurationV1530IEsDedicatedSystemInformationDeliveryConstraints = per.SizeConstraints{}

type RRCReconfiguration_v1530_IEs struct {
	MasterCellGroup                    []byte
	FullConfig                         *int64
	DedicatedNAS_MessageList           []DedicatedNAS_Message
	MasterKeyUpdate                    *MasterKeyUpdate
	DedicatedSIB1_Delivery             []byte
	DedicatedSystemInformationDelivery []byte
	OtherConfig                        *OtherConfig
	NonCriticalExtension               *RRCReconfiguration_v1540_IEs
}

func (ie *RRCReconfiguration_v1530_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1530IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MasterCellGroup != nil, ie.FullConfig != nil, ie.DedicatedNAS_MessageList != nil, ie.MasterKeyUpdate != nil, ie.DedicatedSIB1_Delivery != nil, ie.DedicatedSystemInformationDelivery != nil, ie.OtherConfig != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. masterCellGroup: octet-string
	{
		if ie.MasterCellGroup != nil {
			if err := e.EncodeOctetString(ie.MasterCellGroup, rRCReconfigurationV1530IEsMasterCellGroupConstraints); err != nil {
				return err
			}
		}
	}

	// 3. fullConfig: enumerated
	{
		if ie.FullConfig != nil {
			if err := e.EncodeEnumerated(*ie.FullConfig, rRCReconfigurationV1530IEsFullConfigConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dedicatedNAS-MessageList: sequence-of
	{
		if ie.DedicatedNAS_MessageList != nil {
			seqOf := e.NewSequenceOfEncoder(rRCReconfigurationV1530IEsDedicatedNASMessageListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.DedicatedNAS_MessageList))); err != nil {
				return err
			}
			for i := range ie.DedicatedNAS_MessageList {
				if err := ie.DedicatedNAS_MessageList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. masterKeyUpdate: ref
	{
		if ie.MasterKeyUpdate != nil {
			if err := ie.MasterKeyUpdate.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. dedicatedSIB1-Delivery: octet-string
	{
		if ie.DedicatedSIB1_Delivery != nil {
			if err := e.EncodeOctetString(ie.DedicatedSIB1_Delivery, rRCReconfigurationV1530IEsDedicatedSIB1DeliveryConstraints); err != nil {
				return err
			}
		}
	}

	// 7. dedicatedSystemInformationDelivery: octet-string
	{
		if ie.DedicatedSystemInformationDelivery != nil {
			if err := e.EncodeOctetString(ie.DedicatedSystemInformationDelivery, rRCReconfigurationV1530IEsDedicatedSystemInformationDeliveryConstraints); err != nil {
				return err
			}
		}
	}

	// 8. otherConfig: ref
	{
		if ie.OtherConfig != nil {
			if err := ie.OtherConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1530_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1530IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. masterCellGroup: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1530IEsMasterCellGroupConstraints)
			if err != nil {
				return err
			}
			ie.MasterCellGroup = v
		}
	}

	// 3. fullConfig: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1530IEsFullConfigConstraints)
			if err != nil {
				return err
			}
			ie.FullConfig = &idx
		}
	}

	// 4. dedicatedNAS-MessageList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(rRCReconfigurationV1530IEsDedicatedNASMessageListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DedicatedNAS_MessageList = make([]DedicatedNAS_Message, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DedicatedNAS_MessageList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. masterKeyUpdate: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MasterKeyUpdate = new(MasterKeyUpdate)
			if err := ie.MasterKeyUpdate.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. dedicatedSIB1-Delivery: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1530IEsDedicatedSIB1DeliveryConstraints)
			if err != nil {
				return err
			}
			ie.DedicatedSIB1_Delivery = v
		}
	}

	// 7. dedicatedSystemInformationDelivery: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1530IEsDedicatedSystemInformationDeliveryConstraints)
			if err != nil {
				return err
			}
			ie.DedicatedSystemInformationDelivery = v
		}
	}

	// 8. otherConfig: ref
	{
		if seq.IsComponentPresent(6) {
			ie.OtherConfig = new(OtherConfig)
			if err := ie.OtherConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(7) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1540_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
