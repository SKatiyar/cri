/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package cri

// Version returns the current major and minor version of package protocol.
func Version() (major, minor string) {
	return {{printf "%q" .Major}}, {{printf "%q" .Minor}}
}
