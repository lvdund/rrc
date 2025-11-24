package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoXCG_NR_r16 struct {
	Dl_FreqInfo_NR_r16 *FrequencyConfig_NR_r16 `optional`
	Ul_FreqInfo_NR_r16 *FrequencyConfig_NR_r16 `optional`
}

func (ie *ServCellInfoXCG_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_FreqInfo_NR_r16 != nil, ie.Ul_FreqInfo_NR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_FreqInfo_NR_r16 != nil {
		if err = ie.Dl_FreqInfo_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_FreqInfo_NR_r16", err)
		}
	}
	if ie.Ul_FreqInfo_NR_r16 != nil {
		if err = ie.Ul_FreqInfo_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FreqInfo_NR_r16", err)
		}
	}
	return nil
}

func (ie *ServCellInfoXCG_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	var Dl_FreqInfo_NR_r16Present bool
	if Dl_FreqInfo_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FreqInfo_NR_r16Present bool
	if Ul_FreqInfo_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_FreqInfo_NR_r16Present {
		ie.Dl_FreqInfo_NR_r16 = new(FrequencyConfig_NR_r16)
		if err = ie.Dl_FreqInfo_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_FreqInfo_NR_r16", err)
		}
	}
	if Ul_FreqInfo_NR_r16Present {
		ie.Ul_FreqInfo_NR_r16 = new(FrequencyConfig_NR_r16)
		if err = ie.Ul_FreqInfo_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FreqInfo_NR_r16", err)
		}
	}
	return nil
}
