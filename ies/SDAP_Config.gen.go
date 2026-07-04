// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SDAP-Config (line 14357).

var sDAPConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-Session"},
		{Name: "sdap-HeaderDL"},
		{Name: "sdap-HeaderUL"},
		{Name: "defaultDRB"},
		{Name: "mappedQoS-FlowsToAdd", Optional: true},
		{Name: "mappedQoS-FlowsToRelease", Optional: true},
	},
}

const (
	SDAP_Config_Sdap_HeaderDL_Present = 0
	SDAP_Config_Sdap_HeaderDL_Absent  = 1
)

var sDAPConfigSdapHeaderDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Config_Sdap_HeaderDL_Present, SDAP_Config_Sdap_HeaderDL_Absent},
}

const (
	SDAP_Config_Sdap_HeaderUL_Present = 0
	SDAP_Config_Sdap_HeaderUL_Absent  = 1
)

var sDAPConfigSdapHeaderULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Config_Sdap_HeaderUL_Present, SDAP_Config_Sdap_HeaderUL_Absent},
}

var sDAPConfigMappedQoSFlowsToAddConstraints = per.SizeRange(1, common.MaxNrofQFIs)

var sDAPConfigMappedQoSFlowsToReleaseConstraints = per.SizeRange(1, common.MaxNrofQFIs)

type SDAP_Config struct {
	Pdu_Session              PDU_SessionID
	Sdap_HeaderDL            int64
	Sdap_HeaderUL            int64
	DefaultDRB               bool
	MappedQoS_FlowsToAdd     []QFI
	MappedQoS_FlowsToRelease []QFI
}

func (ie *SDAP_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDAPConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MappedQoS_FlowsToAdd != nil, ie.MappedQoS_FlowsToRelease != nil}); err != nil {
		return err
	}

	// 3. pdu-Session: ref
	{
		if err := ie.Pdu_Session.Encode(e); err != nil {
			return err
		}
	}

	// 4. sdap-HeaderDL: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sdap_HeaderDL, sDAPConfigSdapHeaderDLConstraints); err != nil {
			return err
		}
	}

	// 5. sdap-HeaderUL: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sdap_HeaderUL, sDAPConfigSdapHeaderULConstraints); err != nil {
			return err
		}
	}

	// 6. defaultDRB: boolean
	{
		if err := e.EncodeBoolean(ie.DefaultDRB); err != nil {
			return err
		}
	}

	// 7. mappedQoS-FlowsToAdd: sequence-of
	{
		if ie.MappedQoS_FlowsToAdd != nil {
			seqOf := e.NewSequenceOfEncoder(sDAPConfigMappedQoSFlowsToAddConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.MappedQoS_FlowsToAdd))); err != nil {
				return err
			}
			for i := range ie.MappedQoS_FlowsToAdd {
				if err := ie.MappedQoS_FlowsToAdd[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. mappedQoS-FlowsToRelease: sequence-of
	{
		if ie.MappedQoS_FlowsToRelease != nil {
			seqOf := e.NewSequenceOfEncoder(sDAPConfigMappedQoSFlowsToReleaseConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.MappedQoS_FlowsToRelease))); err != nil {
				return err
			}
			for i := range ie.MappedQoS_FlowsToRelease {
				if err := ie.MappedQoS_FlowsToRelease[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SDAP_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDAPConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdu-Session: ref
	{
		if err := ie.Pdu_Session.Decode(d); err != nil {
			return err
		}
	}

	// 4. sdap-HeaderDL: enumerated
	{
		v1, err := d.DecodeEnumerated(sDAPConfigSdapHeaderDLConstraints)
		if err != nil {
			return err
		}
		ie.Sdap_HeaderDL = v1
	}

	// 5. sdap-HeaderUL: enumerated
	{
		v2, err := d.DecodeEnumerated(sDAPConfigSdapHeaderULConstraints)
		if err != nil {
			return err
		}
		ie.Sdap_HeaderUL = v2
	}

	// 6. defaultDRB: boolean
	{
		v3, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.DefaultDRB = v3
	}

	// 7. mappedQoS-FlowsToAdd: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sDAPConfigMappedQoSFlowsToAddConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MappedQoS_FlowsToAdd = make([]QFI, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MappedQoS_FlowsToAdd[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. mappedQoS-FlowsToRelease: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sDAPConfigMappedQoSFlowsToReleaseConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MappedQoS_FlowsToRelease = make([]QFI, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MappedQoS_FlowsToRelease[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
