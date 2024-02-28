/*
General Util functions for Stacker

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"fmt"
	"time"
)

/*
Time has always been hard to mock in tests, so break some rules and make a back
door way to fix time.
*/
var _time_now_unix_ = int64(-1)

/******************************************************************************/
//MARK: - Functions

/* Add some color to the output */
func color(code int, text string) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", code, text)
}

func Time_now_unix() int64 {
	if 0 < _time_now_unix_ {
		return _time_now_unix_
	}
	return time.Now().Unix()
}
