package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mocks and stubs
var prexiterCalled = false
var lastExitCode int

func fakePrexiter(code int) {
	prexiterCalled = true
	lastExitCode = code
}

var lastExecArgs [][]string

func fakeExecCommand(app string, osArgs []string) {
	args := []string{app}
	args = append(args, osArgs...)
	lastExecArgs = append(lastExecArgs, args)
}

var mockedShouldFinishResult bool
var mockedShouldFinishExitCode int
var lastCallAliasesArgs []string

func fakeDetectAndCallAliases(osArgs []string) (bool, int) {
	lastCallAliasesArgs = osArgs

	return mockedShouldFinishResult, mockedShouldFinishExitCode
}

var lastProxyArgs []string

func fakeProxy(osArgs []string) {
	lastProxyArgs = osArgs
}

// After each cleanup
func afterEach() {
	prexiterCalled = false
	lastExecArgs = [][]string{}
	lastCallAliasesArgs = []string{}
	lastProxyArgs = []string{}
}

// Tests if when calling to show version, it's not calling any other functions.
func TestRouteWithVersion(t *testing.T) {
	defer afterEach()
	prexiter = fakePrexiter
	execCommand = fakeExecCommand

	Route(true, []string{"dkr", "--version"})

	// prexiter was called.
	assert.Equal(t, true, prexiterCalled)
	assert.Equal(t, 0, lastExitCode)
	// `docker --version` was called
	assert.Equal(t, []string{"docker", "--version"}, lastExecArgs[0])
	// `docker-compose --version` was called
	assert.Equal(t, []string{"docker-compose", "--version"}, lastExecArgs[1])
	// CallAliases... was NOT called
	assert.Equal(t, []string(nil), lastCallAliasesArgs)
	// Proxy was NOT called
	assert.Equal(t, []string(nil), lastProxyArgs)
}

// Tests if when calling with known alias it's calling catcher and NOT calling Proxy
func TestRouteCallAliases(t *testing.T) {
	defer afterEach()
	prexiter = fakePrexiter
	execCommand = fakeExecCommand
	detectAndCallAliases = fakeDetectAndCallAliases

	mockedShouldFinishExitCode = 1
	mockedShouldFinishResult = true
	Route(false, []string{"dkr", "killall"})

	assert.Equal(t, true, prexiterCalled)
	assert.Equal(t, 1, lastExitCode)
	assert.Equal(t, []string{"dkr", "killall"}, lastCallAliasesArgs)
	assert.Equal(t, []string{}, lastProxyArgs)
}

// Tests if when calling with simulated unknown command, it will call Proxy
func TestRouteCallProxy(t *testing.T) {
	defer afterEach()
	prexiter = fakePrexiter
	execCommand = fakeExecCommand
	detectAndCallAliases = fakeDetectAndCallAliases
	proxy = fakeProxy

	mockedShouldFinishExitCode = 0
	mockedShouldFinishResult = false
	Route(false, []string{"dkr", "ps"})

	assert.Equal(t, true, prexiterCalled)
	assert.Equal(t, 0, lastExitCode)
	assert.Equal(t, []string{"dkr", "ps"}, lastCallAliasesArgs)
	assert.Equal(t, []string{"dkr", "ps"}, lastProxyArgs)
}
