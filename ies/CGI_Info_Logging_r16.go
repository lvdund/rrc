package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_Info_Logging_r16 struct {
	Plmn_Identity_r16    PLMN_Identity     `madatory`
	CellIdentity_r16     CellIdentity      `madatory`
	TrackingAreaCode_r16 *TrackingAreaCode `optional`
}

func (ie *CGI_Info_Logging_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TrackingAreaCode_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r16", err)
	}
	if err = ie.CellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity_r16", err)
	}
	if ie.TrackingAreaCode_r16 != nil {
		if err = ie.TrackingAreaCode_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TrackingAreaCode_r16", err)
		}
	}
	return nil
}

func (ie *CGI_Info_Logging_r16) Decode(r *uper.UperReader) error {
	var err error
	var TrackingAreaCode_r16Present bool
	if TrackingAreaCode_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r16", err)
	}
	if err = ie.CellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity_r16", err)
	}
	if TrackingAreaCode_r16Present {
		ie.TrackingAreaCode_r16 = new(TrackingAreaCode)
		if err = ie.TrackingAreaCode_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TrackingAreaCode_r16", err)
		}
	}
	return nil
}
