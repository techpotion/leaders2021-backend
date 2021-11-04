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

// Validate checks the field values on PolygonAnalytics with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PolygonAnalytics) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

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

// Validate checks the field values on PolygonParkAnalytics with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonParkAnalytics) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PolygonParkAnalyticsValidationError is the validation error returned by
// PolygonParkAnalytics.Validate if the designated constraints aren't met.
type PolygonParkAnalyticsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonParkAnalyticsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonParkAnalyticsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonParkAnalyticsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonParkAnalyticsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonParkAnalyticsValidationError) ErrorName() string {
	return "PolygonParkAnalyticsValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonParkAnalyticsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonParkAnalytics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonParkAnalyticsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonParkAnalyticsValidationError{}

// Validate checks the field values on PolygonPollutionAnalytics with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonPollutionAnalytics) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PolygonPollutionAnalyticsValidationError is the validation error returned by
// PolygonPollutionAnalytics.Validate if the designated constraints aren't met.
type PolygonPollutionAnalyticsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonPollutionAnalyticsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonPollutionAnalyticsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonPollutionAnalyticsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonPollutionAnalyticsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonPollutionAnalyticsValidationError) ErrorName() string {
	return "PolygonPollutionAnalyticsValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonPollutionAnalyticsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonPollutionAnalytics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonPollutionAnalyticsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonPollutionAnalyticsValidationError{}

// Validate checks the field values on PolygonSubwayAnalytics with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonSubwayAnalytics) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PolygonSubwayAnalyticsValidationError is the validation error returned by
// PolygonSubwayAnalytics.Validate if the designated constraints aren't met.
type PolygonSubwayAnalyticsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonSubwayAnalyticsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonSubwayAnalyticsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonSubwayAnalyticsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonSubwayAnalyticsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonSubwayAnalyticsValidationError) ErrorName() string {
	return "PolygonSubwayAnalyticsValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonSubwayAnalyticsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonSubwayAnalytics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonSubwayAnalyticsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonSubwayAnalyticsValidationError{}

// Validate checks the field values on PolygonAnalyticsDashboard with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonAnalyticsDashboard) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PolygonAnalyticsDashboardValidationError is the validation error returned by
// PolygonAnalyticsDashboard.Validate if the designated constraints aren't met.
type PolygonAnalyticsDashboardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalyticsDashboardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalyticsDashboardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalyticsDashboardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalyticsDashboardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalyticsDashboardValidationError) ErrorName() string {
	return "PolygonAnalyticsDashboardValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalyticsDashboardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalyticsDashboard.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalyticsDashboardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalyticsDashboardValidationError{}

// Validate checks the field values on Marks with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Marks) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MarksValidationError is the validation error returned by Marks.Validate if
// the designated constraints aren't met.
type MarksValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MarksValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MarksValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MarksValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MarksValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MarksValidationError) ErrorName() string { return "MarksValidationError" }

// Error satisfies the builtin error interface
func (e MarksValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMarks.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MarksValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MarksValidationError{}

// Validate checks the field values on PolygonAnalytics_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonAnalytics_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalytics_RequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSportsAreaNames() {
		_, _ = idx, item

		if item != "" {

		}

	}

	for idx, item := range m.GetSportsAreaTypes() {
		_, _ = idx, item

		if item != "" {

		}

	}

	for idx, item := range m.GetSportKinds() {
		_, _ = idx, item

		if item != "" {

		}

	}

	return nil
}

// PolygonAnalytics_RequestValidationError is the validation error returned by
// PolygonAnalytics_Request.Validate if the designated constraints aren't met.
type PolygonAnalytics_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalytics_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalytics_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalytics_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalytics_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalytics_RequestValidationError) ErrorName() string {
	return "PolygonAnalytics_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalytics_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalytics_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalytics_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalytics_RequestValidationError{}

// Validate checks the field values on PolygonAnalytics_Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonAnalytics_Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AreasSquare

	// no validation rules for AreasSquarePer100K

	// no validation rules for AreasAmount

	// no validation rules for AreasAmountPer100K

	// no validation rules for SportsAmount

	// no validation rules for SportsAmountPer100K

	// no validation rules for AreaTypesAmount

	// no validation rules for SportsObjectsAmount

	// no validation rules for SportsObjectsAmountPer100K

	// no validation rules for Density

	return nil
}

