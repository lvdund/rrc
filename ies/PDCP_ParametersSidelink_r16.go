package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_ParametersSidelink_r16 struct {
	OutOfOrderDeliverySidelink_r16 *PDCP_ParametersSidelink_r16_outOfOrderDeliverySidelink_r16 `optional`
}

func (ie *PDCP_ParametersSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OutOfOrderDeliverySidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OutOfOrderDeliverySidelink_r16 != nil {
		if err = ie.OutOfOrderDeliverySidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OutOfOrderDeliverySidelink_r16", err)
		}
	}
	return nil
}

func (ie *PDCP_ParametersSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var OutOfOrderDeliverySidelink_r16Present bool
	if OutOfOrderDeliverySidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OutOfOrderDeliverySidelink_r16Present {
		ie.OutOfOrderDeliverySidelink_r16 = new(PDCP_ParametersSidelink_r16_outOfOrderDeliverySidelink_r16)
		if err = ie.OutOfOrderDeliverySidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OutOfOrderDeliverySidelink_r16", err)
		}
	}
	return nil
}
