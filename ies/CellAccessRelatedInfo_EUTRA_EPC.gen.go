// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellAccessRelatedInfo-EUTRA-EPC (line 5546).

var cellAccessRelatedInfoEUTRAEPCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityList-eutra-epc"},
		{Name: "trackingAreaCode-eutra-epc"},
		{Name: "cellIdentity-eutra-epc"},
	},
}

var cellAccessRelatedInfoEUTRAEPCTrackingAreaCodeEutraEpcConstraints = per.FixedSize(16)

var cellAccessRelatedInfoEUTRAEPCCellIdentityEutraEpcConstraints = per.FixedSize(28)

type CellAccessRelatedInfo_EUTRA_EPC struct {
	Plmn_IdentityList_Eutra_Epc PLMN_IdentityList_EUTRA_EPC
	TrackingAreaCode_Eutra_Epc  per.BitString
	CellIdentity_Eutra_Epc      per.BitString
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellAccessRelatedInfoEUTRAEPCConstraints)
	_ = seq

	// 1. plmn-IdentityList-eutra-epc: ref
	{
		if err := ie.Plmn_IdentityList_Eutra_Epc.Encode(e); err != nil {
			return err
		}
	}

	// 2. trackingAreaCode-eutra-epc: bit-string
	{
		if err := e.EncodeBitString(ie.TrackingAreaCode_Eutra_Epc, cellAccessRelatedInfoEUTRAEPCTrackingAreaCodeEutraEpcConstraints); err != nil {
			return err
		}
	}

	// 3. cellIdentity-eutra-epc: bit-string
	{
		if err := e.EncodeBitString(ie.CellIdentity_Eutra_Epc, cellAccessRelatedInfoEUTRAEPCCellIdentityEutraEpcConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellAccessRelatedInfoEUTRAEPCConstraints)
	_ = seq

	// 1. plmn-IdentityList-eutra-epc: ref
	{
		if err := ie.Plmn_IdentityList_Eutra_Epc.Decode(d); err != nil {
			return err
		}
	}

	// 2. trackingAreaCode-eutra-epc: bit-string
	{
		v1, err := d.DecodeBitString(cellAccessRelatedInfoEUTRAEPCTrackingAreaCodeEutraEpcConstraints)
		if err != nil {
			return err
		}
		ie.TrackingAreaCode_Eutra_Epc = v1
	}

	// 3. cellIdentity-eutra-epc: bit-string
	{
		v2, err := d.DecodeBitString(cellAccessRelatedInfoEUTRAEPCCellIdentityEutraEpcConstraints)
		if err != nil {
			return err
		}
		ie.CellIdentity_Eutra_Epc = v2
	}

	return nil
}
