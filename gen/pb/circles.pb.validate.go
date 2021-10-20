// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: circles.proto

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

// Validate checks the field values on Circles with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Circles) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Circles with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CirclesMultiError, or nil if none found.
func (m *Circles) ValidateAll() error {
	return m.validate(true)
}

func (m *Circles) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CirclesMultiError(errors)
	}
	return nil
}

// CirclesMultiError is an error wrapping multiple validation errors returned
// by Circles.ValidateAll() if the designated constraints aren't met.
type CirclesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CirclesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CirclesMultiError) AllErrors() []error { return m }

// CirclesValidationError is the validation error returned by Circles.Validate
// if the designated constraints aren't met.
type CirclesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CirclesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CirclesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CirclesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CirclesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CirclesValidationError) ErrorName() string { return "CirclesValidationError" }

// Error satisfies the builtin error interface
func (e CirclesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircles.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CirclesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CirclesValidationError{}

// Validate checks the field values on Circles_ListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Circles_ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Circles_ListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Circles_ListRequestMultiError, or nil if none found.
func (m *Circles_ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *Circles_ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Availability

	if len(errors) > 0 {
		return Circles_ListRequestMultiError(errors)
	}
	return nil
}

// Circles_ListRequestMultiError is an error wrapping multiple validation
// errors returned by Circles_ListRequest.ValidateAll() if the designated
// constraints aren't met.
type Circles_ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Circles_ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Circles_ListRequestMultiError) AllErrors() []error { return m }

// Circles_ListRequestValidationError is the validation error returned by
// Circles_ListRequest.Validate if the designated constraints aren't met.
type Circles_ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Circles_ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Circles_ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Circles_ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Circles_ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Circles_ListRequestValidationError) ErrorName() string {
	return "Circles_ListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e Circles_ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircles_ListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Circles_ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Circles_ListRequestValidationError{}

// Validate checks the field values on Circles_ListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Circles_ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Circles_ListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Circles_ListResponseMultiError, or nil if none found.
func (m *Circles_ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *Circles_ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetIntersections() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Circles_ListResponseValidationError{
						field:  fmt.Sprintf("Intersections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Circles_ListResponseValidationError{
						field:  fmt.Sprintf("Intersections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Circles_ListResponseValidationError{
					field:  fmt.Sprintf("Intersections[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetListStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Circles_ListResponseValidationError{
					field:  "ListStats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Circles_ListResponseValidationError{
					field:  "ListStats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetListStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Circles_ListResponseValidationError{
				field:  "ListStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Circles_ListResponseMultiError(errors)
	}
	return nil
}

// Circles_ListResponseMultiError is an error wrapping multiple validation
// errors returned by Circles_ListResponse.ValidateAll() if the designated
// constraints aren't met.
type Circles_ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Circles_ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Circles_ListResponseMultiError) AllErrors() []error { return m }

// Circles_ListResponseValidationError is the validation error returned by
// Circles_ListResponse.Validate if the designated constraints aren't met.
type Circles_ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Circles_ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Circles_ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Circles_ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Circles_ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Circles_ListResponseValidationError) ErrorName() string {
	return "Circles_ListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e Circles_ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircles_ListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Circles_ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Circles_ListResponseValidationError{}
