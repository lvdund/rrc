package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters struct {
	SupportedROHC_Profiles             PDCP_Parameters_supportedROHC_Profiles              `madatory`
	MaxNumberROHC_ContextSessions      PDCP_Parameters_maxNumberROHC_ContextSessions       `madatory`
	UplinkOnlyROHC_Profiles            *PDCP_Parameters_uplinkOnlyROHC_Profiles            `optional`
	ContinueROHC_Context               *PDCP_Parameters_continueROHC_Context               `optional`
	OutOfOrderDelivery                 *PDCP_Parameters_outOfOrderDelivery                 `optional`
	ShortSN                            *PDCP_Parameters_shortSN                            `optional`
	Pdcp_DuplicationSRB                *PDCP_Parameters_pdcp_DuplicationSRB                `optional`
	Pdcp_DuplicationMCG_OrSCG_DRB      *PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB      `optional`
	Drb_IAB_r16                        *PDCP_Parameters_drb_IAB_r16                        `optional,ext-1`
	Non_DRB_IAB_r16                    *PDCP_Parameters_non_DRB_IAB_r16                    `optional,ext-1`
	ExtendedDiscardTimer_r16           *PDCP_Parameters_extendedDiscardTimer_r16           `optional,ext-1`
	ContinueEHC_Context_r16            *PDCP_Parameters_continueEHC_Context_r16            `optional,ext-1`
	Ehc_r16                            *PDCP_Parameters_ehc_r16                            `optional,ext-1`
	MaxNumberEHC_Contexts_r16          *PDCP_Parameters_maxNumberEHC_Contexts_r16          `optional,ext-1`
	JointEHC_ROHC_Config_r16           *PDCP_Parameters_jointEHC_ROHC_Config_r16           `optional,ext-1`
	Pdcp_DuplicationMoreThanTwoRLC_r16 *PDCP_Parameters_pdcp_DuplicationMoreThanTwoRLC_r16 `optional,ext-1`
	LongSN_RedCap_r17                  *PDCP_Parameters_longSN_RedCap_r17                  `optional,ext-2`
	Udc_r17                            *PDCP_Parameters_udc_r17                            `lb:0,ub:15,optional,ext-2`
}

