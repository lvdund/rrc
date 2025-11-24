package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig_transformPrecodingDisabled struct {
	ScramblingID0   *int64                                                        `lb:0,ub:65535,optional`
	ScramblingID1   *int64                                                        `lb:0,ub:65535,optional`
	Dmrs_Uplink_r16 *DMRS_UplinkConfig_transformPrecodingDisabled_dmrs_Uplink_r16 `optional`
}

func (ie *DMRS_UplinkConfig_transformPrecodingDisabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ScramblingID0 != nil, ie.ScramblingID1 != nil, ie.Dmrs_Uplink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ScramblingID0 != nil {
		if err = w.WriteInteger(*ie.ScramblingID0, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode ScramblingID0", err)
		}
	}
	if ie.ScramblingID1 != nil {
		if err = w.WriteInteger(*ie.ScramblingID1, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode ScramblingID1", err)
		}
	}
	if ie.Dmrs_Uplink_r16 != nil {
		if err = ie.Dmrs_Uplink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_Uplink_r16", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig_transformPrecodingDisabled) Decode(r *uper.UperReader) error {
	var err error
	var ScramblingID0Present bool
	if ScramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScramblingID1Present bool
	if ScramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_Uplink_r16Present bool
	if Dmrs_Uplink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ScramblingID0Present {
		var tmp_int_ScramblingID0 int64
		if tmp_int_ScramblingID0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode ScramblingID0", err)
		}
		ie.ScramblingID0 = &tmp_int_ScramblingID0
	}
	if ScramblingID1Present {
		var tmp_int_ScramblingID1 int64
		if tmp_int_ScramblingID1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode ScramblingID1", err)
		}
		ie.ScramblingID1 = &tmp_int_ScramblingID1
	}
	if Dmrs_Uplink_r16Present {
		ie.Dmrs_Uplink_r16 = new(DMRS_UplinkConfig_transformPrecodingDisabled_dmrs_Uplink_r16)
		if err = ie.Dmrs_Uplink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_Uplink_r16", err)
		}
	}
	return nil
}
