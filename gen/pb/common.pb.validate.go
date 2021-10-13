// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

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

// Validate checks the field values on Point with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Point) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Point with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PointMultiError, or nil if none found.
func (m *Point) ValidateAll() error {
	return m.validate(true)
}

func (m *Point) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Lat

	// no validation rules for Lng

	if len(errors) > 0 {
		return PointMultiError(errors)
	}
	return nil
}

// PointMultiError is an error wrapping multiple validation errors returned by
// Point.ValidateAll() if the designated constraints aren't met.
type PointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PointMultiError) AllErrors() []error { return m }

// PointValidationError is the validation error returned by Point.Validate if
// the designated constraints aren't met.
type PointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PointValidationError) ErrorName() string { return "PointValidationError" }

// Error satisfies the builtin error interface
func (e PointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PointValidationError{}

// Validate checks the field values on LineString with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LineString) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LineString with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LineStringMultiError, or
// nil if none found.
func (m *LineString) ValidateAll() error {
	return m.validate(true)
}

func (m *LineString) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPoint() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LineStringValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LineStringValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LineStringValidationError{
					field:  fmt.Sprintf("Point[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LineStringMultiError(errors)
	}
	return nil
}

// LineStringMultiError is an error wrapping multiple validation errors
// returned by LineString.ValidateAll() if the designated constraints aren't met.
type LineStringMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LineStringMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LineStringMultiError) AllErrors() []error { return m }

// LineStringValidationError is the validation error returned by
// LineString.Validate if the designated constraints aren't met.
type LineStringValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LineStringValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LineStringValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LineStringValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LineStringValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LineStringValidationError) ErrorName() string { return "LineStringValidationError" }

// Error satisfies the builtin error interface
func (e LineStringValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLineString.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LineStringValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LineStringValidationError{}

// Validate checks the field values on MultiLineString with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MultiLineString) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MultiLineString with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MultiLineStringMultiError, or nil if none found.
func (m *MultiLineString) ValidateAll() error {
	return m.validate(true)
}

func (m *MultiLineString) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLineString() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MultiLineStringValidationError{
						field:  fmt.Sprintf("LineString[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MultiLineStringValidationError{
						field:  fmt.Sprintf("LineString[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiLineStringValidationError{
					field:  fmt.Sprintf("LineString[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MultiLineStringMultiError(errors)
	}
	return nil
}

// MultiLineStringMultiError is an error wrapping multiple validation errors
// returned by MultiLineString.ValidateAll() if the designated constraints
// aren't met.
type MultiLineStringMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MultiLineStringMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MultiLineStringMultiError) AllErrors() []error { return m }

// MultiLineStringValidationError is the validation error returned by
// MultiLineString.Validate if the designated constraints aren't met.
type MultiLineStringValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiLineStringValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiLineStringValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiLineStringValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiLineStringValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiLineStringValidationError) ErrorName() string { return "MultiLineStringValidationError" }

// Error satisfies the builtin error interface
func (e MultiLineStringValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiLineString.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiLineStringValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiLineStringValidationError{}

// Validate checks the field values on MultiPoint with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MultiPoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MultiPoint with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MultiPointMultiError, or
// nil if none found.
func (m *MultiPoint) ValidateAll() error {
	return m.validate(true)
}

func (m *MultiPoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPoint() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MultiPointValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MultiPointValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiPointValidationError{
					field:  fmt.Sprintf("Point[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MultiPointMultiError(errors)
	}
	return nil
}

// MultiPointMultiError is an error wrapping multiple validation errors
// returned by MultiPoint.ValidateAll() if the designated constraints aren't met.
type MultiPointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MultiPointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MultiPointMultiError) AllErrors() []error { return m }

// MultiPointValidationError is the validation error returned by
// MultiPoint.Validate if the designated constraints aren't met.
type MultiPointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiPointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiPointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiPointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiPointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiPointValidationError) ErrorName() string { return "MultiPointValidationError" }

// Error satisfies the builtin error interface
func (e MultiPointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiPoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiPointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiPointValidationError{}

// Validate checks the field values on Polygon with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Polygon) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Polygon with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PolygonMultiError, or nil if none found.
func (m *Polygon) ValidateAll() error {
	return m.validate(true)
}

func (m *Polygon) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPoint() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PolygonValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PolygonValidationError{
						field:  fmt.Sprintf("Point[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolygonValidationError{
					field:  fmt.Sprintf("Point[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PolygonMultiError(errors)
	}
	return nil
}

// PolygonMultiError is an error wrapping multiple validation errors returned
// by Polygon.ValidateAll() if the designated constraints aren't met.
type PolygonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolygonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolygonMultiError) AllErrors() []error { return m }

// PolygonValidationError is the validation error returned by Polygon.Validate
// if the designated constraints aren't met.
type PolygonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonValidationError) ErrorName() string { return "PolygonValidationError" }

// Error satisfies the builtin error interface
func (e PolygonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonValidationError{}

// Validate checks the field values on MultiPolygon with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MultiPolygon) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MultiPolygon with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MultiPolygonMultiError, or
// nil if none found.
func (m *MultiPolygon) ValidateAll() error {
	return m.validate(true)
}

func (m *MultiPolygon) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPolygon() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MultiPolygonValidationError{
						field:  fmt.Sprintf("Polygon[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MultiPolygonValidationError{
						field:  fmt.Sprintf("Polygon[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiPolygonValidationError{
					field:  fmt.Sprintf("Polygon[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MultiPolygonMultiError(errors)
	}
	return nil
}

// MultiPolygonMultiError is an error wrapping multiple validation errors
// returned by MultiPolygon.ValidateAll() if the designated constraints aren't met.
type MultiPolygonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MultiPolygonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MultiPolygonMultiError) AllErrors() []error { return m }

// MultiPolygonValidationError is the validation error returned by
// MultiPolygon.Validate if the designated constraints aren't met.
type MultiPolygonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiPolygonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiPolygonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiPolygonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiPolygonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiPolygonValidationError) ErrorName() string { return "MultiPolygonValidationError" }

// Error satisfies the builtin error interface
func (e MultiPolygonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiPolygon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiPolygonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiPolygonValidationError{}

// Validate checks the field values on Pagination with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pagination with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginationMultiError, or
// nil if none found.
func (m *Pagination) ValidateAll() error {
	return m.validate(true)
}

func (m *Pagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageNumber

	// no validation rules for ResultsPerPage

	if len(errors) > 0 {
		return PaginationMultiError(errors)
	}
	return nil
}

// PaginationMultiError is an error wrapping multiple validation errors
// returned by Pagination.ValidateAll() if the designated constraints aren't met.
type PaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationMultiError) AllErrors() []error { return m }

// PaginationValidationError is the validation error returned by
// Pagination.Validate if the designated constraints aren't met.
type PaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationValidationError) ErrorName() string { return "PaginationValidationError" }

// Error satisfies the builtin error interface
func (e PaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationValidationError{}

// Validate checks the field values on ListStats with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListStats with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListStatsMultiError, or nil
// if none found.
func (m *ListStats) ValidateAll() error {
	return m.validate(true)
}

func (m *ListStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return ListStatsMultiError(errors)
	}
	return nil
}

// ListStatsMultiError is an error wrapping multiple validation errors returned
// by ListStats.ValidateAll() if the designated constraints aren't met.
type ListStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListStatsMultiError) AllErrors() []error { return m }

