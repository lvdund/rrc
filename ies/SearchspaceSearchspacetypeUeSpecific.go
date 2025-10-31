package ies

// SearchSpace-searchSpaceType-ue-Specific ::= SEQUENCE
// Extensible
type SearchspaceSearchspacetypeUeSpecific struct {
	DciFormats       SearchspaceSearchspacetypeUeSpecificDciFormats
	DciFormatsMtR16  *SearchspaceSearchspacetypeUeSpecificDciFormatsMtR16
	DciFormatsslR16  *SearchspaceSearchspacetypeUeSpecificDciFormatsslR16
	DciFormatsextR16 *SearchspaceSearchspacetypeUeSpecificDciFormatsextR16
}
