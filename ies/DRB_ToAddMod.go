package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRB_ToAddMod struct {
	CnAssociation   *DRB_ToAddMod_cnAssociation   `lb:0,ub:15,optional`
	Drb_Identity    DRB_Identity                  `madatory`
	ReestablishPDCP *DRB_ToAddMod_reestablishPDCP `optional`
	RecoverPDCP     *DRB_ToAddMod_recoverPDCP     `optional`
	Pdcp_Config     *PDCP_Config                  `optional`
	Daps_Config_r16 *DRB_ToAddMod_daps_Config_r16 `optional,ext-1`
}

func (ie *DRB_ToAddMod) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Daps_Config_r16 != nil
	preambleBits := []bool{hasExtensions, ie.CnAssociation != nil, ie.ReestablishPDCP != nil, ie.RecoverPDCP != nil, ie.Pdcp_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CnAssociation != nil {
		if err = ie.CnAssociation.Encode(w); err != nil {
			return utils.WrapError("Encode CnAssociation", err)
		}
	}
	if err = ie.Drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_Identity", err)
	}
	if ie.ReestablishPDCP != nil {
		if err = ie.ReestablishPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishPDCP", err)
		}
	}
	if ie.RecoverPDCP != nil {
		if err = ie.RecoverPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode RecoverPDCP", err)
		}
	}
	if ie.Pdcp_Config != nil {
		if err = ie.Pdcp_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_Config", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Daps_Config_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DRB_ToAddMod", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Daps_Config_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Daps_Config_r16 optional
			if ie.Daps_Config_r16 != nil {
				if err = ie.Daps_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Daps_Config_r16", err)
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

func (ie *DRB_ToAddMod) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CnAssociationPresent bool
	if CnAssociationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishPDCPPresent bool
	if ReestablishPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RecoverPDCPPresent bool
	if RecoverPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_ConfigPresent bool
	if Pdcp_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CnAssociationPresent {
		ie.CnAssociation = new(DRB_ToAddMod_cnAssociation)
		if err = ie.CnAssociation.Decode(r); err != nil {
			return utils.WrapError("Decode CnAssociation", err)
		}
	}
	if err = ie.Drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Drb_Identity", err)
	}
	if ReestablishPDCPPresent {
		ie.ReestablishPDCP = new(DRB_ToAddMod_reestablishPDCP)
		if err = ie.ReestablishPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishPDCP", err)
		}
	}
	if RecoverPDCPPresent {
		ie.RecoverPDCP = new(DRB_ToAddMod_recoverPDCP)
		if err = ie.RecoverPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode RecoverPDCP", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Daps_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Daps_Config_r16 optional
			if Daps_Config_r16Present {
				ie.Daps_Config_r16 = new(DRB_ToAddMod_daps_Config_r16)
				if err = ie.Daps_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Daps_Config_r16", err)
				}
			}
		}
	}
	return nil
}
