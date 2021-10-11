// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sports_objects.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on SportsObjects with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SportsObjects) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SportsObjectsValidationError is the validation error returned by
// SportsObjects.Validate if the designated constraints aren't met.
type SportsObjectsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SportsObjectsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SportsObjectsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SportsObjectsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SportsObjectsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SportsObjectsValidationError) ErrorName() string { return "SportsObjectsValidationError" }

// Error satisfies the builtin error interface
func (e SportsObjectsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSportsObjects.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SportsObjectsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SportsObjectsValidationError{}

// Validate checks the field values on SportsObjects_ListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SportsObjects_ListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SportsObjects_ListRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetObjectName() != "" {

	}

	if m.GetDepartmentalOrganizationId() != 0 {

		if m.GetDepartmentalOrganizationId() < 100000 {
			return SportsObjects_ListRequestValidationError{
				field:  "DepartmentalOrganizationId",
				reason: "value must be greater than or equal to 100000",
			}
		}

	}

	// no validation rules for DepartmentalOrganizationName

	if m.GetSportsAreaName() != "" {

	}

	if m.GetSportsAreaType() != "" {

	}

	// no validation rules for Availability

	if m.GetSportKind() != "" {

	}

	return nil
}

// SportsObjects_ListRequestValidationError is the validation error returned by
// SportsObjects_ListRequest.Validate if the designated constraints aren't met.
type SportsObjects_ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SportsObjects_ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SportsObjects_ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SportsObjects_ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SportsObjects_ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SportsObjects_ListRequestValidationError) ErrorName() string {
	return "SportsObjects_ListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SportsObjects_ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSportsObjects_ListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SportsObjects_ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SportsObjects_ListRequestValidationError{}

// Validate checks the field values on SportsObjects_ListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SportsObjects_ListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSportsObjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SportsObjects_ListResponseValidationError{
					field:  fmt.Sprintf("SportsObjects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetListStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SportsObjects_ListResponseValidationError{
				field:  "ListStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SportsObjects_ListResponseValidationError is the validation error returned
// by SportsObjects_ListResponse.Validate if the designated constraints aren't met.
type SportsObjects_ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SportsObjects_ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SportsObjects_ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SportsObjects_ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SportsObjects_ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SportsObjects_ListResponseValidationError) ErrorName() string {
	return "SportsObjects_ListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SportsObjects_ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSportsObjects_ListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SportsObjects_ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SportsObjects_ListResponseValidationError{}
