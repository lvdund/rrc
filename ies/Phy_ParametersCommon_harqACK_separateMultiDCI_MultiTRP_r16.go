package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16 struct {
	MaxNumberLongPUCCHs_r16 *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16 `optional`
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberLongPUCCHs_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumberLongPUCCHs_r16 != nil {
		if err = ie.MaxNumberLongPUCCHs_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberLongPUCCHs_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16) Decode(r *aper.AperReader) error {
	var err error
	var MaxNumberLongPUCCHs_r16Present bool
	if MaxNumberLongPUCCHs_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumberLongPUCCHs_r16Present {
		ie.MaxNumberLongPUCCHs_r16 = new(Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16)
		if err = ie.MaxNumberLongPUCCHs_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberLongPUCCHs_r16", err)
		}
	}
	return nil
}
