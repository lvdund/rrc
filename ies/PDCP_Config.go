package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config struct {
	Drb                           *PDCP_Config_drb                          `lb:1,ub:16383,optional`
	MoreThanOneRLC                *PDCP_Config_moreThanOneRLC               `optional,ext`
	T_Reordering                  *PDCP_Config_t_Reordering                 `optional,ext`
	CipheringDisabled             *PDCP_Config_cipheringDisabled            `optional,ext-1`
	DiscardTimerExt_r16           *DiscardTimerExt_r16                      `optional,ext-2,setuprelease`
	MoreThanTwoRLC_DRB_r16        *PDCP_Config_moreThanTwoRLC_DRB_r16       `lb:3,ub:3,optional,ext-2`
	EthernetHeaderCompression_r16 *EthernetHeaderCompression_r16            `optional,ext-2,setuprelease`
	SurvivalTimeStateSupport_r17  *PDCP_Config_survivalTimeStateSupport_r17 `optional,ext-3`
	UplinkDataCompression_r17     *UplinkDataCompression_r17                `optional,ext-3,setuprelease`
	DiscardTimerExt2_r17          *DiscardTimerExt2_r17                     `optional,ext-3,setuprelease`
	InitialRX_DELIV_r17           *aper.BitString                           `lb:32,ub:32,optional,ext-3`
}

