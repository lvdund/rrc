package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_moreThanOneRLC struct {
	PrimaryPath           *PDCP_Config_moreThanOneRLC_primaryPath `optional`
	Ul_DataSplitThreshold *UL_DataSplitThreshold                  `optional`
	Pdcp_Duplication      *bool                                   `optional`
}

func (ie *PDCP_Config_moreThanOneRLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PrimaryPath != nil, ie.Ul_DataSplitThreshold != nil, ie.Pdcp_Duplication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PrimaryPath != nil {
		if err = ie.PrimaryPath.Encode(w); err != nil {
			return utils.WrapError("Encode PrimaryPath", err)
		}
	}
	if ie.Ul_DataSplitThreshold != nil {
		if err = ie.Ul_DataSplitThreshold.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_DataSplitThreshold", err)
		}
	}
	if ie.Pdcp_Duplication != nil {
		if err = w.WriteBoolean(*ie.Pdcp_Duplication); err != nil {
			return utils.WrapError("Encode Pdcp_Duplication", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_moreThanOneRLC) Decode(r *uper.UperReader) error {
	var err error
	var PrimaryPathPresent bool
	if PrimaryPathPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_DataSplitThresholdPresent bool
	if Ul_DataSplitThresholdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_DuplicationPresent bool
	if Pdcp_DuplicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PrimaryPathPresent {
		ie.PrimaryPath = new(PDCP_Config_moreThanOneRLC_primaryPath)
		if err = ie.PrimaryPath.Decode(r); err != nil {
			return utils.WrapError("Decode PrimaryPath", err)
		}
	}
	if Ul_DataSplitThresholdPresent {
		ie.Ul_DataSplitThreshold = new(UL_DataSplitThreshold)
		if err = ie.Ul_DataSplitThreshold.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_DataSplitThreshold", err)
		}
	}
	if Pdcp_DuplicationPresent {
		var tmp_bool_Pdcp_Duplication bool
		if tmp_bool_Pdcp_Duplication, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode Pdcp_Duplication", err)
		}
		ie.Pdcp_Duplication = &tmp_bool_Pdcp_Duplication
	}
	return nil
}
