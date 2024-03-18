/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright Contributors to the cpackget project. */

package utils

import (
	"regexp"
	"strings"

	"golang.org/x/mod/semver"
)

// stripLeadingZeros takes in an extended Semantic Version-formatted string
// with optional leading zeros and strip them out. It is required in order
// to re-use `semver` Compare and Major functions
func stripLeadingZeros(version string) string {
	regex := regexp.MustCompile(`\.0*(\d+)`)
	version = strings.TrimLeft(version, "0")
	version = regex.ReplaceAllString(version, ".$1")
	return version
}

// SemverCompare extends `semver.Compare` to work with leading zeros
func SemverCompare(version1, version2 string) int {
	version1 = "v" + stripLeadingZeros(version1)
	version2 = "v" + stripLeadingZeros(version2)
	return semver.Compare(version1, version2)
}

// SemverCompareRange returns an integer comparing version to range given as low[:high]
// according to semantic version precedence.
//
// The result is 0 if version withing range, -1 if < low, 1 if > high,
func SemverCompareRange(version, vrange string) int {
	low, high, found := strings.Cut(vrange, ":")
	if found && high != "" && SemverCompare(version, high) > 0 {
		return 1
	}
	if low != "" && SemverCompare(version, low) < 0 {
		return -1
	}
	return 0
}

// SemverMajor extends `semver.Major` to work with leading zeros
func SemverMajor(version string) string {
	version = "v" + stripLeadingZeros(version)
	version = semver.Major(version)
	return strings.TrimLeft(version, "v")
}

// strips `+meta` from the supplied version string
func SemverStripMeta(version string) string {
	before, _, found := strings.Cut(version, "+")
	if found {
		return before
	}
	return version
}
