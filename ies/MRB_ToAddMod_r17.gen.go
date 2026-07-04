// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-ToAddMod-r17 (line 13180).

var mRBToAddModR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionId-r17", Optional: true},
		{Name: "mrb-Identity-r17"},
		{Name: "mrb-IdentityNew-r17", Optional: true},
		{Name: "reestablishPDCP-r17", Optional: true},
		{Name: "recoverPDCP-r17", Optional: true},
		{Name: "pdcp-Config-r17", Optional: true},
	},
}

const (
	MRB_ToAddMod_r17_ReestablishPDCP_r17_True = 0
)

var mRBToAddModR17ReestablishPDCPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_ToAddMod_r17_ReestablishPDCP_r17_True},
}

const (
	MRB_ToAddMod_r17_RecoverPDCP_r17_True = 0
)

var mRBToAddModR17RecoverPDCPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_ToAddMod_r17_RecoverPDCP_r17_True},
}

type MRB_ToAddMod_r17 struct {
	Mbs_SessionId_r17   *TMGI_r17
	Mrb_Identity_r17    MRB_Identity_r17
	Mrb_IdentityNew_r17 *MRB_Identity_r17
	ReestablishPDCP_r17 *int64
	RecoverPDCP_r17     *int64
	Pdcp_Config_r17     *PDCP_Config
}

func (ie *MRB_ToAddMod_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBToAddModR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_SessionId_r17 != nil, ie.Mrb_IdentityNew_r17 != nil, ie.ReestablishPDCP_r17 != nil, ie.RecoverPDCP_r17 != nil, ie.Pdcp_Config_r17 != nil}); err != nil {
		return err
	}

	// 3. mbs-SessionId-r17: ref
	{
		if ie.Mbs_SessionId_r17 != nil {
			if err := ie.Mbs_SessionId_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mrb-Identity-r17: ref
	{
		if err := ie.Mrb_Identity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. mrb-IdentityNew-r17: ref
	{
		if ie.Mrb_IdentityNew_r17 != nil {
			if err := ie.Mrb_IdentityNew_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. reestablishPDCP-r17: enumerated
	{
		if ie.ReestablishPDCP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishPDCP_r17, mRBToAddModR17ReestablishPDCPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. recoverPDCP-r17: enumerated
	{
		if ie.RecoverPDCP_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RecoverPDCP_r17, mRBToAddModR17RecoverPDCPR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. pdcp-Config-r17: ref
	{
		if ie.Pdcp_Config_r17 != nil {
			if err := ie.Pdcp_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRB_ToAddMod_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBToAddModR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mbs-SessionId-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_SessionId_r17 = new(TMGI_r17)
			if err := ie.Mbs_SessionId_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mrb-Identity-r17: ref
	{
		if err := ie.Mrb_Identity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. mrb-IdentityNew-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrb_IdentityNew_r17 = new(MRB_Identity_r17)
			if err := ie.Mrb_IdentityNew_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. reestablishPDCP-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mRBToAddModR17ReestablishPDCPR17Constraints)
			if err != nil {
				return err
			}
			ie.ReestablishPDCP_r17 = &idx
		}
	}

	// 7. recoverPDCP-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mRBToAddModR17RecoverPDCPR17Constraints)
			if err != nil {
				return err
			}
			ie.RecoverPDCP_r17 = &idx
		}
	}

	// 8. pdcp-Config-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Pdcp_Config_r17 = new(PDCP_Config)
			if err := ie.Pdcp_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
