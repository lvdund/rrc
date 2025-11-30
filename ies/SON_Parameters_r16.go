package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SON_Parameters_r16 struct {
	Rach_Report_r16        *SON_Parameters_r16_rach_Report_r16        `optional`
	RlfReportCHO_r17       *SON_Parameters_r16_rlfReportCHO_r17       `optional,ext-1`
	RlfReportDAPS_r17      *SON_Parameters_r16_rlfReportDAPS_r17      `optional,ext-1`
	Success_HO_Report_r17  *SON_Parameters_r16_success_HO_Report_r17  `optional,ext-1`
	TwoStepRACH_Report_r17 *SON_Parameters_r16_twoStepRACH_Report_r17 `optional,ext-1`
	Pscell_MHI_Report_r17  *SON_Parameters_r16_pscell_MHI_Report_r17  `optional,ext-1`
	OnDemandSI_Report_r17  *SON_Parameters_r16_onDemandSI_Report_r17  `optional,ext-1`
}

func (ie *SON_Parameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.RlfReportCHO_r17 != nil || ie.RlfReportDAPS_r17 != nil || ie.Success_HO_Report_r17 != nil || ie.TwoStepRACH_Report_r17 != nil || ie.Pscell_MHI_Report_r17 != nil || ie.OnDemandSI_Report_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Rach_Report_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rach_Report_r16 != nil {
		if err = ie.Rach_Report_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_Report_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.RlfReportCHO_r17 != nil || ie.RlfReportDAPS_r17 != nil || ie.Success_HO_Report_r17 != nil || ie.TwoStepRACH_Report_r17 != nil || ie.Pscell_MHI_Report_r17 != nil || ie.OnDemandSI_Report_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SON_Parameters_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RlfReportCHO_r17 != nil, ie.RlfReportDAPS_r17 != nil, ie.Success_HO_Report_r17 != nil, ie.TwoStepRACH_Report_r17 != nil, ie.Pscell_MHI_Report_r17 != nil, ie.OnDemandSI_Report_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RlfReportCHO_r17 optional
			if ie.RlfReportCHO_r17 != nil {
				if err = ie.RlfReportCHO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RlfReportCHO_r17", err)
				}
			}
			// encode RlfReportDAPS_r17 optional
			if ie.RlfReportDAPS_r17 != nil {
				if err = ie.RlfReportDAPS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RlfReportDAPS_r17", err)
				}
			}
			// encode Success_HO_Report_r17 optional
			if ie.Success_HO_Report_r17 != nil {
				if err = ie.Success_HO_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Success_HO_Report_r17", err)
				}
			}
			// encode TwoStepRACH_Report_r17 optional
			if ie.TwoStepRACH_Report_r17 != nil {
				if err = ie.TwoStepRACH_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoStepRACH_Report_r17", err)
				}
			}
			// encode Pscell_MHI_Report_r17 optional
			if ie.Pscell_MHI_Report_r17 != nil {
				if err = ie.Pscell_MHI_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pscell_MHI_Report_r17", err)
				}
			}
			// encode OnDemandSI_Report_r17 optional
			if ie.OnDemandSI_Report_r17 != nil {
				if err = ie.OnDemandSI_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OnDemandSI_Report_r17", err)
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

func (ie *SON_Parameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rach_Report_r16Present bool
	if Rach_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rach_Report_r16Present {
		ie.Rach_Report_r16 = new(SON_Parameters_r16_rach_Report_r16)
		if err = ie.Rach_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_Report_r16", err)
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

			RlfReportCHO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RlfReportDAPS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Success_HO_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TwoStepRACH_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pscell_MHI_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OnDemandSI_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RlfReportCHO_r17 optional
			if RlfReportCHO_r17Present {
				ie.RlfReportCHO_r17 = new(SON_Parameters_r16_rlfReportCHO_r17)
				if err = ie.RlfReportCHO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode RlfReportCHO_r17", err)
				}
			}
			// decode RlfReportDAPS_r17 optional
			if RlfReportDAPS_r17Present {
				ie.RlfReportDAPS_r17 = new(SON_Parameters_r16_rlfReportDAPS_r17)
				if err = ie.RlfReportDAPS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode RlfReportDAPS_r17", err)
				}
			}
			// decode Success_HO_Report_r17 optional
			if Success_HO_Report_r17Present {
				ie.Success_HO_Report_r17 = new(SON_Parameters_r16_success_HO_Report_r17)
				if err = ie.Success_HO_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Success_HO_Report_r17", err)
				}
			}
			// decode TwoStepRACH_Report_r17 optional
			if TwoStepRACH_Report_r17Present {
				ie.TwoStepRACH_Report_r17 = new(SON_Parameters_r16_twoStepRACH_Report_r17)
				if err = ie.TwoStepRACH_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoStepRACH_Report_r17", err)
				}
			}
			// decode Pscell_MHI_Report_r17 optional
			if Pscell_MHI_Report_r17Present {
				ie.Pscell_MHI_Report_r17 = new(SON_Parameters_r16_pscell_MHI_Report_r17)
				if err = ie.Pscell_MHI_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pscell_MHI_Report_r17", err)
				}
			}
			// decode OnDemandSI_Report_r17 optional
			if OnDemandSI_Report_r17Present {
				ie.OnDemandSI_Report_r17 = new(SON_Parameters_r16_onDemandSI_Report_r17)
				if err = ie.OnDemandSI_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode OnDemandSI_Report_r17", err)
				}
			}
		}
	}
	return nil
}