func (ie *PDCP_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Drb_IAB_r16 != nil || ie.Non_DRB_IAB_r16 != nil || ie.ExtendedDiscardTimer_r16 != nil || ie.ContinueEHC_Context_r16 != nil || ie.Ehc_r16 != nil || ie.MaxNumberEHC_Contexts_r16 != nil || ie.JointEHC_ROHC_Config_r16 != nil || ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil || ie.LongSN_RedCap_r17 != nil || ie.Udc_r17 != nil
	preambleBits := []bool{hasExtensions, ie.UplinkOnlyROHC_Profiles != nil, ie.ContinueROHC_Context != nil, ie.OutOfOrderDelivery != nil, ie.ShortSN != nil, ie.Pdcp_DuplicationSRB != nil, ie.Pdcp_DuplicationMCG_OrSCG_DRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedROHC_Profiles.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedROHC_Profiles", err)
	}
	if err = ie.MaxNumberROHC_ContextSessions.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberROHC_ContextSessions", err)
	}
	if ie.UplinkOnlyROHC_Profiles != nil {
		if err = ie.UplinkOnlyROHC_Profiles.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkOnlyROHC_Profiles", err)
		}
	}
	if ie.ContinueROHC_Context != nil {
		if err = ie.ContinueROHC_Context.Encode(w); err != nil {
			return utils.WrapError("Encode ContinueROHC_Context", err)
		}
	}
	if ie.OutOfOrderDelivery != nil {
		if err = ie.OutOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode OutOfOrderDelivery", err)
		}
	}
	if ie.ShortSN != nil {
		if err = ie.ShortSN.Encode(w); err != nil {
			return utils.WrapError("Encode ShortSN", err)
		}
	}
	if ie.Pdcp_DuplicationSRB != nil {
		if err = ie.Pdcp_DuplicationSRB.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_DuplicationSRB", err)
		}
	}
	if ie.Pdcp_DuplicationMCG_OrSCG_DRB != nil {
		if err = ie.Pdcp_DuplicationMCG_OrSCG_DRB.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_DuplicationMCG_OrSCG_DRB", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Drb_IAB_r16 != nil || ie.Non_DRB_IAB_r16 != nil || ie.ExtendedDiscardTimer_r16 != nil || ie.ContinueEHC_Context_r16 != nil || ie.Ehc_r16 != nil || ie.MaxNumberEHC_Contexts_r16 != nil || ie.JointEHC_ROHC_Config_r16 != nil || ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil, ie.LongSN_RedCap_r17 != nil || ie.Udc_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCP_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Drb_IAB_r16 != nil, ie.Non_DRB_IAB_r16 != nil, ie.ExtendedDiscardTimer_r16 != nil, ie.ContinueEHC_Context_r16 != nil, ie.Ehc_r16 != nil, ie.MaxNumberEHC_Contexts_r16 != nil, ie.JointEHC_ROHC_Config_r16 != nil, ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Drb_IAB_r16 optional
			if ie.Drb_IAB_r16 != nil {
				if err = ie.Drb_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drb_IAB_r16", err)
				}
			}
			// encode Non_DRB_IAB_r16 optional
			if ie.Non_DRB_IAB_r16 != nil {
				if err = ie.Non_DRB_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Non_DRB_IAB_r16", err)
				}
			}
			// encode ExtendedDiscardTimer_r16 optional
			if ie.ExtendedDiscardTimer_r16 != nil {
				if err = ie.ExtendedDiscardTimer_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedDiscardTimer_r16", err)
				}
			}
			// encode ContinueEHC_Context_r16 optional
			if ie.ContinueEHC_Context_r16 != nil {
				if err = ie.ContinueEHC_Context_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ContinueEHC_Context_r16", err)
				}
			}
			// encode Ehc_r16 optional
			if ie.Ehc_r16 != nil {
				if err = ie.Ehc_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ehc_r16", err)
				}
			}
			// encode MaxNumberEHC_Contexts_r16 optional
			if ie.MaxNumberEHC_Contexts_r16 != nil {
				if err = ie.MaxNumberEHC_Contexts_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberEHC_Contexts_r16", err)
				}
			}
			// encode JointEHC_ROHC_Config_r16 optional
			if ie.JointEHC_ROHC_Config_r16 != nil {
				if err = ie.JointEHC_ROHC_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode JointEHC_ROHC_Config_r16", err)
				}
			}
			// encode Pdcp_DuplicationMoreThanTwoRLC_r16 optional
			if ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil {
				if err = ie.Pdcp_DuplicationMoreThanTwoRLC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcp_DuplicationMoreThanTwoRLC_r16", err)
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
			optionals_ext_2 := []bool{ie.LongSN_RedCap_r17 != nil, ie.Udc_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode LongSN_RedCap_r17 optional
			if ie.LongSN_RedCap_r17 != nil {
				if err = ie.LongSN_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LongSN_RedCap_r17", err)
				}
			}
			// encode Udc_r17 optional
			if ie.Udc_r17 != nil {
				if err = ie.Udc_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Udc_r17", err)
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

func (ie *PDCP_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkOnlyROHC_ProfilesPresent bool
	if UplinkOnlyROHC_ProfilesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ContinueROHC_ContextPresent bool
	if ContinueROHC_ContextPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OutOfOrderDeliveryPresent bool
	if OutOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ShortSNPresent bool
	if ShortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_DuplicationSRBPresent bool
	if Pdcp_DuplicationSRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_DuplicationMCG_OrSCG_DRBPresent bool
	if Pdcp_DuplicationMCG_OrSCG_DRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedROHC_Profiles.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedROHC_Profiles", err)
	}
	if err = ie.MaxNumberROHC_ContextSessions.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberROHC_ContextSessions", err)
	}
	if UplinkOnlyROHC_ProfilesPresent {
		ie.UplinkOnlyROHC_Profiles = new(PDCP_Parameters_uplinkOnlyROHC_Profiles)
		if err = ie.UplinkOnlyROHC_Profiles.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkOnlyROHC_Profiles", err)
		}
	}
	if ContinueROHC_ContextPresent {
		ie.ContinueROHC_Context = new(PDCP_Parameters_continueROHC_Context)
		if err = ie.ContinueROHC_Context.Decode(r); err != nil {
			return utils.WrapError("Decode ContinueROHC_Context", err)
		}
	}
	if OutOfOrderDeliveryPresent {
		ie.OutOfOrderDelivery = new(PDCP_Parameters_outOfOrderDelivery)
		if err = ie.OutOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode OutOfOrderDelivery", err)
		}
	}
	if ShortSNPresent {
		ie.ShortSN = new(PDCP_Parameters_shortSN)
		if err = ie.ShortSN.Decode(r); err != nil {
			return utils.WrapError("Decode ShortSN", err)
		}
	}
	if Pdcp_DuplicationSRBPresent {
		ie.Pdcp_DuplicationSRB = new(PDCP_Parameters_pdcp_DuplicationSRB)
		if err = ie.Pdcp_DuplicationSRB.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_DuplicationSRB", err)
		}
	}
	if Pdcp_DuplicationMCG_OrSCG_DRBPresent {
		ie.Pdcp_DuplicationMCG_OrSCG_DRB = new(PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB)
		if err = ie.Pdcp_DuplicationMCG_OrSCG_DRB.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_DuplicationMCG_OrSCG_DRB", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			Drb_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Non_DRB_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedDiscardTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ContinueEHC_Context_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ehc_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberEHC_Contexts_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			JointEHC_ROHC_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcp_DuplicationMoreThanTwoRLC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Drb_IAB_r16 optional
			if Drb_IAB_r16Present {
				ie.Drb_IAB_r16 = new(PDCP_Parameters_drb_IAB_r16)
				if err = ie.Drb_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drb_IAB_r16", err)
				}
			}
			// decode Non_DRB_IAB_r16 optional
			if Non_DRB_IAB_r16Present {
				ie.Non_DRB_IAB_r16 = new(PDCP_Parameters_non_DRB_IAB_r16)
				if err = ie.Non_DRB_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Non_DRB_IAB_r16", err)
				}
			}
			// decode ExtendedDiscardTimer_r16 optional
			if ExtendedDiscardTimer_r16Present {
				ie.ExtendedDiscardTimer_r16 = new(PDCP_Parameters_extendedDiscardTimer_r16)
				if err = ie.ExtendedDiscardTimer_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedDiscardTimer_r16", err)
				}
			}
			// decode ContinueEHC_Context_r16 optional
			if ContinueEHC_Context_r16Present {
				ie.ContinueEHC_Context_r16 = new(PDCP_Parameters_continueEHC_Context_r16)
				if err = ie.ContinueEHC_Context_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ContinueEHC_Context_r16", err)
				}
			}
			// decode Ehc_r16 optional
			if Ehc_r16Present {
				ie.Ehc_r16 = new(PDCP_Parameters_ehc_r16)
				if err = ie.Ehc_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ehc_r16", err)
				}
			}
			// decode MaxNumberEHC_Contexts_r16 optional
			if MaxNumberEHC_Contexts_r16Present {
				ie.MaxNumberEHC_Contexts_r16 = new(PDCP_Parameters_maxNumberEHC_Contexts_r16)
				if err = ie.MaxNumberEHC_Contexts_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberEHC_Contexts_r16", err)
				}
			}
			// decode JointEHC_ROHC_Config_r16 optional
			if JointEHC_ROHC_Config_r16Present {
				ie.JointEHC_ROHC_Config_r16 = new(PDCP_Parameters_jointEHC_ROHC_Config_r16)
				if err = ie.JointEHC_ROHC_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode JointEHC_ROHC_Config_r16", err)
				}
			}
			// decode Pdcp_DuplicationMoreThanTwoRLC_r16 optional
			if Pdcp_DuplicationMoreThanTwoRLC_r16Present {
				ie.Pdcp_DuplicationMoreThanTwoRLC_r16 = new(PDCP_Parameters_pdcp_DuplicationMoreThanTwoRLC_r16)
				if err = ie.Pdcp_DuplicationMoreThanTwoRLC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcp_DuplicationMoreThanTwoRLC_r16", err)
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

			LongSN_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Udc_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode LongSN_RedCap_r17 optional
			if LongSN_RedCap_r17Present {
				ie.LongSN_RedCap_r17 = new(PDCP_Parameters_longSN_RedCap_r17)
				if err = ie.LongSN_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode LongSN_RedCap_r17", err)
				}
			}
			// decode Udc_r17 optional
			if Udc_r17Present {
				ie.Udc_r17 = new(PDCP_Parameters_udc_r17)
				if err = ie.Udc_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Udc_r17", err)
				}
			}
		}
	}
	return nil
}
