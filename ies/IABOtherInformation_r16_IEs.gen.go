// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IABOtherInformation-r16-IEs (line 436).

var iABOtherInformationR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ip-InfoType-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var iABOtherInformation_r16_IEsIpInfoTypeR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iab-IP-Request-r16"},
		{Name: "iab-IP-Report-r16"},
	},
}

const (
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16 = 0
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16  = 1
)

var iABOtherInformationR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

var iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-IPv4-AddressNumReq-r16", Optional: true},
		{Name: "iab-IPv6-AddressReq-r16", Optional: true},
	},
}

var iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16IabIPv6AddressReqR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iab-IPv6-AddressNumReq-r16"},
		{Name: "iab-IPv6-AddressPrefixReq-r16"},
	},
}

const (
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressNumReq_r16    = 0
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressPrefixReq_r16 = 1
)

var iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-IPv4-AddressReport-r16", Optional: true},
		{Name: "iab-IPv6-Report-r16", Optional: true},
	},
}

var iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16IabIPv6ReportR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iab-IPv6-AddressReport-r16"},
		{Name: "iab-IPv6-PrefixReport-r16"},
	},
}

const (
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_AddressReport_r16 = 0
	IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_PrefixReport_r16  = 1
)

type IABOtherInformation_r16_IEs struct {
	Ip_InfoType_r16 struct {
		Choice             int
		Iab_IP_Request_r16 *struct {
			Iab_IPv4_AddressNumReq_r16 *IAB_IP_AddressNumReq_r16
			Iab_IPv6_AddressReq_r16    *struct {
				Choice                        int
				Iab_IPv6_AddressNumReq_r16    *IAB_IP_AddressNumReq_r16
				Iab_IPv6_AddressPrefixReq_r16 *IAB_IP_AddressPrefixReq_r16
			}
		}
		Iab_IP_Report_r16 *struct {
			Iab_IPv4_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
			Iab_IPv6_Report_r16        *struct {
				Choice                     int
				Iab_IPv6_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
				Iab_IPv6_PrefixReport_r16  *IAB_IP_PrefixAndTraffic_r16
			}
		}
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *IABOtherInformation_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABOtherInformationR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ip-InfoType-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(iABOtherInformation_r16_IEsIpInfoTypeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Ip_InfoType_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Ip_InfoType_r16.Choice {
		case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16:
			c := (*ie.Ip_InfoType_r16.Iab_IP_Request_r16)
			iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq := e.NewSequenceEncoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Constraints)
			if err := iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq.EncodePreamble([]bool{c.Iab_IPv4_AddressNumReq_r16 != nil, c.Iab_IPv6_AddressReq_r16 != nil}); err != nil {
				return err
			}
			if c.Iab_IPv4_AddressNumReq_r16 != nil {
				if err := c.Iab_IPv4_AddressNumReq_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Iab_IPv6_AddressReq_r16 != nil {
				choiceEnc := e.NewChoiceEncoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16IabIPv6AddressReqR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Iab_IPv6_AddressReq_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Iab_IPv6_AddressReq_r16).Choice {
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressNumReq_r16:
					if err := (*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressNumReq_r16.Encode(e); err != nil {
						return err
					}
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressPrefixReq_r16:
					if err := (*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressPrefixReq_r16.Encode(e); err != nil {
						return err
					}
				}
			}
		case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16:
			c := (*ie.Ip_InfoType_r16.Iab_IP_Report_r16)
			iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq := e.NewSequenceEncoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Constraints)
			if err := iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq.EncodePreamble([]bool{c.Iab_IPv4_AddressReport_r16 != nil, c.Iab_IPv6_Report_r16 != nil}); err != nil {
				return err
			}
			if c.Iab_IPv4_AddressReport_r16 != nil {
				if err := c.Iab_IPv4_AddressReport_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Iab_IPv6_Report_r16 != nil {
				choiceEnc := e.NewChoiceEncoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16IabIPv6ReportR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Iab_IPv6_Report_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Iab_IPv6_Report_r16).Choice {
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_AddressReport_r16:
					if err := (*c.Iab_IPv6_Report_r16).Iab_IPv6_AddressReport_r16.Encode(e); err != nil {
						return err
					}
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_PrefixReport_r16:
					if err := (*c.Iab_IPv6_Report_r16).Iab_IPv6_PrefixReport_r16.Encode(e); err != nil {
						return err
					}
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Ip_InfoType_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, iABOtherInformationR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *IABOtherInformation_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABOtherInformationR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ip-InfoType-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(iABOtherInformation_r16_IEsIpInfoTypeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Ip_InfoType_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16:
			ie.Ip_InfoType_r16.Iab_IP_Request_r16 = &struct {
				Iab_IPv4_AddressNumReq_r16 *IAB_IP_AddressNumReq_r16
				Iab_IPv6_AddressReq_r16    *struct {
					Choice                        int
					Iab_IPv6_AddressNumReq_r16    *IAB_IP_AddressNumReq_r16
					Iab_IPv6_AddressPrefixReq_r16 *IAB_IP_AddressPrefixReq_r16
				}
			}{}
			c := (*ie.Ip_InfoType_r16.Iab_IP_Request_r16)
			iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq := d.NewSequenceDecoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Constraints)
			if err := iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq.IsComponentPresent(0) {
				c.Iab_IPv4_AddressNumReq_r16 = new(IAB_IP_AddressNumReq_r16)
				if err := (*c.Iab_IPv4_AddressNumReq_r16).Decode(d); err != nil {
					return err
				}
			}
			if iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16Seq.IsComponentPresent(1) {
				c.Iab_IPv6_AddressReq_r16 = &struct {
					Choice                        int
					Iab_IPv6_AddressNumReq_r16    *IAB_IP_AddressNumReq_r16
					Iab_IPv6_AddressPrefixReq_r16 *IAB_IP_AddressPrefixReq_r16
				}{}
				choiceDec := d.NewChoiceDecoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPRequestR16IabIPv6AddressReqR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Iab_IPv6_AddressReq_r16).Choice = int(index)
				switch index {
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressNumReq_r16:
					(*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressNumReq_r16 = new(IAB_IP_AddressNumReq_r16)
					if err := (*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressNumReq_r16.Decode(d); err != nil {
						return err
					}
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Request_r16_Iab_IPv6_AddressReq_r16_Iab_IPv6_AddressPrefixReq_r16:
					(*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressPrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16)
					if err := (*c.Iab_IPv6_AddressReq_r16).Iab_IPv6_AddressPrefixReq_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16:
			ie.Ip_InfoType_r16.Iab_IP_Report_r16 = &struct {
				Iab_IPv4_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
				Iab_IPv6_Report_r16        *struct {
					Choice                     int
					Iab_IPv6_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
					Iab_IPv6_PrefixReport_r16  *IAB_IP_PrefixAndTraffic_r16
				}
			}{}
			c := (*ie.Ip_InfoType_r16.Iab_IP_Report_r16)
			iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq := d.NewSequenceDecoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Constraints)
			if err := iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq.IsComponentPresent(0) {
				c.Iab_IPv4_AddressReport_r16 = new(IAB_IP_AddressAndTraffic_r16)
				if err := (*c.Iab_IPv4_AddressReport_r16).Decode(d); err != nil {
					return err
				}
			}
			if iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16Seq.IsComponentPresent(1) {
				c.Iab_IPv6_Report_r16 = &struct {
					Choice                     int
					Iab_IPv6_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
					Iab_IPv6_PrefixReport_r16  *IAB_IP_PrefixAndTraffic_r16
				}{}
				choiceDec := d.NewChoiceDecoder(iABOtherInformationR16IEsIpInfoTypeR16IabIPReportR16IabIPv6ReportR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Iab_IPv6_Report_r16).Choice = int(index)
				switch index {
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_AddressReport_r16:
					(*c.Iab_IPv6_Report_r16).Iab_IPv6_AddressReport_r16 = new(IAB_IP_AddressAndTraffic_r16)
					if err := (*c.Iab_IPv6_Report_r16).Iab_IPv6_AddressReport_r16.Decode(d); err != nil {
						return err
					}
				case IABOtherInformation_r16_IEs_Ip_InfoType_r16_Iab_IP_Report_r16_Iab_IPv6_Report_r16_Iab_IPv6_PrefixReport_r16:
					(*c.Iab_IPv6_Report_r16).Iab_IPv6_PrefixReport_r16 = new(IAB_IP_PrefixAndTraffic_r16)
					if err := (*c.Iab_IPv6_Report_r16).Iab_IPv6_PrefixReport_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(iABOtherInformationR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
