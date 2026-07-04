// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CGI-Info-Logging-r16 (line 5960).

var cGIInfoLoggingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-r16"},
		{Name: "cellIdentity-r16"},
		{Name: "trackingAreaCode-r16", Optional: true},
	},
}

type CGI_Info_Logging_r16 struct {
	Plmn_Identity_r16    PLMN_Identity
	CellIdentity_r16     CellIdentity
	TrackingAreaCode_r16 *TrackingAreaCode
}

func (ie *CGI_Info_Logging_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGIInfoLoggingR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TrackingAreaCode_r16 != nil}); err != nil {
		return err
	}

	// 2. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. cellIdentity-r16: ref
	{
		if err := ie.CellIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. trackingAreaCode-r16: ref
	{
		if ie.TrackingAreaCode_r16 != nil {
			if err := ie.TrackingAreaCode_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CGI_Info_Logging_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGIInfoLoggingR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. cellIdentity-r16: ref
	{
		if err := ie.CellIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. trackingAreaCode-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.TrackingAreaCode_r16 = new(TrackingAreaCode)
			if err := ie.TrackingAreaCode_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
