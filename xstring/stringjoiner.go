/*
 *    Copyright 2021 chenquan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package xstring

import "strings"

type (
	// Joiner is used to construct a sequence of characters separated by a delimiter and optionally starting with a supplied prefix and ending with a supplied suffix.
	Joiner struct {
		b    *strings.Builder
		opts *joinerOptions
		n    int // n is length of prefix and suffix for
	}
	// joinerOptions a option.
	joinerOptions struct {
		prefix string
		step   string
		suffix string
	}

	JoinerOption func(*joinerOptions)
)

// WithJoinerStep returns a JoinerOption with step settings.
func WithJoinerStep(step string) JoinerOption {
	return func(options *joinerOptions) {
		options.step = step
	}
}

// WithJoinerPrefix returns a JoinerOption with prefix settings.
func WithJoinerPrefix(prefix string) JoinerOption {
	return func(options *joinerOptions) {
		options.prefix = prefix
	}
}

// WithJoinerSuffix returns a JoinerOption with suffix settings.
func WithJoinerSuffix(suffix string) JoinerOption {
	return func(options *joinerOptions) {
		options.suffix = suffix
	}
}

// WithJoiner returns a JoinerOption with prefix, step and suffix settings.
func WithJoiner(prefix, step, suffix string) JoinerOption {
	return func(options *joinerOptions) {
		options.prefix = prefix
		options.step = step
		options.suffix = suffix
	}
}

// NewJoiner returns a Joiner.
func NewJoiner(opts ...JoinerOption) *Joiner {
	j := &Joiner{}
	j.loadOpts(opts...)

	return j
}

func (j *Joiner) loadOpts(opts ...JoinerOption) {
	op := new(joinerOptions)
	for _, opt := range opts {
		opt(op)
	}

	j.opts = op
	j.n = len(op.prefix) + len(op.suffix)
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (j *Joiner) WriteRune(r rune) (int, error) {
	j.tryWriteStep()
	n, _ := j.b.WriteRune(r)

	return n, nil
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (j *Joiner) WriteString(s string) (int, error) {
	j.tryWriteStep()
	n, _ := j.b.WriteString(s)

	return n, nil
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (j *Joiner) WriteByte(b byte) error {
	j.tryWriteStep()
	_ = j.b.WriteByte(b)

	return nil
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (j *Joiner) Write(p []byte) (int, error) {
	j.tryWriteStep()
	n, _ := j.b.Write(p)

	return n, nil
}

// String returns the accumulated string.
func (j *Joiner) String() string {
	var s string
	if j.b != nil {
		s = j.b.String()
	}

	return j.opts.prefix + s + j.opts.suffix
}

func (j *Joiner) tryWriteStep() {
	if j.b == nil {
		j.b = &strings.Builder{}
	} else {
		j.b.WriteString(j.opts.step)
	}
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (j *Joiner) Grow(n int) {
	if j.b == nil {
		j.b = &strings.Builder{}
	}

	j.b.Grow(n)
}

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (j *Joiner) Cap() int {
	if j.b == nil {
		return j.n
	}

	return j.b.Cap() + j.n
}

// Reset resets the Builder to be empty.
func (j *Joiner) Reset() {
	if j.b != nil {
		j.b.Reset()
	}
}

// Len returns the len of accumulated string.
func (j *Joiner) Len() int {
	if j.b == nil {
		return j.n
	}

	return j.b.Len() + j.n
}
