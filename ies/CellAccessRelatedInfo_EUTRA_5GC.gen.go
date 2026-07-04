// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellAccessRelatedInfo-EUTRA-5GC (line 5524).

var cellAccessRelatedInfoEUTRA5GCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityList-eutra-5gc"},
		{Name: "trackingAreaCode-eutra-5gc"},
		{Name: "ranac-5gc", Optional: true},
		{Name: "cellIdentity-eutra-5gc"},
	},
}

type CellAccessRelatedInfo_EUTRA_5GC struct {
	Plmn_IdentityList_Eutra_5gc PLMN_IdentityList_EUTRA_5GC
	TrackingAreaCode_Eutra_5gc  TrackingAreaCode
	Ranac_5gc                   *RAN_AreaCode
	CellIdentity_Eutra_5gc      CellIdentity_EUTRA_5GC
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellAccessRelatedInfoEUTRA5GCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ranac_5gc != nil}); err != nil {
		return err
	}

	// 2. plmn-IdentityList-eutra-5gc: ref
	{
		if err := ie.Plmn_IdentityList_Eutra_5gc.Encode(e); err != nil {
			return err
		}
	}

	// 3. trackingAreaCode-eutra-5gc: ref
	{
		if err := ie.TrackingAreaCode_Eutra_5gc.Encode(e); err != nil {
			return err
		}
	}

	// 4. ranac-5gc: ref
	{
		if ie.Ranac_5gc != nil {
			if err := ie.Ranac_5gc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. cellIdentity-eutra-5gc: ref
	{
		if err := ie.CellIdentity_Eutra_5gc.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellAccessRelatedInfoEUTRA5GCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-IdentityList-eutra-5gc: ref
	{
		if err := ie.Plmn_IdentityList_Eutra_5gc.Decode(d); err != nil {
			return err
		}
	}

	// 3. trackingAreaCode-eutra-5gc: ref
	{
		if err := ie.TrackingAreaCode_Eutra_5gc.Decode(d); err != nil {
			return err
		}
	}

	// 4. ranac-5gc: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ranac_5gc = new(RAN_AreaCode)
			if err := ie.Ranac_5gc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. cellIdentity-eutra-5gc: ref
	{
		if err := ie.CellIdentity_Eutra_5gc.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
