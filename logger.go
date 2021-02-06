// Copyright 2021 XinRui Hua.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rx_log

import (
	"fmt"
	"io"
	"os"
	"time"
)

func LogRequest(writer ...io.Writer) {
	reqWriter =  io.MultiWriter(writer...)
}


// normal request log
var reqWriter io.Writer
// panic log
var panicWriter io.Writer
// level log
var levelWriter io.Writer

type requestFormatter struct {
	// TimeStamp shows log current time
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int16
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
}

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param requestFormatter) string {
	return fmt.Sprintf("[RX] %v |%3d| %13v | %15s |%-7s %#v\n",
		param.TimeStamp.Format("01/02 15:04:05"),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
	)
}

func ReqLog(startT, endT time.Time, ip, method, path string, status int16) {
	if reqWriter == nil {
		reqWriter = os.Stdout
	}
	start := startT
	end := endT
	param := requestFormatter{}

	param.TimeStamp = time.Now()
	param.Latency = end.Sub(start)

	param.ClientIP = ip
	param.Method = method
	param.StatusCode = status
	param.Path = path

	fmt.Fprint(reqWriter, defaultLogFormatter(param))
}
