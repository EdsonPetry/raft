package raft

import "log"

const Debug = false

func DPrintf(format string, a ...any) {
	if Debug {
		log.Printf(format, a...)
	}
}
