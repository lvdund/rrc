package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters_udc_r17 struct {
	StandardDictionary_r17  *PDCP_Parameters_udc_r17_standardDictionary_r17  `optional`
	OperatorDictionary_r17  *PDCP_Parameters_udc_r17_operatorDictionary_r17  `lb:0,ub:15,optional`
	ContinueUDC_r17         *PDCP_Parameters_udc_r17_continueUDC_r17         `optional`
	SupportOfBufferSize_r17 *PDCP_Parameters_udc_r17_supportOfBufferSize_r17 `optional`
}

func (ie *PDCP_Parameters_udc_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.StandardDictionary_r17 != nil, ie.OperatorDictionary_r17 != nil, ie.ContinueUDC_r17 != nil, ie.SupportOfBufferSize_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.StandardDictionary_r17 != nil {
		if err = ie.StandardDictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode StandardDictionary_r17", err)
		}
	}
	if ie.OperatorDictionary_r17 != nil {
		if err = ie.OperatorDictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OperatorDictionary_r17", err)
		}
	}
	if ie.ContinueUDC_r17 != nil {
		if err = ie.ContinueUDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ContinueUDC_r17", err)
		}
	}
	if ie.SupportOfBufferSize_r17 != nil {
		if err = ie.SupportOfBufferSize_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportOfBufferSize_r17", err)
		}
	}
	return nil
}

func (ie *PDCP_Parameters_udc_r17) Decode(r *aper.AperReader) error {
	var err error
	var StandardDictionary_r17Present bool
	if StandardDictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OperatorDictionary_r17Present bool
	if OperatorDictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ContinueUDC_r17Present bool
	if ContinueUDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportOfBufferSize_r17Present bool
	if SupportOfBufferSize_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if StandardDictionary_r17Present {
		ie.StandardDictionary_r17 = new(PDCP_Parameters_udc_r17_standardDictionary_r17)
		if err = ie.StandardDictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode StandardDictionary_r17", err)
		}
	}
	if OperatorDictionary_r17Present {
		ie.OperatorDictionary_r17 = new(PDCP_Parameters_udc_r17_operatorDictionary_r17)
		if err = ie.OperatorDictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OperatorDictionary_r17", err)
		}
	}
	if ContinueUDC_r17Present {
		ie.ContinueUDC_r17 = new(PDCP_Parameters_udc_r17_continueUDC_r17)
		if err = ie.ContinueUDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ContinueUDC_r17", err)
		}
	}
	if SupportOfBufferSize_r17Present {
		ie.SupportOfBufferSize_r17 = new(PDCP_Parameters_udc_r17_supportOfBufferSize_r17)
		if err = ie.SupportOfBufferSize_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportOfBufferSize_r17", err)
		}
	}
	return nil
}
