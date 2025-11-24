package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkCommon struct {
	GenericParameters  BWP                 `madatory`
	Pdcch_ConfigCommon *PDCCH_ConfigCommon `optional,setuprelease`
	Pdsch_ConfigCommon *PDSCH_ConfigCommon `optional,setuprelease`
}

func (ie *BWP_DownlinkCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcch_ConfigCommon != nil, ie.Pdsch_ConfigCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.GenericParameters.Encode(w); err != nil {
		return utils.WrapError("Encode GenericParameters", err)
	}
	if ie.Pdcch_ConfigCommon != nil {
		tmp_Pdcch_ConfigCommon := utils.SetupRelease[*PDCCH_ConfigCommon]{
			Setup: ie.Pdcch_ConfigCommon,
		}
		if err = tmp_Pdcch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_ConfigCommon", err)
		}
	}
	if ie.Pdsch_ConfigCommon != nil {
		tmp_Pdsch_ConfigCommon := utils.SetupRelease[*PDSCH_ConfigCommon]{
			Setup: ie.Pdsch_ConfigCommon,
		}
		if err = tmp_Pdsch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ConfigCommon", err)
		}
	}
	return nil
}

func (ie *BWP_DownlinkCommon) Decode(r *uper.UperReader) error {
	var err error
	var Pdcch_ConfigCommonPresent bool
	if Pdcch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigCommonPresent bool
	if Pdsch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.GenericParameters.Decode(r); err != nil {
		return utils.WrapError("Decode GenericParameters", err)
	}
	if Pdcch_ConfigCommonPresent {
		tmp_Pdcch_ConfigCommon := utils.SetupRelease[*PDCCH_ConfigCommon]{}
		if err = tmp_Pdcch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_ConfigCommon", err)
		}
		ie.Pdcch_ConfigCommon = tmp_Pdcch_ConfigCommon.Setup
	}
	if Pdsch_ConfigCommonPresent {
		tmp_Pdsch_ConfigCommon := utils.SetupRelease[*PDSCH_ConfigCommon]{}
		if err = tmp_Pdsch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ConfigCommon", err)
		}
		ie.Pdsch_ConfigCommon = tmp_Pdsch_ConfigCommon.Setup
	}
	return nil
}
