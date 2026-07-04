// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CGI-InfoEUTRALogging (line 5926).

var cGIInfoEUTRALoggingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-eutra-5gc", Optional: true},
		{Name: "trackingAreaCode-eutra-5gc", Optional: true},
		{Name: "cellIdentity-eutra-5gc", Optional: true},
		{Name: "plmn-Identity-eutra-epc", Optional: true},
		{Name: "trackingAreaCode-eutra-epc", Optional: true},
		{Name: "cellIdentity-eutra-epc", Optional: true},
	},
}

var cGIInfoEUTRALoggingCellIdentityEutra5gcConstraints = per.FixedSize(28)

var cGIInfoEUTRALoggingTrackingAreaCodeEutraEpcConstraints = per.FixedSize(16)

var cGIInfoEUTRALoggingCellIdentityEutraEpcConstraints = per.FixedSize(28)

type CGI_InfoEUTRALogging struct {
	Plmn_Identity_Eutra_5gc    *PLMN_Identity
	TrackingAreaCode_Eutra_5gc *TrackingAreaCode
	CellIdentity_Eutra_5gc     *per.BitString
	Plmn_Identity_Eutra_Epc    *PLMN_Identity
	TrackingAreaCode_Eutra_Epc *per.BitString
	CellIdentity_Eutra_Epc     *per.BitString
}

func (ie *CGI_InfoEUTRALogging) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGIInfoEUTRALoggingConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Plmn_Identity_Eutra_5gc != nil, ie.TrackingAreaCode_Eutra_5gc != nil, ie.CellIdentity_Eutra_5gc != nil, ie.Plmn_Identity_Eutra_Epc != nil, ie.TrackingAreaCode_Eutra_Epc != nil, ie.CellIdentity_Eutra_Epc != nil}); err != nil {
		return err
	}

	// 2. plmn-Identity-eutra-5gc: ref
	{
		if ie.Plmn_Identity_Eutra_5gc != nil {
			if err := ie.Plmn_Identity_Eutra_5gc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. trackingAreaCode-eutra-5gc: ref
	{
		if ie.TrackingAreaCode_Eutra_5gc != nil {
			if err := ie.TrackingAreaCode_Eutra_5gc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. cellIdentity-eutra-5gc: bit-string
	{
		if ie.CellIdentity_Eutra_5gc != nil {
			if err := e.EncodeBitString(*ie.CellIdentity_Eutra_5gc, cGIInfoEUTRALoggingCellIdentityEutra5gcConstraints); err != nil {
				return err
			}
		}
	}

	// 5. plmn-Identity-eutra-epc: ref
	{
		if ie.Plmn_Identity_Eutra_Epc != nil {
			if err := ie.Plmn_Identity_Eutra_Epc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. trackingAreaCode-eutra-epc: bit-string
	{
		if ie.TrackingAreaCode_Eutra_Epc != nil {
			if err := e.EncodeBitString(*ie.TrackingAreaCode_Eutra_Epc, cGIInfoEUTRALoggingTrackingAreaCodeEutraEpcConstraints); err != nil {
				return err
			}
		}
	}

	// 7. cellIdentity-eutra-epc: bit-string
	{
		if ie.CellIdentity_Eutra_Epc != nil {
			if err := e.EncodeBitString(*ie.CellIdentity_Eutra_Epc, cGIInfoEUTRALoggingCellIdentityEutraEpcConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CGI_InfoEUTRALogging) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGIInfoEUTRALoggingConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-Identity-eutra-5gc: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Plmn_Identity_Eutra_5gc = new(PLMN_Identity)
			if err := ie.Plmn_Identity_Eutra_5gc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. trackingAreaCode-eutra-5gc: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TrackingAreaCode_Eutra_5gc = new(TrackingAreaCode)
			if err := ie.TrackingAreaCode_Eutra_5gc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. cellIdentity-eutra-5gc: bit-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBitString(cGIInfoEUTRALoggingCellIdentityEutra5gcConstraints)
			if err != nil {
				return err
			}
			ie.CellIdentity_Eutra_5gc = &v
		}
	}

	// 5. plmn-Identity-eutra-epc: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Plmn_Identity_Eutra_Epc = new(PLMN_Identity)
			if err := ie.Plmn_Identity_Eutra_Epc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. trackingAreaCode-eutra-epc: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(cGIInfoEUTRALoggingTrackingAreaCodeEutraEpcConstraints)
			if err != nil {
				return err
			}
			ie.TrackingAreaCode_Eutra_Epc = &v
		}
	}

	// 7. cellIdentity-eutra-epc: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(cGIInfoEUTRALoggingCellIdentityEutraEpcConstraints)
			if err != nil {
				return err
			}
			ie.CellIdentity_Eutra_Epc = &v
		}
	}

	return nil
}
