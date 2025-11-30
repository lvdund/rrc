package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb_headerCompression_rohc struct {
	MaxCID           int64                                                    `lb:1,ub:16383,madatory`
	Profiles         PDCP_Config_drb_headerCompression_rohc_profiles          `madatory`
	Drb_ContinueROHC *PDCP_Config_drb_headerCompression_rohc_drb_ContinueROHC `optional`
}

func (ie *PDCP_Config_drb_headerCompression_rohc) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drb_ContinueROHC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.MaxCID, &aper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger MaxCID", err)
	}
	if err = ie.Profiles.Encode(w); err != nil {
		return utils.WrapError("Encode Profiles", err)
	}
	if ie.Drb_ContinueROHC != nil {
		if err = ie.Drb_ContinueROHC.Encode(w); err != nil {
			return utils.WrapError("Encode Drb_ContinueROHC", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_drb_headerCompression_rohc) Decode(r *aper.AperReader) error {
	var err error
	var Drb_ContinueROHCPresent bool
	if Drb_ContinueROHCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_MaxCID int64
	if tmp_int_MaxCID, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger MaxCID", err)
	}
	ie.MaxCID = tmp_int_MaxCID
	if err = ie.Profiles.Decode(r); err != nil {
		return utils.WrapError("Decode Profiles", err)
	}
	if Drb_ContinueROHCPresent {
		ie.Drb_ContinueROHC = new(PDCP_Config_drb_headerCompression_rohc_drb_ContinueROHC)
		if err = ie.Drb_ContinueROHC.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_ContinueROHC", err)
		}
	}
	return nil
}
