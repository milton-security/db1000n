//go:build linux
// +build linux

package ota

import (
	"strings"
	"testing"
)

//nolint: paralleltest // No need to test in parallel
func TestMergeExtraArgs(t *testing.T) {
	type testCase struct {
		Name         string
		OSArgs       []string
		ExtraArgs    []string
		ExpectedArgs []string
	}

	testCases := []testCase{
		{
			Name:         "no extra args",
			OSArgs:       []string{"--pprof=:8080"},
			ExpectedArgs: []string{"--pprof=:8080"},
		},
		{
			Name:         "unique extra args",
			OSArgs:       []string{"--pprof=:8080"},
			ExtraArgs:    []string{"--enable-self-update", "--make-yourself-comfortable"},
			ExpectedArgs: []string{"--pprof=:8080", "--enable-self-update", "--make-yourself-comfortable"},
		},
		{
			Name:         "overlapping extra args",
			OSArgs:       []string{"--pprof=:8080"},
			ExpectedArgs: []string{"--pprof=:8080", "--enable-self-update", "--make-yourself-comfortable"},
			ExtraArgs:    []string{"--pprof=:8080", "--enable-self-update", "--make-yourself-comfortable"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			mergedArgs := appendArgIfNotPresent(tc.OSArgs, tc.ExtraArgs)

			gotRawArgs := strings.Join(mergedArgs, " ")
			expectedRawArgs := strings.Join(tc.ExpectedArgs, " ")

			if gotRawArgs != expectedRawArgs {
				t.Errorf("Unexpected merge results:\nexp: %s\ngot: %s",
					expectedRawArgs, gotRawArgs)
			}
		})
	}
}
