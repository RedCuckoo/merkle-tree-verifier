// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server.proto

package proto

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

// Validate checks the field values on UploadFilesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadFilesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFilesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadFilesRequestMultiError, or nil if none found.
func (m *UploadFilesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFilesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UploadFilesRequestMultiError(errors)
	}

	return nil
}

// UploadFilesRequestMultiError is an error wrapping multiple validation errors
// returned by UploadFilesRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadFilesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFilesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFilesRequestMultiError) AllErrors() []error { return m }

// UploadFilesRequestValidationError is the validation error returned by
// UploadFilesRequest.Validate if the designated constraints aren't met.
type UploadFilesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFilesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFilesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFilesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFilesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFilesRequestValidationError) ErrorName() string {
	return "UploadFilesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadFilesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFilesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFilesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFilesRequestValidationError{}

// Validate checks the field values on UploadFilesReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UploadFilesReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFilesReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadFilesReplyMultiError, or nil if none found.
func (m *UploadFilesReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFilesReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MerkleTreeRoot

	if len(errors) > 0 {
		return UploadFilesReplyMultiError(errors)
	}

	return nil
}

// UploadFilesReplyMultiError is an error wrapping multiple validation errors
// returned by UploadFilesReply.ValidateAll() if the designated constraints
// aren't met.
type UploadFilesReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFilesReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFilesReplyMultiError) AllErrors() []error { return m }

// UploadFilesReplyValidationError is the validation error returned by
// UploadFilesReply.Validate if the designated constraints aren't met.
type UploadFilesReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFilesReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFilesReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFilesReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFilesReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFilesReplyValidationError) ErrorName() string { return "UploadFilesReplyValidationError" }

// Error satisfies the builtin error interface
func (e UploadFilesReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFilesReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFilesReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFilesReplyValidationError{}

// Validate checks the field values on DownloadFileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadFileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadFileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadFileRequestMultiError, or nil if none found.
func (m *DownloadFileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadFileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileIndex

	if len(errors) > 0 {
		return DownloadFileRequestMultiError(errors)
	}

	return nil
}

// DownloadFileRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadFileRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadFileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadFileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadFileRequestMultiError) AllErrors() []error { return m }

// DownloadFileRequestValidationError is the validation error returned by
// DownloadFileRequest.Validate if the designated constraints aren't met.
type DownloadFileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadFileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadFileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadFileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadFileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadFileRequestValidationError) ErrorName() string {
	return "DownloadFileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadFileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadFileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadFileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadFileRequestValidationError{}

// Validate checks the field values on MerkleProof with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MerkleProof) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MerkleProof with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MerkleProofMultiError, or
// nil if none found.
func (m *MerkleProof) ValidateAll() error {
	return m.validate(true)
}

func (m *MerkleProof) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MerkleProofMultiError(errors)
	}

	return nil
}

// MerkleProofMultiError is an error wrapping multiple validation errors
// returned by MerkleProof.ValidateAll() if the designated constraints aren't met.
type MerkleProofMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MerkleProofMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MerkleProofMultiError) AllErrors() []error { return m }

// MerkleProofValidationError is the validation error returned by
// MerkleProof.Validate if the designated constraints aren't met.
type MerkleProofValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MerkleProofValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MerkleProofValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MerkleProofValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MerkleProofValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MerkleProofValidationError) ErrorName() string { return "MerkleProofValidationError" }

// Error satisfies the builtin error interface
func (e MerkleProofValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMerkleProof.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MerkleProofValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MerkleProofValidationError{}

// Validate checks the field values on DownloadFileReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DownloadFileReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadFileReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadFileReplyMultiError, or nil if none found.
func (m *DownloadFileReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadFileReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for File

	if all {
		switch v := interface{}(m.GetMerkleProof()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadFileReplyValidationError{
					field:  "MerkleProof",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadFileReplyValidationError{
					field:  "MerkleProof",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMerkleProof()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadFileReplyValidationError{
				field:  "MerkleProof",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadFileReplyMultiError(errors)
	}

	return nil
}

// DownloadFileReplyMultiError is an error wrapping multiple validation errors
// returned by DownloadFileReply.ValidateAll() if the designated constraints
// aren't met.
type DownloadFileReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadFileReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadFileReplyMultiError) AllErrors() []error { return m }

// DownloadFileReplyValidationError is the validation error returned by
// DownloadFileReply.Validate if the designated constraints aren't met.
type DownloadFileReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadFileReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadFileReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadFileReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadFileReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadFileReplyValidationError) ErrorName() string {
	return "DownloadFileReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadFileReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadFileReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadFileReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadFileReplyValidationError{}

// Validate checks the field values on ListRemoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRemoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRemoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRemoteRequestMultiError, or nil if none found.
func (m *ListRemoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRemoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRemoteRequestMultiError(errors)
	}

	return nil
}

// ListRemoteRequestMultiError is an error wrapping multiple validation errors
// returned by ListRemoteRequest.ValidateAll() if the designated constraints
// aren't met.
type ListRemoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRemoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRemoteRequestMultiError) AllErrors() []error { return m }

// ListRemoteRequestValidationError is the validation error returned by
// ListRemoteRequest.Validate if the designated constraints aren't met.
type ListRemoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRemoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRemoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRemoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRemoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRemoteRequestValidationError) ErrorName() string {
	return "ListRemoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRemoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRemoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRemoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRemoteRequestValidationError{}

// Validate checks the field values on ListRemoteReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRemoteReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRemoteReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRemoteReplyMultiError, or nil if none found.
func (m *ListRemoteReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRemoteReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRemoteReplyMultiError(errors)
	}

	return nil
}

// ListRemoteReplyMultiError is an error wrapping multiple validation errors
// returned by ListRemoteReply.ValidateAll() if the designated constraints
// aren't met.
type ListRemoteReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRemoteReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRemoteReplyMultiError) AllErrors() []error { return m }

// ListRemoteReplyValidationError is the validation error returned by
// ListRemoteReply.Validate if the designated constraints aren't met.
type ListRemoteReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRemoteReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRemoteReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRemoteReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRemoteReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRemoteReplyValidationError) ErrorName() string { return "ListRemoteReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListRemoteReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRemoteReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRemoteReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRemoteReplyValidationError{}

// Validate checks the field values on ResetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResetRequestMultiError, or
// nil if none found.
func (m *ResetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ResetRequestMultiError(errors)
	}

	return nil
}

// ResetRequestMultiError is an error wrapping multiple validation errors
// returned by ResetRequest.ValidateAll() if the designated constraints aren't met.
type ResetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetRequestMultiError) AllErrors() []error { return m }

// ResetRequestValidationError is the validation error returned by
// ResetRequest.Validate if the designated constraints aren't met.
type ResetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetRequestValidationError) ErrorName() string { return "ResetRequestValidationError" }

// Error satisfies the builtin error interface
func (e ResetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetRequestValidationError{}

// Validate checks the field values on ResetReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResetReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResetReplyMultiError, or
// nil if none found.
func (m *ResetReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Successful

	if len(errors) > 0 {
		return ResetReplyMultiError(errors)
	}

	return nil
}

// ResetReplyMultiError is an error wrapping multiple validation errors
// returned by ResetReply.ValidateAll() if the designated constraints aren't met.
type ResetReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetReplyMultiError) AllErrors() []error { return m }

// ResetReplyValidationError is the validation error returned by
// ResetReply.Validate if the designated constraints aren't met.
type ResetReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetReplyValidationError) ErrorName() string { return "ResetReplyValidationError" }

// Error satisfies the builtin error interface
func (e ResetReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetReplyValidationError{}
