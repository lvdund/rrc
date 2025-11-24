package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRB_ToAddMod struct {
	Srb_Identity       SRB_Identity                  `madatory`
	ReestablishPDCP    *SRB_ToAddMod_reestablishPDCP `optional`
	DiscardOnPDCP      *SRB_ToAddMod_discardOnPDCP   `optional`
	Pdcp_Config        *PDCP_Config                  `optional`
	Srb_Identity_v1700 *SRB_Identity_v1700           `optional,ext-1`
}

func (ie *SRB_ToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Srb_Identity_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.ReestablishPDCP != nil, ie.DiscardOnPDCP != nil, ie.Pdcp_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Srb_Identity", err)
	}
	if ie.ReestablishPDCP != nil {
		if err = ie.ReestablishPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishPDCP", err)
		}
	}
	if ie.DiscardOnPDCP != nil {
		if err = ie.DiscardOnPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode DiscardOnPDCP", err)
		}
	}
	if ie.Pdcp_Config != nil {
		if err = ie.Pdcp_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_Config", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Srb_Identity_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRB_ToAddMod", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Srb_Identity_v1700 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Srb_Identity_v1700 optional
			if ie.Srb_Identity_v1700 != nil {
				if err = ie.Srb_Identity_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srb_Identity_v1700", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SRB_ToAddMod) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishPDCPPresent bool
	if ReestablishPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DiscardOnPDCPPresent bool
	if DiscardOnPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_ConfigPresent bool
	if Pdcp_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Srb_Identity", err)
	}
	if ReestablishPDCPPresent {
		ie.ReestablishPDCP = new(SRB_ToAddMod_reestablishPDCP)
		if err = ie.ReestablishPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishPDCP", err)
		}
	}
	if DiscardOnPDCPPresent {
		ie.DiscardOnPDCP = new(SRB_ToAddMod_discardOnPDCP)
		if err = ie.DiscardOnPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode DiscardOnPDCP", err)
		}
	}
	if Pdcp_ConfigPresent {
		ie.Pdcp_Config = new(PDCP_Config)
		if err = ie.Pdcp_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_Config", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Srb_Identity_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Srb_Identity_v1700 optional
			if Srb_Identity_v1700Present {
				ie.Srb_Identity_v1700 = new(SRB_Identity_v1700)
				if err = ie.Srb_Identity_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srb_Identity_v1700", err)
				}
			}
		}
	}
	return nil
}
