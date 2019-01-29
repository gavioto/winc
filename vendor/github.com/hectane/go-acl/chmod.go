package acl

import (
	"os"
)

// Change the permissions of the specified file. Only the nine
// least-significant bytes are used, allowing access by the file's owner, the
// file's group, and everyone else to be explicitly controlled.
func Chmod(name string, mode os.FileMode) error {
	return Apply(
		name,
		true,
		false,
		GrantName((uint32(mode)&0700)<<23, "S-1-3-0"),
		GrantName((uint32(mode)&0070)<<26, "S-1-3-1"),
		GrantName((uint32(mode)&0007)<<29, "S-1-1-0"),
	)
}
