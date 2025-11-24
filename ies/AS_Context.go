package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AS_Context struct {
	ReestablishmentInfo              *ReestablishmentInfo          `optional`
	ConfigRestrictInfo               *ConfigRestrictInfoSCG        `optional`
	Ran_NotificationAreaInfo         *RAN_NotificationAreaInfo     `optional,ext-1`
	UeAssistanceInformation          *[]byte                       `optional,ext-2`
	SelectedBandCombinationSN        *BandCombinationInfoSN        `optional,ext-3`
	ConfigRestrictInfoDAPS_r16       *ConfigRestrictInfoDAPS_r16   `optional,ext-4`
	SidelinkUEInformationNR_r16      *[]byte                       `optional,ext-4`
	SidelinkUEInformationEUTRA_r16   *[]byte                       `optional,ext-4`
	UeAssistanceInformationEUTRA_r16 *[]byte                       `optional,ext-4`
	UeAssistanceInformationSCG_r16   *[]byte                       `optional,ext-4`
	NeedForGapsInfoNR_r16            *NeedForGapsInfoNR_r16        `optional,ext-4`
	ConfigRestrictInfoDAPS_v1640     *ConfigRestrictInfoDAPS_v1640 `optional,ext-5`
	NeedForGapNCSG_InfoNR_r17        *NeedForGapNCSG_InfoNR_r17    `optional,ext-6`
	NeedForGapNCSG_InfoEUTRA_r17     *NeedForGapNCSG_InfoEUTRA_r17 `optional,ext-6`
	MbsInterestIndication_r17        *[]byte                       `optional,ext-6`
}

func (ie *AS_Context) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Ran_NotificationAreaInfo != nil || ie.UeAssistanceInformation != nil || ie.SelectedBandCombinationSN != nil || ie.ConfigRestrictInfoDAPS_r16 != nil || ie.SidelinkUEInformationNR_r16 != nil || ie.SidelinkUEInformationEUTRA_r16 != nil || ie.UeAssistanceInformationEUTRA_r16 != nil || ie.UeAssistanceInformationSCG_r16 != nil || ie.NeedForGapsInfoNR_r16 != nil || ie.ConfigRestrictInfoDAPS_v1640 != nil || ie.NeedForGapNCSG_InfoNR_r17 != nil || ie.NeedForGapNCSG_InfoEUTRA_r17 != nil || ie.MbsInterestIndication_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ReestablishmentInfo != nil, ie.ConfigRestrictInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReestablishmentInfo != nil {
		if err = ie.ReestablishmentInfo.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishmentInfo", err)
		}
	}
	if ie.ConfigRestrictInfo != nil {
		if err = ie.ConfigRestrictInfo.Encode(w); err != nil {
			return utils.WrapError("Encode ConfigRestrictInfo", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 6 bits for 6 extension groups
		extBitmap := []bool{ie.Ran_NotificationAreaInfo != nil, ie.UeAssistanceInformation != nil, ie.SelectedBandCombinationSN != nil, ie.ConfigRestrictInfoDAPS_r16 != nil || ie.SidelinkUEInformationNR_r16 != nil || ie.SidelinkUEInformationEUTRA_r16 != nil || ie.UeAssistanceInformationEUTRA_r16 != nil || ie.UeAssistanceInformationSCG_r16 != nil || ie.NeedForGapsInfoNR_r16 != nil, ie.ConfigRestrictInfoDAPS_v1640 != nil, ie.NeedForGapNCSG_InfoNR_r17 != nil || ie.NeedForGapNCSG_InfoEUTRA_r17 != nil || ie.MbsInterestIndication_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap AS_Context", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ran_NotificationAreaInfo != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ran_NotificationAreaInfo optional
			if ie.Ran_NotificationAreaInfo != nil {
				if err = ie.Ran_NotificationAreaInfo.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ran_NotificationAreaInfo", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.UeAssistanceInformation != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode UeAssistanceInformation optional
			if ie.UeAssistanceInformation != nil {
				if err = extWriter.WriteOctetString(*ie.UeAssistanceInformation, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode UeAssistanceInformation", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.SelectedBandCombinationSN != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SelectedBandCombinationSN optional
			if ie.SelectedBandCombinationSN != nil {
				if err = ie.SelectedBandCombinationSN.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SelectedBandCombinationSN", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.ConfigRestrictInfoDAPS_r16 != nil, ie.SidelinkUEInformationNR_r16 != nil, ie.SidelinkUEInformationEUTRA_r16 != nil, ie.UeAssistanceInformationEUTRA_r16 != nil, ie.UeAssistanceInformationSCG_r16 != nil, ie.NeedForGapsInfoNR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ConfigRestrictInfoDAPS_r16 optional
			if ie.ConfigRestrictInfoDAPS_r16 != nil {
				if err = ie.ConfigRestrictInfoDAPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfigRestrictInfoDAPS_r16", err)
				}
			}
			// encode SidelinkUEInformationNR_r16 optional
			if ie.SidelinkUEInformationNR_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.SidelinkUEInformationNR_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode SidelinkUEInformationNR_r16", err)
				}
			}
			// encode SidelinkUEInformationEUTRA_r16 optional
			if ie.SidelinkUEInformationEUTRA_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.SidelinkUEInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode SidelinkUEInformationEUTRA_r16", err)
				}
			}
			// encode UeAssistanceInformationEUTRA_r16 optional
			if ie.UeAssistanceInformationEUTRA_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.UeAssistanceInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode UeAssistanceInformationEUTRA_r16", err)
				}
			}
			// encode UeAssistanceInformationSCG_r16 optional
			if ie.UeAssistanceInformationSCG_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.UeAssistanceInformationSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode UeAssistanceInformationSCG_r16", err)
				}
			}
			// encode NeedForGapsInfoNR_r16 optional
			if ie.NeedForGapsInfoNR_r16 != nil {
				if err = ie.NeedForGapsInfoNR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NeedForGapsInfoNR_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.ConfigRestrictInfoDAPS_v1640 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ConfigRestrictInfoDAPS_v1640 optional
			if ie.ConfigRestrictInfoDAPS_v1640 != nil {
				if err = ie.ConfigRestrictInfoDAPS_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfigRestrictInfoDAPS_v1640", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.NeedForGapNCSG_InfoNR_r17 != nil, ie.NeedForGapNCSG_InfoEUTRA_r17 != nil, ie.MbsInterestIndication_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NeedForGapNCSG_InfoNR_r17 optional
			if ie.NeedForGapNCSG_InfoNR_r17 != nil {
				if err = ie.NeedForGapNCSG_InfoNR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NeedForGapNCSG_InfoNR_r17", err)
				}
			}
			// encode NeedForGapNCSG_InfoEUTRA_r17 optional
			if ie.NeedForGapNCSG_InfoEUTRA_r17 != nil {
				if err = ie.NeedForGapNCSG_InfoEUTRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NeedForGapNCSG_InfoEUTRA_r17", err)
				}
			}
			// encode MbsInterestIndication_r17 optional
			if ie.MbsInterestIndication_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.MbsInterestIndication_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode MbsInterestIndication_r17", err)
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

