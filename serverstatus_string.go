// Code generated by "stringer -type=serverStatus"; DO NOT EDIT.

package main

import "fmt"

const _serverStatus_name = "unknownuninitializedsealedstandbyactive"

var _serverStatus_index = [...]uint8{0, 7, 20, 26, 33, 39}

func (i serverStatus) String() string {
	if i < 0 || i >= serverStatus(len(_serverStatus_index)-1) {
		return fmt.Sprintf("serverStatus(%d)", i)
	}
	return _serverStatus_name[_serverStatus_index[i]:_serverStatus_index[i+1]]
}
