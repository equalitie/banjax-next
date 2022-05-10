//go:build integration

package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	exit_code := m.Run()
	tearDown()
	os.Exit(exit_code)
}

func TestBanjaxEndpoint(t *testing.T) {
	banjax_resources := []TestResource{
		{"GET", "/auth_request", 200, randomXClientIP(), nil},
		{"POST", "/auth_request", 200, randomXClientIP(), nil},
		{"PUT", "/auth_request", 200, randomXClientIP(), nil},
		{"PATCH", "/auth_request", 200, randomXClientIP(), nil},
		{"HEAD", "/auth_request", 200, randomXClientIP(), nil},
		{"OPTIONS", "/auth_request", 200, randomXClientIP(), nil},
		{"DELETE", "/auth_request", 200, randomXClientIP(), nil},
		{"CONNECT", "/auth_request", 200, randomXClientIP(), nil},
		{"TRACE", "/auth_request", 200, randomXClientIP(), nil},
		{"GET", "/info", 200, randomXClientIP(), nil},
		{"GET", "/decision_lists", 200, randomXClientIP(), nil},
		{"GET", "/rate_limit_states", 200, randomXClientIP(), nil},
	}
	httpTester(t, banjax_resources)
}

func TestProtectedResources(t *testing.T) {
	defer reloadConfig(fixtureConfigTest)

	prefix := "/auth_request?path="
	httpTester(t, []TestResource{
		{"GET", "/info", 200, nil, []string{"2022-01-02"}},
		// this variation shouldn't be protected
		{"GET", prefix + "wp-adm/in", 200, randomXClientIP(), nil},
		// protected resources
		{"GET", prefix + "wp-admin", 401, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin", 401, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin//", 401, randomXClientIP(), nil},
		{"GET", prefix + "wp-admin/admin.php", 401, randomXClientIP(), nil},
		{"GET", prefix + "wp-admin/admin.php#test", 401, randomXClientIP(), nil},
		{"GET", prefix + "wp-admin/admin.php?a=1&b=2", 401, randomXClientIP(), nil},
		// exceptions
		{"GET", prefix + "wp-admin/admin-ajax.php", 200, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin/admin-ajax.php", 200, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin/admin-ajax.php?a=1", 200, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin/admin-ajax.php?a=1&b=2", 200, randomXClientIP(), nil},
		{"GET", prefix + "/wp-admin/admin-ajax.php#test", 200, randomXClientIP(), nil},
		{"GET", prefix + "wp-admin/admin-ajax.php/", 200, randomXClientIP(), nil},
	})

	reloadConfig(fixtureConfigTestReload)
	httpTester(t, []TestResource{
		{"GET", "/info", 200, nil, []string{"2022-02-03"}},
		{"GET", prefix + "wp-admin2", 401, randomXClientIP(), nil},
	})
}