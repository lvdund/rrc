// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RRCResumeComplete-IEs (line 1597).

var rRCResumeCompleteIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dedicatedNAS-Message", Optional: true},
		{Name: "selectedPLMN-Identity", Optional: true},
		{Name: "uplinkTxDirectCurrentList", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCResumeCompleteIEsSelectedPLMNIdentityConstraints = per.Constrained(1, common.MaxPLMN)

var rRCResumeCompleteIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCResumeComplete_IEs struct {
	DedicatedNAS_Message      *DedicatedNAS_Message
	SelectedPLMN_Identity     *int64
	UplinkTxDirectCurrentList *UplinkTxDirectCurrentList
	LateNonCriticalExtension  []byte
	NonCriticalExtension      *RRCResumeComplete_v1610_IEs
}

func (ie *RRCResumeComplete_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DedicatedNAS_Message != nil, ie.SelectedPLMN_Identity != nil, ie.UplinkTxDirectCurrentList != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. dedicatedNAS-Message: ref
	{
		if ie.DedicatedNAS_Message != nil {
			if err := ie.DedicatedNAS_Message.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. selectedPLMN-Identity: integer
	{
		if ie.SelectedPLMN_Identity != nil {
			if err := e.EncodeInteger(*ie.SelectedPLMN_Identity, rRCResumeCompleteIEsSelectedPLMNIdentityConstraints); err != nil {
				return err
			}
		}
	}

	// 4. uplinkTxDirectCurrentList: ref
	{
		if ie.UplinkTxDirectCurrentList != nil {
			if err := ie.UplinkTxDirectCurrentList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCResumeCompleteIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *RRCResumeComplete_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dedicatedNAS-Message: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DedicatedNAS_Message = new(DedicatedNAS_Message)
			if err := ie.DedicatedNAS_Message.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. selectedPLMN-Identity: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rRCResumeCompleteIEsSelectedPLMNIdentityConstraints)
			if err != nil {
				return err
			}
			ie.SelectedPLMN_Identity = &v
		}
	}

	// 4. uplinkTxDirectCurrentList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.UplinkTxDirectCurrentList = new(UplinkTxDirectCurrentList)
			if err := ie.UplinkTxDirectCurrentList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(rRCResumeCompleteIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(RRCResumeComplete_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
