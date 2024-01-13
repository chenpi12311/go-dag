// Copyright (c) 2023 - 2024 vistart. All rights reserved.
// Use of this source code is governed by Apache-2.0 license
// that can be found in the LICENSE file.

package simple

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ErrChannelNotInitialized reports when the channel list is not initialized.
var ErrChannelNotInitialized = errors.New("the channel map is not initialized")

// ErrChannelNotExist indicates that the specified channel does not exist.
type ErrChannelNotExist struct {
	name string
	error
}

func (e ErrChannelNotExist) Error() string {
	return fmt.Sprintf("channel[%s] not exist", e.name)
}

// ErrChannelInputEmpty indicates that the input channel is empty.
var ErrChannelInputEmpty = errors.New("the input channel is empty")

// ErrChannelOutputEmpty indicates that the output channel is empty.
var ErrChannelOutputEmpty = errors.New("the output channel is empty")

// ErrWorkerPanicked reports when the worker is panicked.
type ErrWorkerPanicked struct {
	panic any
	error
}

func (e ErrWorkerPanicked) Error() string {
	return fmt.Sprintf("worker panicked.")
}

// ErrChannelNameExisted indicates that the specified channel has existed.
type ErrChannelNameExisted struct {
	name string
	error
}

func (e ErrChannelNameExisted) Error() string {
	return fmt.Sprintf("the channel[%s] has existed.", e.name)
}

// ErrValueTypeMismatch defines that the data type output by the node is inconsistent with expectation.
type ErrValueTypeMismatch struct {
	expect any
	actual any
	input  string
	error
}

func (e ErrValueTypeMismatch) Error() string {
	return fmt.Sprintf("Expectation[%s] does not match actual[%s] received on channel[%s].",
		reflect.TypeOf(e.expect), reflect.TypeOf(e.actual), e.input)
}

// NewErrValueTypeMismatch instantiates an error to indicate that the type of received data on in input channel
// does not match the actual.
func NewErrValueTypeMismatch(expect, actual any, input string) ErrValueTypeMismatch {
	return ErrValueTypeMismatch{expect: expect, actual: actual, input: input}
}

// ErrRedundantChannels indicates that there are unused channelInputs.
type ErrRedundantChannels struct {
	channels []string
	error
}

func (e ErrRedundantChannels) Error() string {
	return fmt.Sprintf("Redundant channelInputs: %v", strings.Join(e.channels, ", "))
}

// ErrTransitChannelNonExist indicates that the channel(s) to be used by the specified node does not exist.
type ErrTransitChannelNonExist struct {
	transitName    string
	channelInputs  []string
	channelOutputs []string
	error
}

func (e ErrTransitChannelNonExist) Error() string {
	return fmt.Sprintf("The specified channel(s) does not exist: input[%v], output[%v]",
		strings.Join(e.channelInputs, ", "), strings.Join(e.channelOutputs, ", "))
}
