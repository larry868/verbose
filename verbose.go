/*
verbose package generate formatted output only when verbose mode is turned on.

MessageType defines the header of the message and its color
  - INFO: `>>info:` in cyan before the message
  - WARNING: `>>warn:` in orange before the messsage
  - ALERT: `>>alert:` in red before the messsage

# Usage

	verbose.Println(verbose.INFO, "everything is ok")
	verbose.Printf(verbose.WARNING, "value should be greater than %v", value)
*/
package verbose

import (
	"fmt"
)

// Set this flag to generate verbose output
var IsOn bool = false

// Type of verbose message, defines the header of the message and its color
//   - INFO: `>>info:` in cyan before the message
//   - WARNING: `>>warn:` in orange before the messsage
//   - ALERT: `>>alert:` in red before the messsage
type MessageType int

const (
	INFO    MessageType = 0 // will output ">>info:" in cyan before the messsage
	WARNING MessageType = 1 // will output ">>warn:" in orange before the messsage
	ALERT   MessageType = 2 // will output ">>alert:" in red before the messsage
)

var messageTypeStrings []string = []string{
	">>\x1b[0;36minfo\x1b[0m:",
	">>\x1b[38;5;208mwarn\x1b[0m:",
	">>\x1b[0;31malert\x1b[0m:"}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
func Println(msgtype MessageType, params ...interface{}) {
	if !IsOn {
		return
	}
	var xparams []interface{}
	xparams = append(xparams, messageTypeStrings[msgtype])
	xparams = append(xparams, params...)
	fmt.Println(xparams...)
}

// Print formats and calls Output to print to the standard stream.
// Arguments are handled in the manner of fmt.Print.
func Print(msgtype MessageType, params ...interface{}) {
	if !IsOn {
		return
	}
	var xparams []interface{}
	xparams = append(xparams, messageTypeStrings[msgtype])
	xparams = append(xparams, params...)
	fmt.Print(xparams...)
}

// Printf formats and calls Output to print to the standard stream.
// Arguments are handled in the manner of fmt.Printf.
func Printf(msgtype MessageType, format string, params ...interface{}) {
	if !IsOn {
		return
	}
	fmt.Printf(messageTypeStrings[msgtype]+" "+format, params...)
}

// Error formats and calls Output to print to the standard stream,
// like Println with the messageType ALERT and only if err is not nil
func Error(context string, err error) error {
	if IsOn && err != nil {
		fmt.Printf("%s [%s] %s\n", messageTypeStrings[ALERT], context, err.Error())
	}
	return err
}

// Assert panic with a formated message if ok is false
func Assert(ok bool, format string, params ...interface{}) {
	if !ok {
		str := fmt.Sprintf(messageTypeStrings[ALERT]+" "+format, params...)
		panic(str)
	}
}