// ListStatsValidationError is the validation error returned by
// ListStats.Validate if the designated constraints aren't met.
type ListStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStatsValidationError) ErrorName() string { return "ListStatsValidationError" }

// Error satisfies the builtin error interface
func (e ListStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStatsValidationError{}

// Validate checks the field values on SportsObject with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SportsObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SportsObject with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SportsObjectMultiError, or
// nil if none found.
func (m *SportsObject) ValidateAll() error {
	return m.validate(true)
}

func (m *SportsObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ObjectId

	// no validation rules for ObjectName

	// no validation rules for ObjectAddress

	if all {
		switch v := interface{}(m.GetObjectPoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SportsObjectValidationError{
					field:  "ObjectPoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SportsObjectValidationError{
					field:  "ObjectPoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObjectPoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SportsObjectValidationError{
				field:  "ObjectPoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DepartmentalOrganizationId

	// no validation rules for DepartmentalOrganizationName

	// no validation rules for Availability

	// no validation rules for SportKind

	if len(errors) > 0 {
		return SportsObjectMultiError(errors)
	}
	return nil
}

// SportsObjectMultiError is an error wrapping multiple validation errors
// returned by SportsObject.ValidateAll() if the designated constraints aren't met.
type SportsObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SportsObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SportsObjectMultiError) AllErrors() []error { return m }

// SportsObjectValidationError is the validation error returned by
// SportsObject.Validate if the designated constraints aren't met.
type SportsObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SportsObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SportsObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SportsObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SportsObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SportsObjectValidationError) ErrorName() string { return "SportsObjectValidationError" }

// Error satisfies the builtin error interface
func (e SportsObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSportsObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SportsObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SportsObjectValidationError{}

// Validate checks the field values on SportsObjectDetailed with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SportsObjectDetailed) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SportsObjectDetailed with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SportsObjectDetailedMultiError, or nil if none found.
func (m *SportsObjectDetailed) ValidateAll() error {
	return m.validate(true)
}

func (m *SportsObjectDetailed) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ObjectId

	// no validation rules for ObjectName

	// no validation rules for SportsAreaAddress

	if all {
		switch v := interface{}(m.GetObjectPoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SportsObjectDetailedValidationError{
					field:  "ObjectPoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SportsObjectDetailedValidationError{
					field:  "ObjectPoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObjectPoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SportsObjectDetailedValidationError{
				field:  "ObjectPoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DepartmentalOrganizationId

	// no validation rules for DepartmentalOrganizationName

	// no validation rules for SportsAreaId

	// no validation rules for SportsAreaName

	// no validation rules for SportsAreaType

	// no validation rules for SportsAreaSquare

	// no validation rules for Availability

	// no validation rules for SportKind

	if len(errors) > 0 {
		return SportsObjectDetailedMultiError(errors)
	}
	return nil
}

// SportsObjectDetailedMultiError is an error wrapping multiple validation
// errors returned by SportsObjectDetailed.ValidateAll() if the designated
// constraints aren't met.
type SportsObjectDetailedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SportsObjectDetailedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SportsObjectDetailedMultiError) AllErrors() []error { return m }

// SportsObjectDetailedValidationError is the validation error returned by
// SportsObjectDetailed.Validate if the designated constraints aren't met.
type SportsObjectDetailedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SportsObjectDetailedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SportsObjectDetailedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SportsObjectDetailedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SportsObjectDetailedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SportsObjectDetailedValidationError) ErrorName() string {
	return "SportsObjectDetailedValidationError"
}

// Error satisfies the builtin error interface
func (e SportsObjectDetailedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSportsObjectDetailed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SportsObjectDetailedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SportsObjectDetailedValidationError{}