func (ie *AS_Context) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishmentInfoPresent bool
	if ReestablishmentInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfigRestrictInfoPresent bool
	if ConfigRestrictInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ReestablishmentInfoPresent {
		ie.ReestablishmentInfo = new(ReestablishmentInfo)
		if err = ie.ReestablishmentInfo.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishmentInfo", err)
		}
	}
	if ConfigRestrictInfoPresent {
		ie.ConfigRestrictInfo = new(ConfigRestrictInfoSCG)
		if err = ie.ConfigRestrictInfo.Decode(r); err != nil {
			return utils.WrapError("Decode ConfigRestrictInfo", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 6 bits for 6 extension groups
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

			Ran_NotificationAreaInfoPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ran_NotificationAreaInfo optional
			if Ran_NotificationAreaInfoPresent {
				ie.Ran_NotificationAreaInfo = new(RAN_NotificationAreaInfo)
				if err = ie.Ran_NotificationAreaInfo.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ran_NotificationAreaInfo", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			UeAssistanceInformationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode UeAssistanceInformation optional
			if UeAssistanceInformationPresent {
				var tmp_os_UeAssistanceInformation []byte
				if tmp_os_UeAssistanceInformation, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode UeAssistanceInformation", err)
				}
				ie.UeAssistanceInformation = &tmp_os_UeAssistanceInformation
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SelectedBandCombinationSNPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SelectedBandCombinationSN optional
			if SelectedBandCombinationSNPresent {
				ie.SelectedBandCombinationSN = new(BandCombinationInfoSN)
				if err = ie.SelectedBandCombinationSN.Decode(extReader); err != nil {
					return utils.WrapError("Decode SelectedBandCombinationSN", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ConfigRestrictInfoDAPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SidelinkUEInformationNR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SidelinkUEInformationEUTRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UeAssistanceInformationEUTRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UeAssistanceInformationSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NeedForGapsInfoNR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ConfigRestrictInfoDAPS_r16 optional
			if ConfigRestrictInfoDAPS_r16Present {
				ie.ConfigRestrictInfoDAPS_r16 = new(ConfigRestrictInfoDAPS_r16)
				if err = ie.ConfigRestrictInfoDAPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfigRestrictInfoDAPS_r16", err)
				}
			}
			// decode SidelinkUEInformationNR_r16 optional
			if SidelinkUEInformationNR_r16Present {
				var tmp_os_SidelinkUEInformationNR_r16 []byte
				if tmp_os_SidelinkUEInformationNR_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode SidelinkUEInformationNR_r16", err)
				}
				ie.SidelinkUEInformationNR_r16 = &tmp_os_SidelinkUEInformationNR_r16
			}
			// decode SidelinkUEInformationEUTRA_r16 optional
			if SidelinkUEInformationEUTRA_r16Present {
				var tmp_os_SidelinkUEInformationEUTRA_r16 []byte
				if tmp_os_SidelinkUEInformationEUTRA_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode SidelinkUEInformationEUTRA_r16", err)
				}
				ie.SidelinkUEInformationEUTRA_r16 = &tmp_os_SidelinkUEInformationEUTRA_r16
			}
			// decode UeAssistanceInformationEUTRA_r16 optional
			if UeAssistanceInformationEUTRA_r16Present {
				var tmp_os_UeAssistanceInformationEUTRA_r16 []byte
				if tmp_os_UeAssistanceInformationEUTRA_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode UeAssistanceInformationEUTRA_r16", err)
				}
				ie.UeAssistanceInformationEUTRA_r16 = &tmp_os_UeAssistanceInformationEUTRA_r16
			}
			// decode UeAssistanceInformationSCG_r16 optional
			if UeAssistanceInformationSCG_r16Present {
				var tmp_os_UeAssistanceInformationSCG_r16 []byte
				if tmp_os_UeAssistanceInformationSCG_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode UeAssistanceInformationSCG_r16", err)
				}
				ie.UeAssistanceInformationSCG_r16 = &tmp_os_UeAssistanceInformationSCG_r16
			}
			// decode NeedForGapsInfoNR_r16 optional
			if NeedForGapsInfoNR_r16Present {
				ie.NeedForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
				if err = ie.NeedForGapsInfoNR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode NeedForGapsInfoNR_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ConfigRestrictInfoDAPS_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ConfigRestrictInfoDAPS_v1640 optional
			if ConfigRestrictInfoDAPS_v1640Present {
				ie.ConfigRestrictInfoDAPS_v1640 = new(ConfigRestrictInfoDAPS_v1640)
				if err = ie.ConfigRestrictInfoDAPS_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfigRestrictInfoDAPS_v1640", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			NeedForGapNCSG_InfoNR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NeedForGapNCSG_InfoEUTRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MbsInterestIndication_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NeedForGapNCSG_InfoNR_r17 optional
			if NeedForGapNCSG_InfoNR_r17Present {
				ie.NeedForGapNCSG_InfoNR_r17 = new(NeedForGapNCSG_InfoNR_r17)
				if err = ie.NeedForGapNCSG_InfoNR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NeedForGapNCSG_InfoNR_r17", err)
				}
			}
			// decode NeedForGapNCSG_InfoEUTRA_r17 optional
			if NeedForGapNCSG_InfoEUTRA_r17Present {
				ie.NeedForGapNCSG_InfoEUTRA_r17 = new(NeedForGapNCSG_InfoEUTRA_r17)
				if err = ie.NeedForGapNCSG_InfoEUTRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NeedForGapNCSG_InfoEUTRA_r17", err)
				}
			}
			// decode MbsInterestIndication_r17 optional
			if MbsInterestIndication_r17Present {
				var tmp_os_MbsInterestIndication_r17 []byte
				if tmp_os_MbsInterestIndication_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode MbsInterestIndication_r17", err)
				}
				ie.MbsInterestIndication_r17 = &tmp_os_MbsInterestIndication_r17
			}
		}
	}
	return nil
}
