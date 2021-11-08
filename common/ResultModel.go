/**
 * Copyright 2017 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */

package common

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//Implements IResult
type commsResult struct {
	debug                bool                   //True if the result should log Debug level logs
	beautify_logs        bool                   //True if the result should beautify the logs
	errors_only          bool                   //True if the result should only print results that contain errors
	messages             []string               // An array of accumulated logs with context
	was_successful       bool                   //True if the result of the operation was successful
	log_importance_level int                    //0-2, 0 if only debugs were written, 1 if Info was written, 2 if Error was written
	status_code          int                    //The Status_code that was returned from an external request (if any)
	response_message     string                 //The Error Message returned from an external request (if any)
	parent               chan asyncLogPackage   //If this is a child, this will be the parent's channel
	children             []chan asyncLogPackage //A slice of channels connected to any children that are created
}

type asyncLogPackage struct {
	messages             []string
	log_importance_level int
}

/*
	MakeCommsResult is the factory method for generating a commsResult for use.
	@params
		configs IConfigGetter A configuration getter to retrieve the logging level
	@returns
		IResult A pointer to a commsResult (which implements IResult)
*/
func MakeCommsResult(configs IConfigGetter) IResult {
	debug := false
	beautify_logs := false
	errors_only := false
	log_level := configs.SafeGetConfigVar("LOGGING_LEVEL")

	if log_level == "DEBUG" {
		debug = true
	}
	if log_level == "DEV" {
		debug = true
		beautify_logs = true
	}
	if log_level == "ERRORS_ONLY" {
		errors_only = true
	}
	return &commsResult{
		debug:         debug,
		beautify_logs: beautify_logs,
		errors_only:   errors_only,
	}
}

/*
	MakeDefaultCommsResult is used for testing purposes or to avoid having to set the logging level. Instead, it just
	set's the logging level to debug.
*/
func MakeDefaultCommsResult() IResult {
	return &commsResult{
		debug:         true,
		beautify_logs: false,
	}
}

func (this *commsResult) GetChild() IResult {
	channel := make(chan asyncLogPackage)
	child := &commsResult{
		debug:         this.debug,
		beautify_logs: this.beautify_logs,
		parent:        channel,
	}
	this.children = append(this.children, channel)

	parent_log := fmt.Sprintf("[CHILD #%s STARTED]", strconv.Itoa(len(this.children)))
	this.addLog(parent_log, "")
	child_log := fmt.Sprintf("[CHILD #%s OUTPUT]", strconv.Itoa(len(this.children)))
	child.addLog(child_log, "")

	return child
}

//WasSuccessful Returns a boolean that's true if everything went well, false if there was an error
func (this *commsResult) WasSuccessful() bool {
	return this.was_successful
}

//Succeed set's the results was_successful flag to true
func (this *commsResult) Succeed() {
	this.was_successful = true
}

//Fail set's the results was_successful flag to false
func (this *commsResult) Fail() {
	this.was_successful = false
}

//Implements error interface
func (this *commsResult) Error() string {
	if this.messages == nil || len(this.messages) == 0 {
		return ""
	}
	return this.messages[len(this.messages)-1]
}

/*
	MergeWithResult is a function for merging a result that was returned from a function call with the caller's
	result to be returned to the level above.
	@params
		r IResult The result from the function below to merge with this one
*/
func (this *commsResult) MergeWithResult(r IResult) {
	if r == nil {
		return
	}
	for _, v := range r.GetMessages() {
		this.messages = append(this.messages, v)
	}

	if this.log_importance_level < r.GetLogLevel() {
		this.log_importance_level = r.GetLogLevel()
	}

	this.children = append(this.children, r.GetChildren()...)
	this.response_message = r.GetResponseMessage()
	this.status_code = r.GetStatusCode()
}

func (this *commsResult) GetChildren() []chan asyncLogPackage {
	return this.children
}

//GetMessages Returns the []string messages in this result
func (this *commsResult) GetMessages() []string {
	return this.messages
}

//GetMessages Returns the current logging level. This goes up if Infof or Errorf are called.
func (this *commsResult) GetLogLevel() int {
	return this.log_importance_level
}

//Get's status code
func (this *commsResult) GetStatusCode() int {
	return this.status_code
}

//Set's status code
func (this *commsResult) SetStatusCode(code int) {
	this.status_code = code
}

//Get the reponse error to show to the user
func (this *commsResult) GetResponseMessage() string {
	return this.response_message
}

//Set the reponse error to show to the user
func (this *commsResult) SetResponseMessage(msg string) {
	this.response_message = msg
}

