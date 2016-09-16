// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gerrit

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var gerritGitCookieFile, gerritPassword, gerritUser, gerritURL string

func init() {
	flag.StringVar(
		&gerritGitCookieFile,
		"gerrit-gitcookie-file",
		filepath.Join(os.Getenv("HOME"), ".gitcookies"),
		"Git cookie file for gitcookiefile authentication",
	)
	flag.StringVar(
		&gerritPassword,
		"gerrit-password",
		"tbot",
		"Gerrit password for basic or digest authentication",
	)
	flag.StringVar(
		&gerritUser,
		"gerrit-user",
		"tbot",
		"Gerrit user for basic or digest authentication",
	)
	flag.StringVar(
		&gerritURL,
		"gerrit-url",
		"https://go-review.googlesource.com",
		"Gerrit URL",
	)
	flag.Parse()
}

func getChanges(t *testing.T, c *Client) {
	_, err := c.QueryChanges("is:open", QueryChangesOpt{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBasicAuth(t *testing.T) {
	getChanges(t, NewClient(gerritURL, BasicAuth(gerritUser, gerritPassword)))
}

func TestDigestAuth(t *testing.T) {
	getChanges(t, NewClient(gerritURL, DigestAuth(gerritUser, gerritPassword)))
}

func TestGitCookiesAuth(t *testing.T) {
	getChanges(t, NewClient(gerritURL, GitCookiesAuth()))
}

func TestGitCookieFileAuth(t *testing.T) {
	getChanges(t, NewClient(gerritURL, GitCookieFileAuth(gerritGitCookieFile)))
}

func TestNoAuth(t *testing.T) {
	getChanges(t, NewClient(gerritURL, NoAuth))
}
