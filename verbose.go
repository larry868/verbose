// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

/*
verbose package generates formatted output only when verbose mode is turned on.

MessageType defines the header of the message and its color
  - INFO: `>>info:` in cyan before the message
  - TRACKER: `>>tracker:` in green before the messsage
  - WARNING: `>>warning:` in orange before the messsage
  - ALERT: `>>alert:` in red before the messsage

# Usage

	verbose.Println(verbose.INFO, "everything is ok")
	verbose.Printf(verbose.WARNING, "value should be greater than %v", value)
*/
package verbose

import (
	"fmt"
	"time"
)

// Set this flag to generate verbose output
var IsOn bool = false
var IsDebugging bool = false

// Type of verbose message, defines the header of the message and its color
//   - INFO: 	`>>info:` in cyan before the message
//   - TRACK: 	`>>track:` in red before the messsage
//   - WARNING: `>>warn:` in orange before the messsage
//   - ALERT: 	`>>alert:` in red before the messsage
type MessageType int

const (
	INFO    MessageType = 0 // will output ">>info:" in cyan before the messsage
	WARNING MessageType = 1 // will output ">>warning:" in orange before the messsage
	ALERT   MessageType = 2 // will output ">>alert:" in red before the messsage
	TRACK   MessageType = 3 // will output ">>track: {timestamp}" in green before the messsage
	DEBUG   MessageType = 4 // will output ">>debug: in yellow before the messsage
)

var messageTypeStrings []string = []string{
	"\x1b[0;36m>>info   :\x1b[0m",
	"\x1b[38;5;208m>>warning:\x1b[0m",
	"\x1b[0;31m>>alert  :\x1b[0m",
	"\033[1;32m>>track  :\033[0m",
	"\x1b[0;33m>>debug  :\x1b[0m",
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
func Println(msgtype MessageType, params ...interface{}) {
	if !IsOn && !(IsDebugging && msgtype == DEBUG) {
		return
	}
	var xparams []interface{}
	xparams = append(xparams, messageTypeStrings[msgtype])
	if msgtype == TRACK {
		xparams = append(xparams, time.Now().Format("20060102 15:04:05 "))
	}
	xparams = append(xparams, params...)
	fmt.Println(xparams...)
}

// Print formats and calls Output to print to the standard stream.
// Arguments are handled in the manner of fmt.Print.
func Print(msgtype MessageType, params ...interface{}) {
	if !IsOn && !(IsDebugging && msgtype == DEBUG) {
		return
	}
	var xparams []interface{}
	xparams = append(xparams, messageTypeStrings[msgtype])
	if msgtype == TRACK {
		xparams = append(xparams, time.Now().Format("20060102 15:04:05 "))
	}
	xparams = append(xparams, params...)
	fmt.Print(xparams...)
}

// Printf formats and calls Output to print to the standard stream.
// Arguments are handled in the manner of fmt.Printf.
func Printf(msgtype MessageType, format string, params ...interface{}) {
	if !IsOn && !(IsDebugging && msgtype == DEBUG) {
		return
	}
	var strtrack string
	if msgtype == TRACK {
		strtrack = time.Now().Format("20060102 15:04:05 ")
	}
	fmt.Printf(fmt.Sprintf("%s %s%s", messageTypeStrings[msgtype], strtrack, format), params...)
}

// PrintfIf formats and calls Output to print to the standard stream.
// Print out only if verbose IsOn and ok is true.
// Arguments are handled in the manner of fmt.Printf.
func PrintfIf(ok bool, msgtype MessageType, format string, params ...interface{}) {
	if !ok || (!IsOn && !(IsDebugging && msgtype == DEBUG)) {
		return
	}
	Printf(msgtype, format, params...)
}

// Error formats and calls Output to print to the standard stream,
// like Println with the messageType ALERT and only if err is not nil.
func Error(context string, err error) error {
	if IsOn && err != nil {
		fmt.Printf("%s [%s] %s\n", messageTypeStrings[ALERT], context, err.Error())
	}
	return err
}

// Assert panics with a formated message if not ok, whatever IsOn and IsDebugging.
func Assert(ok bool, format string, params ...interface{}) {
	if !ok {
		str := fmt.Sprintf(messageTypeStrings[ALERT]+" "+format, params...)
		panic(str)
	}
}

// Track generate output a time tracker and a message if verbose IsOn, otherwise does nothing.
func Track(start time.Time, format string, params ...interface{}) {
	if !IsOn {
		return
	}
	fmt.Printf(fmt.Sprintf("%s %s \033[1;32m<< %s\033[0m\n", messageTypeStrings[TRACK], format, time.Since(start)), params...)
}

// Debug prints a formated string to the standard stream only if versbode IsDebugging is true, otherwise does nothing.
func Debug(format string, params ...interface{}) {
	if !IsDebugging {
		return
	}
	fmt.Printf(fmt.Sprintf("%s %s\n", messageTypeStrings[DEBUG], format), params...)
}
