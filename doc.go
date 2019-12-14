// Command bofh is a "Bastard Operator From Hell" style excuse server.
// Inspired by http://pages.cs.wisc.edu/~ballard/bofh.
//
// Usage
//      bofh [-addr] [-port]
// Arguments
//      -addr  listen address for the server [default: localhost]
//      -port  listen port for the server    [default: 666]
package main

import "fmt"

func usage() {
	fmt.Printf(
		"bofh: The Bastard Operator From Hell excuse server.\n" +
			"usage:\n" +
			"  bofh [-addr] [-port]\n" +
			"arguments:\n" +
			"  -addr  listen address for the server [default: localhost]\n" +
			"  -port  listen port for the server    [default: 666]\n",
	)
}
