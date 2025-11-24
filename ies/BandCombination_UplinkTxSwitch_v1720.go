package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1720 struct {
	BandCombination_v1720                   *BandCombination_v1720                                                        `optional`
	UplinkTxSwitching_OptionSupport2T2T_r17 *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v1720 != nil, ie.UplinkTxSwitching_OptionSupport2T2T_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v1720 != nil {
		if err = ie.BandCombination_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1720", err)
		}
	}
	if ie.UplinkTxSwitching_OptionSupport2T2T_r17 != nil {
		if err = ie.UplinkTxSwitching_OptionSupport2T2T_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxSwitching_OptionSupport2T2T_r17", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Decode(r *uper.UperReader) error {
	var err error
	var BandCombination_v1720Present bool
	if BandCombination_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkTxSwitching_OptionSupport2T2T_r17Present bool
	if UplinkTxSwitching_OptionSupport2T2T_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v1720Present {
		ie.BandCombination_v1720 = new(BandCombination_v1720)
		if err = ie.BandCombination_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1720", err)
		}
	}
	if UplinkTxSwitching_OptionSupport2T2T_r17Present {
		ie.UplinkTxSwitching_OptionSupport2T2T_r17 = new(BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17)
		if err = ie.UplinkTxSwitching_OptionSupport2T2T_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxSwitching_OptionSupport2T2T_r17", err)
		}
	}
	return nil
}
