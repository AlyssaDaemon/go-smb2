package smb2

import (
	"errors"
	"os"
	"strings"
)

const PathSeparator = '\\'

func IsPathSeparator(c uint8) bool {
	return c == '\\'
}

func base(path string) string {
	j := len(path)
	for j > 0 && IsPathSeparator(path[j-1]) {
		j--
	}

	if j == 0 {
		return ""
	}

	i := j - 1
	for i > 0 && !IsPathSeparator(path[i-1]) {
		i--
	}

	return path[i:j]
}

func dir(path string) string {
	if path == "" {
		return ""
	}

	i := len(path)
	for i > 0 && IsPathSeparator(path[i-1]) {
		i--
	}

	if i == 0 {
		return "\\"
	}

	i--
	for i > 0 && !IsPathSeparator(path[i-1]) {
		i--
	}

	if i == 0 {
		return ""
	}

	i--
	for i > 0 && IsPathSeparator(path[i-1]) {
		i--
	}

	if i == 0 {
		return "\\"
	}

	return path[:i]
}

func validatePath(op string, path string, allowAbs bool) error {
	if len(path) == 0 {
		return nil
	}

	if strings.ContainsRune(path, '/') {
		return &os.PathError{Op: op, Path: path, Err: errors.New("can't use '/' as a path separator; use '\\' instead")}
	}

	if !allowAbs && path[0] == '\\' {
		return &os.PathError{Op: op, Path: path, Err: errors.New("leading '\\' is not allowed in this operation")}
	}

	return nil
}

func normPath(path string) string {
	for strings.HasPrefix(path, `.\`) {
		path = path[2:]
	}
	if path == "." {
		return ""
	}
	return path
}
