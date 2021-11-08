/**
 * Copyright 2017,2018 InsideSales.com Inc.
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

/*
	IResult is the wrapper for a struct that is meant to be returned from a function as a more verbose error.
*/
type IResult interface {
	GetChild() IResult
	GetChildren() []chan asyncLogPackage
	WasSuccessful() bool
	Succeed()
	Fail()
	Error() string
	MergeWithResult(r IResult)
	GetMessages() []string
	GetLogLevel() int
	GetStatusCode() int
	SetStatusCode(int)
	GetResponseMessage() string
	SetResponseMessage(string)
	Flush()
	Debugf(template string, args ...interface{})
	DebugMessagef(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}
