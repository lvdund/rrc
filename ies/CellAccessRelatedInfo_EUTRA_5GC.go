package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo_EUTRA_5GC struct {
	Plmn_IdentityList_eutra_5gc PLMN_IdentityList_EUTRA_5GC `madatory`
	TrackingAreaCode_eutra_5gc  TrackingAreaCode            `madatory`
	Ranac_5gc                   *RAN_AreaCode               `optional`
	CellIdentity_eutra_5gc      CellIdentity_EUTRA_5GC      `madatory`
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ranac_5gc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Plmn_IdentityList_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_eutra_5gc", err)
	}
	if err = ie.TrackingAreaCode_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaCode_eutra_5gc", err)
	}
	if ie.Ranac_5gc != nil {
		if err = ie.Ranac_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode Ranac_5gc", err)
		}
	}
	if err = ie.CellIdentity_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity_eutra_5gc", err)
	}
	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Decode(r *aper.AperReader) error {
	var err error
	var Ranac_5gcPresent bool
	if Ranac_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Plmn_IdentityList_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_eutra_5gc", err)
	}
	if err = ie.TrackingAreaCode_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode TrackingAreaCode_eutra_5gc", err)
	}
	if Ranac_5gcPresent {
		ie.Ranac_5gc = new(RAN_AreaCode)
		if err = ie.Ranac_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode Ranac_5gc", err)
		}
	}
	if err = ie.CellIdentity_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity_eutra_5gc", err)
	}
	return nil
}
