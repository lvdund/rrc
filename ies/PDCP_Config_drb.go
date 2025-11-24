package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb struct {
	DiscardTimer         *PDCP_Config_drb_discardTimer         `optional`
	Pdcp_SN_SizeUL       *PDCP_Config_drb_pdcp_SN_SizeUL       `optional`
	Pdcp_SN_SizeDL       *PDCP_Config_drb_pdcp_SN_SizeDL       `optional`
	HeaderCompression    *PDCP_Config_drb_headerCompression    `lb:1,ub:16383,optional`
	IntegrityProtection  *PDCP_Config_drb_integrityProtection  `optional`
	StatusReportRequired *PDCP_Config_drb_statusReportRequired `optional`
	OutOfOrderDelivery   *PDCP_Config_drb_outOfOrderDelivery   `optional`
}

func (ie *PDCP_Config_drb) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.DiscardTimer != nil, ie.Pdcp_SN_SizeUL != nil, ie.Pdcp_SN_SizeDL != nil, ie.HeaderCompression != nil, ie.IntegrityProtection != nil, ie.StatusReportRequired != nil, ie.OutOfOrderDelivery != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DiscardTimer != nil {
		if err = ie.DiscardTimer.Encode(w); err != nil {
			return utils.WrapError("Encode DiscardTimer", err)
		}
	}
	if ie.Pdcp_SN_SizeUL != nil {
		if err = ie.Pdcp_SN_SizeUL.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_SN_SizeUL", err)
		}
	}
	if ie.Pdcp_SN_SizeDL != nil {
		if err = ie.Pdcp_SN_SizeDL.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_SN_SizeDL", err)
		}
	}
	if ie.HeaderCompression != nil {
		if err = ie.HeaderCompression.Encode(w); err != nil {
			return utils.WrapError("Encode HeaderCompression", err)
		}
	}
	if ie.IntegrityProtection != nil {
		if err = ie.IntegrityProtection.Encode(w); err != nil {
			return utils.WrapError("Encode IntegrityProtection", err)
		}
	}
	if ie.StatusReportRequired != nil {
		if err = ie.StatusReportRequired.Encode(w); err != nil {
			return utils.WrapError("Encode StatusReportRequired", err)
		}
	}
	if ie.OutOfOrderDelivery != nil {
		if err = ie.OutOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode OutOfOrderDelivery", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_drb) Decode(r *uper.UperReader) error {
	var err error
	var DiscardTimerPresent bool
	if DiscardTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_SN_SizeULPresent bool
	if Pdcp_SN_SizeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_SN_SizeDLPresent bool
	if Pdcp_SN_SizeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var HeaderCompressionPresent bool
	if HeaderCompressionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IntegrityProtectionPresent bool
	if IntegrityProtectionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var StatusReportRequiredPresent bool
	if StatusReportRequiredPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OutOfOrderDeliveryPresent bool
	if OutOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DiscardTimerPresent {
		ie.DiscardTimer = new(PDCP_Config_drb_discardTimer)
		if err = ie.DiscardTimer.Decode(r); err != nil {
			return utils.WrapError("Decode DiscardTimer", err)
		}
	}
	if Pdcp_SN_SizeULPresent {
		ie.Pdcp_SN_SizeUL = new(PDCP_Config_drb_pdcp_SN_SizeUL)
		if err = ie.Pdcp_SN_SizeUL.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_SN_SizeUL", err)
		}
	}
	if Pdcp_SN_SizeDLPresent {
		ie.Pdcp_SN_SizeDL = new(PDCP_Config_drb_pdcp_SN_SizeDL)
		if err = ie.Pdcp_SN_SizeDL.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_SN_SizeDL", err)
		}
	}
	if HeaderCompressionPresent {
		ie.HeaderCompression = new(PDCP_Config_drb_headerCompression)
		if err = ie.HeaderCompression.Decode(r); err != nil {
			return utils.WrapError("Decode HeaderCompression", err)
		}
	}
	if IntegrityProtectionPresent {
		ie.IntegrityProtection = new(PDCP_Config_drb_integrityProtection)
		if err = ie.IntegrityProtection.Decode(r); err != nil {
			return utils.WrapError("Decode IntegrityProtection", err)
		}
	}
	if StatusReportRequiredPresent {
		ie.StatusReportRequired = new(PDCP_Config_drb_statusReportRequired)
		if err = ie.StatusReportRequired.Decode(r); err != nil {
			return utils.WrapError("Decode StatusReportRequired", err)
		}
	}
	if OutOfOrderDeliveryPresent {
		ie.OutOfOrderDelivery = new(PDCP_Config_drb_outOfOrderDelivery)
		if err = ie.OutOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode OutOfOrderDelivery", err)
		}
	}
	return nil
}
