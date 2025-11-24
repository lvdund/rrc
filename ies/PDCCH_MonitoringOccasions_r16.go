package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_MonitoringOccasions_r16 struct {
	Period7span3_r16 *PDCCH_MonitoringOccasions_r16_period7span3_r16 `optional`
	Period4span3_r16 *PDCCH_MonitoringOccasions_r16_period4span3_r16 `optional`
	Period2span2_r16 *PDCCH_MonitoringOccasions_r16_period2span2_r16 `optional`
}

func (ie *PDCCH_MonitoringOccasions_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Period7span3_r16 != nil, ie.Period4span3_r16 != nil, ie.Period2span2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Period7span3_r16 != nil {
		if err = ie.Period7span3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Period7span3_r16", err)
		}
	}
	if ie.Period4span3_r16 != nil {
		if err = ie.Period4span3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Period4span3_r16", err)
		}
	}
	if ie.Period2span2_r16 != nil {
		if err = ie.Period2span2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Period2span2_r16", err)
		}
	}
	return nil
}

func (ie *PDCCH_MonitoringOccasions_r16) Decode(r *uper.UperReader) error {
	var err error
	var Period7span3_r16Present bool
	if Period7span3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Period4span3_r16Present bool
	if Period4span3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Period2span2_r16Present bool
	if Period2span2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Period7span3_r16Present {
		ie.Period7span3_r16 = new(PDCCH_MonitoringOccasions_r16_period7span3_r16)
		if err = ie.Period7span3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Period7span3_r16", err)
		}
	}
	if Period4span3_r16Present {
		ie.Period4span3_r16 = new(PDCCH_MonitoringOccasions_r16_period4span3_r16)
		if err = ie.Period4span3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Period4span3_r16", err)
		}
	}
	if Period2span2_r16Present {
		ie.Period2span2_r16 = new(PDCCH_MonitoringOccasions_r16_period2span2_r16)
		if err = ie.Period2span2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Period2span2_r16", err)
		}
	}
	return nil
}
