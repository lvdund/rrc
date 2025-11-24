package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16 struct {
	BundleSizeSet1_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16 `optional`
	BundleSizeSet2_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet2_r16 `optional`
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BundleSizeSet1_r16 != nil, ie.BundleSizeSet2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BundleSizeSet1_r16 != nil {
		if err = ie.BundleSizeSet1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSizeSet1_r16", err)
		}
	}
	if ie.BundleSizeSet2_r16 != nil {
		if err = ie.BundleSizeSet2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSizeSet2_r16", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16) Decode(r *uper.UperReader) error {
	var err error
	var BundleSizeSet1_r16Present bool
	if BundleSizeSet1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BundleSizeSet2_r16Present bool
	if BundleSizeSet2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BundleSizeSet1_r16Present {
		ie.BundleSizeSet1_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16)
		if err = ie.BundleSizeSet1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSizeSet1_r16", err)
		}
	}
	if BundleSizeSet2_r16Present {
		ie.BundleSizeSet2_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet2_r16)
		if err = ie.BundleSizeSet2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSizeSet2_r16", err)
		}
	}
	return nil
}
