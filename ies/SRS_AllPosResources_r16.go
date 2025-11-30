package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_AllPosResources_r16 struct {
	Srs_PosResources_r16  SRS_PosResources_r16   `madatory`
	Srs_PosResourceAP_r16 *SRS_PosResourceAP_r16 `optional`
	Srs_PosResourceSP_r16 *SRS_PosResourceSP_r16 `optional`
}

func (ie *SRS_AllPosResources_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_PosResourceAP_r16 != nil, ie.Srs_PosResourceSP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srs_PosResources_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_PosResources_r16", err)
	}
	if ie.Srs_PosResourceAP_r16 != nil {
		if err = ie.Srs_PosResourceAP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceAP_r16", err)
		}
	}
	if ie.Srs_PosResourceSP_r16 != nil {
		if err = ie.Srs_PosResourceSP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceSP_r16", err)
		}
	}
	return nil
}

func (ie *SRS_AllPosResources_r16) Decode(r *aper.AperReader) error {
	var err error
	var Srs_PosResourceAP_r16Present bool
	if Srs_PosResourceAP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_PosResourceSP_r16Present bool
	if Srs_PosResourceSP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srs_PosResources_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_PosResources_r16", err)
	}
	if Srs_PosResourceAP_r16Present {
		ie.Srs_PosResourceAP_r16 = new(SRS_PosResourceAP_r16)
		if err = ie.Srs_PosResourceAP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PosResourceAP_r16", err)
		}
	}
	if Srs_PosResourceSP_r16Present {
		ie.Srs_PosResourceSP_r16 = new(SRS_PosResourceSP_r16)
		if err = ie.Srs_PosResourceSP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PosResourceSP_r16", err)
		}
	}
	return nil
}
