package mapping

import "errors"

var (
	// ErrEntryTypeNotSupported entry type has no appropriate mapper type
	ErrEntryTypeNotSupported = errors.New(`entry type not supported for mapping`)

	// ErrEntryTypeNotDefined
	ErrStateMappingNotFound = errors.New(`state mapping not found`)

	// ErrEventMappingNotFound
	ErrEventMappingNotFound = errors.New(`event mapping not found`)
)
