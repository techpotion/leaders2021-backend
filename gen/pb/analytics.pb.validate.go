// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: analytics.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on PolygonAnalytics with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PolygonAnalytics) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolygonAnalytics with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PolygonAnalyticsMultiError, or nil if none found.
func (m *PolygonAnalytics) ValidateAll() error {
	return m.validate(true)
}

func (m *PolygonAnalytics) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PolygonAnalyticsMultiError(errors)
	}
	return nil
}

// PolygonAnalyticsMultiError is an error wrapping multiple validation errors
// returned by PolygonAnalytics.ValidateAll() if the designated constraints
// aren't met.
type PolygonAnalyticsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolygonAnalyticsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolygonAnalyticsMultiError) AllErrors() []error { return m }

// PolygonAnalyticsValidationError is the validation error returned by
// PolygonAnalytics.Validate if the designated constraints aren't met.
type PolygonAnalyticsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalyticsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalyticsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalyticsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalyticsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalyticsValidationError) ErrorName() string { return "PolygonAnalyticsValidationError" }

// Error satisfies the builtin error interface
func (e PolygonAnalyticsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalytics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalyticsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalyticsValidationError{}

// Validate checks the field values on PolygonAnalytics_GetRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PolygonAnalytics_GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolygonAnalytics_GetRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PolygonAnalytics_GetRequestMultiError, or nil if none found.
func (m *PolygonAnalytics_GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PolygonAnalytics_GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPolygon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolygonAnalytics_GetRequestValidationError{
					field:  "Polygon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolygonAnalytics_GetRequestValidationError{
					field:  "Polygon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalytics_GetRequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSportsAreaType() != "" {

	}

	// no validation rules for Availability

	if m.GetSportKind() != "" {

	}

	if len(errors) > 0 {
		return PolygonAnalytics_GetRequestMultiError(errors)
	}
	return nil
}

// PolygonAnalytics_GetRequestMultiError is an error wrapping multiple
// validation errors returned by PolygonAnalytics_GetRequest.ValidateAll() if
// the designated constraints aren't met.
type PolygonAnalytics_GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolygonAnalytics_GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolygonAnalytics_GetRequestMultiError) AllErrors() []error { return m }

// PolygonAnalytics_GetRequestValidationError is the validation error returned
// by PolygonAnalytics_GetRequest.Validate if the designated constraints
// aren't met.
type PolygonAnalytics_GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalytics_GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalytics_GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalytics_GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalytics_GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalytics_GetRequestValidationError) ErrorName() string {
	return "PolygonAnalytics_GetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalytics_GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalytics_GetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalytics_GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalytics_GetRequestValidationError{}

// Validate checks the field values on PolygonAnalytics_GetResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PolygonAnalytics_GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolygonAnalytics_GetResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PolygonAnalytics_GetResponseMultiError, or nil if none found.
func (m *PolygonAnalytics_GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PolygonAnalytics_GetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AreasSquare

	// no validation rules for AreasAmount

	// no validation rules for SportsAmount

	if len(errors) > 0 {
		return PolygonAnalytics_GetResponseMultiError(errors)
	}
	return nil
}

// PolygonAnalytics_GetResponseMultiError is an error wrapping multiple
// validation errors returned by PolygonAnalytics_GetResponse.ValidateAll() if
// the designated constraints aren't met.
type PolygonAnalytics_GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolygonAnalytics_GetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolygonAnalytics_GetResponseMultiError) AllErrors() []error { return m }

// PolygonAnalytics_GetResponseValidationError is the validation error returned
// by PolygonAnalytics_GetResponse.Validate if the designated constraints
// aren't met.
type PolygonAnalytics_GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalytics_GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalytics_GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalytics_GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalytics_GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalytics_GetResponseValidationError) ErrorName() string {
	return "PolygonAnalytics_GetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalytics_GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalytics_GetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalytics_GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalytics_GetResponseValidationError{}
