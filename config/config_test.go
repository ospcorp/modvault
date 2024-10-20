package config

import (
	"testing"
)

func TestMergeOnto(t *testing.T) {
	// TODO: convert this table driven test
	ccA := ConfigChoices{"backendA", "prefixA", "ttlA"}
	ccB := ConfigChoices{"backendB", "", "ttlB"}

	ccResult := ccB.MergeOnto(ccA)

	backendCorrect := ccB.DefaultBackend
	if ccResult.DefaultBackend != backendCorrect {
		t.Errorf("Config merging failed. DefaultBackend should be %q but was %q\n",
			backendCorrect, ccResult.DefaultBackend)
	}

	prefixCorrect := ccA.DefaultPrefix
	if ccResult.DefaultPrefix != prefixCorrect {
		t.Errorf("Config merging failed. DefaultPrefix should be %q but was %q\n",
			prefixCorrect, ccResult.DefaultPrefix)
	}

	ttlCorrect := ccB.DefaultTTL
	if ccResult.DefaultTTL != ttlCorrect {
		t.Errorf("Config merging failed. DefaultTTL should be %q but was %q\n",
			ttlCorrect, ccResult.DefaultTTL)
	}
}