/*
	Flush is meant to be a deferred function call at the top-most level of the request. When called, it formats the
	messages that it has gathered, gets all of the messages from any children that it has, and outputs them
	either through the base fmt package so as to avoid unintentional styling, or if the result has a parent it will
	write to them.

	Sample output:

	0[Info] Message: I'm in the main function. My name is bob
 	Timestamp: 2017-11-02 13:58:18.5391358 -0600 MDT
 	Caller: C:/Users/brandon.echols/Documents/Coding/Go/src/playground/old stuff/CommsResult.go::18

	1[Info] Message: Hello There. I'm AnotherFunction
 	Timestamp: 2017-11-02 13:58:18.5401363 -0600 MDT
 	Caller: C:/Users/brandon.echols/Documents/Coding/Go/src/playground/old stuff/CommsResult.go::31
*/
func (this *commsResult) Flush() {
	go func() { //In case we have to wait for children or parents, let the calling function exit
		my_logs_length := len(this.messages)

		//For each child, get their output and append it to our own
		for i, child := range this.children {
			//Blocks until Flush is called on this child, or 5 minutes has passed
			select {
			case child_output := <-child:
				if this.log_importance_level < child_output.log_importance_level {
					this.log_importance_level = child_output.log_importance_level
				}
				this.messages = append(this.messages, child_output.messages...)
			case <-time.After(time.Minute * 5):
				this.Errorf("CHILD %d DID NOT COME HOME!! We're flushing without them", i+1)
			}
		}

		output := ""
		if this.beautify_logs {
			for i, msg := range this.messages {
				if i < my_logs_length {
					num := strconv.Itoa(i)
					this.messages[i] = num + " " + msg
				} else {
					this.messages[i] = "-" + msg
				}
				output += this.messages[i]
			}
		} else {
			for i, msg := range this.messages {
				if i < my_logs_length {
					num := strconv.Itoa(i)
					this.messages[i] = num + ") " + msg
				} else {
					this.messages[i] = "-" + msg
				}
				output += this.messages[i]
			}
			output = strings.Replace(output, "\n", "  :|: ", -1)
		}

		if this.parent != nil { //We are not the top, so we'll pass on our stuff
			log_pack := asyncLogPackage{
				messages:             this.messages,
				log_importance_level: this.log_importance_level,
			}
			//Blocks until Flush is called on parent, or 5 minutes has passed
			select {
			case this.parent <- log_pack: //Send all of our output to the parent
			case <-time.After(time.Minute * 5):
				this.Errorf("PARENT NOT LISTENING!!! We'll move on without them")
				fmt.Println(output)
			}
		} else { //We're the top so we'll print
			if !(this.log_importance_level < 2 && this.errors_only) {
				fmt.Println(output)
			}
		}

		this.messages = []string{}
	}()
}

/*
	The following methods are mimics of the ZapLogger methods, but instead of logging them out, we append the
	message and contextual information to the result's list.
	@params
		template string The formattable template string of the message
		args ...interface{} A list of arguments to inject into the template
*/
func (this *commsResult) Debugf(template string, args ...interface{}) {
	if !this.debug {
		return
	}

	original_message := fmt.Sprintf(template, args...)
	this.addLog("[Debug]", original_message)
}

func (this *commsResult) Infof(template string, args ...interface{}) {
	if this.log_importance_level < 1 {
		this.log_importance_level = 1
	}
	original_message := fmt.Sprintf(template, args...)
	this.addLog("[Info]", original_message)
}

func (this *commsResult) Errorf(template string, args ...interface{}) {
	if this.log_importance_level < 2 {
		this.log_importance_level = 2
	}

	original_message := fmt.Sprintf(template, args...)
	this.addLog("[Error]", original_message)
}

//addLog is a helper function for the *f methods.
func (this *commsResult) addLog(header string, org_msg string) {
	_, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)

	org_msg = strings.TrimSuffix(org_msg, "\n")

	output := fmt.Sprintf(
		header+" %s  %s  %s::%d",
		org_msg,
		time.Now().Format(time.RFC3339),
		fileName,
		line,
	)
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}

	this.messages = append(this.messages, output)
}

/*
	The following methods work the same as above, but we leave out the contextual information.
	@params
		template string The formattable template string of the message
		args ...interface{} A list of arguments to inject into the template
*/
func (this *commsResult) DebugMessagef(template string, args ...interface{}) {
	if !this.debug {
		return
	}

	if !strings.HasSuffix(template, "\n") {
		template += "\n"
	}

	original_message := fmt.Sprintf(template, args...)
	this.messages = append(this.messages, fmt.Sprintf(
		"[Message] %s",
		original_message,
	))
}
