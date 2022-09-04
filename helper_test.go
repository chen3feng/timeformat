package timeformat_test

import "testing"

func expectFalse(t *testing.T, value bool) {
	if value {
		t.Helper()
		t.Errorf("Expect false, got true")
	}
}

func expectTrue(t *testing.T, value bool) {
	if !value {
		t.Helper()
		t.Errorf("Expect true, got false")
	}
}

func expectEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Helper()
		t.Errorf("Expect %v == %v", a, b)
	}
}

func expectNe[T comparable](t *testing.T, a, b T) {
	if a == b {
		t.Helper()
		t.Errorf("Expect %v != %v", a, b)
	}
}

func expectNoError(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Errorf("Expect %v != nil", err)
	}
}
