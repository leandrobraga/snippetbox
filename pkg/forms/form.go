package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form is a struc
type Form struct {
	url.Values
	Errors errors
}

// New is new
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required check field in form is required
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be a blank")
		}
	}
}

//MaxLength check max size field
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

//PermittedValues is this it
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)

	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

//Valid return true if errors is equal 0
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
