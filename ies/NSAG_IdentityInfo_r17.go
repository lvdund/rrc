package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NSAG_IdentityInfo_r17 struct {
	Nsag_ID_r17          NSAG_ID_r17       `madatory`
	TrackingAreaCode_r17 *TrackingAreaCode `optional`
}

func (ie *NSAG_IdentityInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TrackingAreaCode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Nsag_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nsag_ID_r17", err)
	}
	if ie.TrackingAreaCode_r17 != nil {
		if err = ie.TrackingAreaCode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TrackingAreaCode_r17", err)
		}
	}
	return nil
}

func (ie *NSAG_IdentityInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var TrackingAreaCode_r17Present bool
	if TrackingAreaCode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Nsag_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Nsag_ID_r17", err)
	}
	if TrackingAreaCode_r17Present {
		ie.TrackingAreaCode_r17 = new(TrackingAreaCode)
		if err = ie.TrackingAreaCode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TrackingAreaCode_r17", err)
		}
	}
	return nil
}
