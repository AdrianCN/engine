// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

type Align int

const (
	AlignNone   = Align(iota) // No aligh
	AlignLeft                 // Align horizontally at left
	AlignRight                // Align horizontally at right
	AlignWidth                // Align horizontally using all width
	AlignTop                  // Align vertically at the top
	AlignBottom               // Align vertically at the cnter
	AlignHeight               // Align vertically using all height
	AlignCenter               // Align horizontally/vertically at the center
)