// PolygonAnalytics_ResponseValidationError is the validation error returned by
// PolygonAnalytics_Response.Validate if the designated constraints aren't met.
type PolygonAnalytics_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalytics_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalytics_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalytics_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalytics_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalytics_ResponseValidationError) ErrorName() string {
	return "PolygonAnalytics_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalytics_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalytics_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalytics_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalytics_ResponseValidationError{}

// Validate checks the field values on PolygonParkAnalytics_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonParkAnalytics_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonParkAnalytics_RequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HasSportground

	return nil
}

// PolygonParkAnalytics_RequestValidationError is the validation error returned
// by PolygonParkAnalytics_Request.Validate if the designated constraints
// aren't met.
type PolygonParkAnalytics_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonParkAnalytics_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonParkAnalytics_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonParkAnalytics_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonParkAnalytics_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonParkAnalytics_RequestValidationError) ErrorName() string {
	return "PolygonParkAnalytics_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonParkAnalytics_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonParkAnalytics_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonParkAnalytics_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonParkAnalytics_RequestValidationError{}

// Validate checks the field values on PolygonParkAnalytics_Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonParkAnalytics_Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetParks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolygonParkAnalytics_ResponseValidationError{
					field:  fmt.Sprintf("Parks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetListStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonParkAnalytics_ResponseValidationError{
				field:  "ListStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PolygonParkAnalytics_ResponseValidationError is the validation error
// returned by PolygonParkAnalytics_Response.Validate if the designated
// constraints aren't met.
type PolygonParkAnalytics_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonParkAnalytics_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonParkAnalytics_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonParkAnalytics_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonParkAnalytics_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonParkAnalytics_ResponseValidationError) ErrorName() string {
	return "PolygonParkAnalytics_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonParkAnalytics_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonParkAnalytics_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonParkAnalytics_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonParkAnalytics_ResponseValidationError{}

// Validate checks the field values on PolygonPollutionAnalytics_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PolygonPollutionAnalytics_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonPollutionAnalytics_RequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsPolluted

	// no validation rules for ReturnPoints

	return nil
}

// PolygonPollutionAnalytics_RequestValidationError is the validation error
// returned by PolygonPollutionAnalytics_Request.Validate if the designated
// constraints aren't met.
type PolygonPollutionAnalytics_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonPollutionAnalytics_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonPollutionAnalytics_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonPollutionAnalytics_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonPollutionAnalytics_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonPollutionAnalytics_RequestValidationError) ErrorName() string {
	return "PolygonPollutionAnalytics_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonPollutionAnalytics_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonPollutionAnalytics_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonPollutionAnalytics_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonPollutionAnalytics_RequestValidationError{}

// Validate checks the field values on PolygonPollutionAnalytics_Response with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PolygonPollutionAnalytics_Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolygonPollutionAnalytics_ResponseValidationError{
					field:  fmt.Sprintf("Points[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PollutionPercentage

	if v, ok := interface{}(m.GetListStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonPollutionAnalytics_ResponseValidationError{
				field:  "ListStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PolygonPollutionAnalytics_ResponseValidationError is the validation error
// returned by PolygonPollutionAnalytics_Response.Validate if the designated
// constraints aren't met.
type PolygonPollutionAnalytics_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonPollutionAnalytics_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonPollutionAnalytics_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonPollutionAnalytics_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonPollutionAnalytics_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonPollutionAnalytics_ResponseValidationError) ErrorName() string {
	return "PolygonPollutionAnalytics_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonPollutionAnalytics_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonPollutionAnalytics_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonPollutionAnalytics_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonPollutionAnalytics_ResponseValidationError{}

// Validate checks the field values on PolygonSubwayAnalytics_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonSubwayAnalytics_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonSubwayAnalytics_RequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PolygonSubwayAnalytics_RequestValidationError is the validation error
// returned by PolygonSubwayAnalytics_Request.Validate if the designated
// constraints aren't met.
type PolygonSubwayAnalytics_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonSubwayAnalytics_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonSubwayAnalytics_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonSubwayAnalytics_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonSubwayAnalytics_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonSubwayAnalytics_RequestValidationError) ErrorName() string {
	return "PolygonSubwayAnalytics_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonSubwayAnalytics_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonSubwayAnalytics_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonSubwayAnalytics_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonSubwayAnalytics_RequestValidationError{}

// Validate checks the field values on PolygonSubwayAnalytics_Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PolygonSubwayAnalytics_Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolygonSubwayAnalytics_ResponseValidationError{
					field:  fmt.Sprintf("Points[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetListStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonSubwayAnalytics_ResponseValidationError{
				field:  "ListStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PolygonSubwayAnalytics_ResponseValidationError is the validation error
// returned by PolygonSubwayAnalytics_Response.Validate if the designated
// constraints aren't met.
type PolygonSubwayAnalytics_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonSubwayAnalytics_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonSubwayAnalytics_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonSubwayAnalytics_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonSubwayAnalytics_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonSubwayAnalytics_ResponseValidationError) ErrorName() string {
	return "PolygonSubwayAnalytics_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonSubwayAnalytics_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonSubwayAnalytics_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonSubwayAnalytics_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonSubwayAnalytics_ResponseValidationError{}

// Validate checks the field values on PolygonAnalyticsDashboard_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PolygonAnalyticsDashboard_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPolygon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalyticsDashboard_RequestValidationError{
				field:  "Polygon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSportsAreaNames() {
		_, _ = idx, item

		if item != "" {

		}

	}

	for idx, item := range m.GetSportsAreaTypes() {
		_, _ = idx, item

		if item != "" {

		}

	}

	for idx, item := range m.GetSportKinds() {
		_, _ = idx, item

		if item != "" {

		}

	}

	return nil
}

// PolygonAnalyticsDashboard_RequestValidationError is the validation error
// returned by PolygonAnalyticsDashboard_Request.Validate if the designated
// constraints aren't met.
type PolygonAnalyticsDashboard_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalyticsDashboard_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalyticsDashboard_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalyticsDashboard_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalyticsDashboard_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalyticsDashboard_RequestValidationError) ErrorName() string {
	return "PolygonAnalyticsDashboard_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalyticsDashboard_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalyticsDashboard_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalyticsDashboard_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalyticsDashboard_RequestValidationError{}

// Validate checks the field values on PolygonAnalyticsDashboard_Response with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PolygonAnalyticsDashboard_Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBasicAnalytics()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalyticsDashboard_ResponseValidationError{
				field:  "BasicAnalytics",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetParkAnalytics()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalyticsDashboard_ResponseValidationError{
				field:  "ParkAnalytics",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPollutionAnalytics()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalyticsDashboard_ResponseValidationError{
				field:  "PollutionAnalytics",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSubwayAnalytics()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolygonAnalyticsDashboard_ResponseValidationError{
				field:  "SubwayAnalytics",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Mark

	return nil
}

// PolygonAnalyticsDashboard_ResponseValidationError is the validation error
// returned by PolygonAnalyticsDashboard_Response.Validate if the designated
// constraints aren't met.
type PolygonAnalyticsDashboard_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolygonAnalyticsDashboard_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolygonAnalyticsDashboard_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolygonAnalyticsDashboard_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolygonAnalyticsDashboard_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolygonAnalyticsDashboard_ResponseValidationError) ErrorName() string {
	return "PolygonAnalyticsDashboard_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PolygonAnalyticsDashboard_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolygonAnalyticsDashboard_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolygonAnalyticsDashboard_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolygonAnalyticsDashboard_ResponseValidationError{}

// Validate checks the field values on Marks_GetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Marks_GetRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AreasAmountPer100K

	// no validation rules for SportsAmountPer100K

	// no validation rules for AreasSquarePer100K

	// no validation rules for SubwayDistance

	// no validation rules for PollutionPercentage

	return nil
}

// Marks_GetRequestValidationError is the validation error returned by
// Marks_GetRequest.Validate if the designated constraints aren't met.
type Marks_GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Marks_GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Marks_GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Marks_GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Marks_GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Marks_GetRequestValidationError) ErrorName() string { return "Marks_GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e Marks_GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMarks_GetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Marks_GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Marks_GetRequestValidationError{}

// Validate checks the field values on Marks_GetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Marks_GetResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mark

	return nil
}

// Marks_GetResponseValidationError is the validation error returned by
// Marks_GetResponse.Validate if the designated constraints aren't met.
type Marks_GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Marks_GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Marks_GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Marks_GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Marks_GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Marks_GetResponseValidationError) ErrorName() string {
	return "Marks_GetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e Marks_GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMarks_GetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Marks_GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Marks_GetResponseValidationError{}