func (ie *PDCP_Config) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.CipheringDisabled != nil || ie.DiscardTimerExt_r16 != nil || ie.MoreThanTwoRLC_DRB_r16 != nil || ie.EthernetHeaderCompression_r16 != nil || ie.SurvivalTimeStateSupport_r17 != nil || ie.UplinkDataCompression_r17 != nil || ie.DiscardTimerExt2_r17 != nil || ie.InitialRX_DELIV_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Drb != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drb != nil {
		if err = ie.Drb.Encode(w); err != nil {
			return utils.WrapError("Encode Drb", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.CipheringDisabled != nil, ie.DiscardTimerExt_r16 != nil || ie.MoreThanTwoRLC_DRB_r16 != nil || ie.EthernetHeaderCompression_r16 != nil, ie.SurvivalTimeStateSupport_r17 != nil || ie.UplinkDataCompression_r17 != nil || ie.DiscardTimerExt2_r17 != nil || ie.InitialRX_DELIV_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCP_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CipheringDisabled != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CipheringDisabled optional
			if ie.CipheringDisabled != nil {
				if err = ie.CipheringDisabled.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CipheringDisabled", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.DiscardTimerExt_r16 != nil, ie.MoreThanTwoRLC_DRB_r16 != nil, ie.EthernetHeaderCompression_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DiscardTimerExt_r16 optional
			if ie.DiscardTimerExt_r16 != nil {
				tmp_DiscardTimerExt_r16 := utils.SetupRelease[*DiscardTimerExt_r16]{
					Setup: ie.DiscardTimerExt_r16,
				}
				if err = tmp_DiscardTimerExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscardTimerExt_r16", err)
				}
			}
			// encode MoreThanTwoRLC_DRB_r16 optional
			if ie.MoreThanTwoRLC_DRB_r16 != nil {
				if err = ie.MoreThanTwoRLC_DRB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MoreThanTwoRLC_DRB_r16", err)
				}
			}
			// encode EthernetHeaderCompression_r16 optional
			if ie.EthernetHeaderCompression_r16 != nil {
				tmp_EthernetHeaderCompression_r16 := utils.SetupRelease[*EthernetHeaderCompression_r16]{
					Setup: ie.EthernetHeaderCompression_r16,
				}
				if err = tmp_EthernetHeaderCompression_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EthernetHeaderCompression_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.SurvivalTimeStateSupport_r17 != nil, ie.UplinkDataCompression_r17 != nil, ie.DiscardTimerExt2_r17 != nil, ie.InitialRX_DELIV_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SurvivalTimeStateSupport_r17 optional
			if ie.SurvivalTimeStateSupport_r17 != nil {
				if err = ie.SurvivalTimeStateSupport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SurvivalTimeStateSupport_r17", err)
				}
			}
			// encode UplinkDataCompression_r17 optional
			if ie.UplinkDataCompression_r17 != nil {
				tmp_UplinkDataCompression_r17 := utils.SetupRelease[*UplinkDataCompression_r17]{
					Setup: ie.UplinkDataCompression_r17,
				}
				if err = tmp_UplinkDataCompression_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkDataCompression_r17", err)
				}
			}
			// encode DiscardTimerExt2_r17 optional
			if ie.DiscardTimerExt2_r17 != nil {
				tmp_DiscardTimerExt2_r17 := utils.SetupRelease[*DiscardTimerExt2_r17]{
					Setup: ie.DiscardTimerExt2_r17,
				}
				if err = tmp_DiscardTimerExt2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscardTimerExt2_r17", err)
				}
			}
			// encode InitialRX_DELIV_r17 optional
			if ie.InitialRX_DELIV_r17 != nil {
				if err = extWriter.WriteBitString(ie.InitialRX_DELIV_r17.Bytes, uint(ie.InitialRX_DELIV_r17.NumBits), &aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode InitialRX_DELIV_r17", err)
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

func (ie *PDCP_Config) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DrbPresent bool
	if DrbPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DrbPresent {
		ie.Drb = new(PDCP_Config_drb)
		if err = ie.Drb.Decode(r); err != nil {
			return utils.WrapError("Decode Drb", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			CipheringDisabledPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CipheringDisabled optional
			if CipheringDisabledPresent {
				ie.CipheringDisabled = new(PDCP_Config_cipheringDisabled)
				if err = ie.CipheringDisabled.Decode(extReader); err != nil {
					return utils.WrapError("Decode CipheringDisabled", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			DiscardTimerExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MoreThanTwoRLC_DRB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EthernetHeaderCompression_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DiscardTimerExt_r16 optional
			if DiscardTimerExt_r16Present {
				tmp_DiscardTimerExt_r16 := utils.SetupRelease[*DiscardTimerExt_r16]{}
				if err = tmp_DiscardTimerExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscardTimerExt_r16", err)
				}
				ie.DiscardTimerExt_r16 = tmp_DiscardTimerExt_r16.Setup
			}
			// decode MoreThanTwoRLC_DRB_r16 optional
			if MoreThanTwoRLC_DRB_r16Present {
				ie.MoreThanTwoRLC_DRB_r16 = new(PDCP_Config_moreThanTwoRLC_DRB_r16)
				if err = ie.MoreThanTwoRLC_DRB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MoreThanTwoRLC_DRB_r16", err)
				}
			}
			// decode EthernetHeaderCompression_r16 optional
			if EthernetHeaderCompression_r16Present {
				tmp_EthernetHeaderCompression_r16 := utils.SetupRelease[*EthernetHeaderCompression_r16]{}
				if err = tmp_EthernetHeaderCompression_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EthernetHeaderCompression_r16", err)
				}
				ie.EthernetHeaderCompression_r16 = tmp_EthernetHeaderCompression_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SurvivalTimeStateSupport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkDataCompression_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DiscardTimerExt2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InitialRX_DELIV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SurvivalTimeStateSupport_r17 optional
			if SurvivalTimeStateSupport_r17Present {
				ie.SurvivalTimeStateSupport_r17 = new(PDCP_Config_survivalTimeStateSupport_r17)
				if err = ie.SurvivalTimeStateSupport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SurvivalTimeStateSupport_r17", err)
				}
			}
			// decode UplinkDataCompression_r17 optional
			if UplinkDataCompression_r17Present {
				tmp_UplinkDataCompression_r17 := utils.SetupRelease[*UplinkDataCompression_r17]{}
				if err = tmp_UplinkDataCompression_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkDataCompression_r17", err)
				}
				ie.UplinkDataCompression_r17 = tmp_UplinkDataCompression_r17.Setup
			}
			// decode DiscardTimerExt2_r17 optional
			if DiscardTimerExt2_r17Present {
				tmp_DiscardTimerExt2_r17 := utils.SetupRelease[*DiscardTimerExt2_r17]{}
				if err = tmp_DiscardTimerExt2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscardTimerExt2_r17", err)
				}
				ie.DiscardTimerExt2_r17 = tmp_DiscardTimerExt2_r17.Setup
			}
			// decode InitialRX_DELIV_r17 optional
			if InitialRX_DELIV_r17Present {
				var tmp_bs_InitialRX_DELIV_r17 []byte
				var n_InitialRX_DELIV_r17 uint
				if tmp_bs_InitialRX_DELIV_r17, n_InitialRX_DELIV_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode InitialRX_DELIV_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_InitialRX_DELIV_r17,
					NumBits: uint64(n_InitialRX_DELIV_r17),
				}
				ie.InitialRX_DELIV_r17 = &tmp_bitstring
			}
		}
	}
	return nil
}
