/*
 * Copyright 2022 TikTok Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package analyzer

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// ErrParseFailed refers to the ANTLR file parse failure
var ErrParseFailed = errors.New("parse failed")

// ErrorListener for handling error during ANTLR parsing tree traversing
type ErrorListener struct {
	*antlr.DefaultErrorListener
}

// NewErrorListener return new error listener object
func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

// SyntaxError handles the syntax error of ANTLR file parsing
func (el *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Errorf("%w at line %d:%d - %s with error (%v)", ErrParseFailed, line, column, msg, e))
}
