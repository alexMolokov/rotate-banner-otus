// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/RotatorService.proto

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

// Validate checks the field values on AddBannerToSlotRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddBannerToSlotRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddBannerToSlotRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddBannerToSlotRequestMultiError, or nil if none found.
func (m *AddBannerToSlotRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddBannerToSlotRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBannerId() < 0 {
		err := AddBannerToSlotRequestValidationError{
			field:  "BannerId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSlotId() < 0 {
		err := AddBannerToSlotRequestValidationError{
			field:  "SlotId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddBannerToSlotRequestMultiError(errors)
	}

	return nil
}

// AddBannerToSlotRequestMultiError is an error wrapping multiple validation
// errors returned by AddBannerToSlotRequest.ValidateAll() if the designated
// constraints aren't met.
type AddBannerToSlotRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddBannerToSlotRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddBannerToSlotRequestMultiError) AllErrors() []error { return m }

// AddBannerToSlotRequestValidationError is the validation error returned by
// AddBannerToSlotRequest.Validate if the designated constraints aren't met.
type AddBannerToSlotRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddBannerToSlotRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddBannerToSlotRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddBannerToSlotRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddBannerToSlotRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddBannerToSlotRequestValidationError) ErrorName() string {
	return "AddBannerToSlotRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddBannerToSlotRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddBannerToSlotRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddBannerToSlotRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddBannerToSlotRequestValidationError{}

// Validate checks the field values on RemoveBannerFromSlotRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveBannerFromSlotRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveBannerFromSlotRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveBannerFromSlotRequestMultiError, or nil if none found.
func (m *RemoveBannerFromSlotRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveBannerFromSlotRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBannerId() < 0 {
		err := RemoveBannerFromSlotRequestValidationError{
			field:  "BannerId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSlotId() < 0 {
		err := RemoveBannerFromSlotRequestValidationError{
			field:  "SlotId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveBannerFromSlotRequestMultiError(errors)
	}

	return nil
}

// RemoveBannerFromSlotRequestMultiError is an error wrapping multiple
// validation errors returned by RemoveBannerFromSlotRequest.ValidateAll() if
// the designated constraints aren't met.
type RemoveBannerFromSlotRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveBannerFromSlotRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveBannerFromSlotRequestMultiError) AllErrors() []error { return m }

// RemoveBannerFromSlotRequestValidationError is the validation error returned
// by RemoveBannerFromSlotRequest.Validate if the designated constraints
// aren't met.
type RemoveBannerFromSlotRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveBannerFromSlotRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveBannerFromSlotRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveBannerFromSlotRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveBannerFromSlotRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveBannerFromSlotRequestValidationError) ErrorName() string {
	return "RemoveBannerFromSlotRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveBannerFromSlotRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveBannerFromSlotRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveBannerFromSlotRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveBannerFromSlotRequestValidationError{}

// Validate checks the field values on CountTransitionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CountTransitionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountTransitionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CountTransitionRequestMultiError, or nil if none found.
func (m *CountTransitionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CountTransitionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSlotId() < 0 {
		err := CountTransitionRequestValidationError{
			field:  "SlotId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetBannerId() < 0 {
		err := CountTransitionRequestValidationError{
			field:  "BannerId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSgId() < 0 {
		err := CountTransitionRequestValidationError{
			field:  "SgId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CountTransitionRequestMultiError(errors)
	}

	return nil
}

// CountTransitionRequestMultiError is an error wrapping multiple validation
// errors returned by CountTransitionRequest.ValidateAll() if the designated
// constraints aren't met.
type CountTransitionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountTransitionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountTransitionRequestMultiError) AllErrors() []error { return m }

// CountTransitionRequestValidationError is the validation error returned by
// CountTransitionRequest.Validate if the designated constraints aren't met.
type CountTransitionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountTransitionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountTransitionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountTransitionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountTransitionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountTransitionRequestValidationError) ErrorName() string {
	return "CountTransitionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CountTransitionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountTransitionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountTransitionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountTransitionRequestValidationError{}

// Validate checks the field values on ChooseBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChooseBannerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChooseBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChooseBannerRequestMultiError, or nil if none found.
func (m *ChooseBannerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChooseBannerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSlotId() < 0 {
		err := ChooseBannerRequestValidationError{
			field:  "SlotId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSgId() < 0 {
		err := ChooseBannerRequestValidationError{
			field:  "SgId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChooseBannerRequestMultiError(errors)
	}

	return nil
}

// ChooseBannerRequestMultiError is an error wrapping multiple validation
// errors returned by ChooseBannerRequest.ValidateAll() if the designated
// constraints aren't met.
type ChooseBannerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChooseBannerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChooseBannerRequestMultiError) AllErrors() []error { return m }

// ChooseBannerRequestValidationError is the validation error returned by
// ChooseBannerRequest.Validate if the designated constraints aren't met.
type ChooseBannerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChooseBannerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChooseBannerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChooseBannerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChooseBannerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChooseBannerRequestValidationError) ErrorName() string {
	return "ChooseBannerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChooseBannerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChooseBannerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChooseBannerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChooseBannerRequestValidationError{}

// Validate checks the field values on ChooseBannerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChooseBannerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChooseBannerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChooseBannerResponseMultiError, or nil if none found.
func (m *ChooseBannerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChooseBannerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BannerId

	if len(errors) > 0 {
		return ChooseBannerResponseMultiError(errors)
	}

	return nil
}

// ChooseBannerResponseMultiError is an error wrapping multiple validation
// errors returned by ChooseBannerResponse.ValidateAll() if the designated
// constraints aren't met.
type ChooseBannerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChooseBannerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChooseBannerResponseMultiError) AllErrors() []error { return m }

// ChooseBannerResponseValidationError is the validation error returned by
// ChooseBannerResponse.Validate if the designated constraints aren't met.
type ChooseBannerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChooseBannerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChooseBannerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChooseBannerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChooseBannerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChooseBannerResponseValidationError) ErrorName() string {
	return "ChooseBannerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChooseBannerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChooseBannerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChooseBannerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChooseBannerResponseValidationError{}