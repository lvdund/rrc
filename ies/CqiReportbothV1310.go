package ies

import "rrc/utils"

// CQI-ReportBoth-v1310 ::= SEQUENCE
type CqiReportbothV1310 struct {
	CsiImConfigtoreleaselistextR13 *CsiImConfigtoreleaselistextR13
	CsiImConfigtoaddmodlistextR13  *CsiImConfigtoaddmodlistextR13
}
