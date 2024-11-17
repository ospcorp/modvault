package config

import (
	"testing"
)

func TestMergeOnto(t *testing.T) {
	// TODO: convert this table driven test
	ccA := Configuration{"backendA", "prefixA", "ttlA"}
	ccB := Configuration{"backendB", "", "ttlB"}

	ccResult := ccB.MergeOnto(ccA)

	backendCorrect := ccB.BackendName
	if ccResult.BackendName != backendCorrect {
		t.Errorf("Config merging failed. BackendName should be %q but was %q\n",
			backendCorrect, ccResult.BackendName)
	}

	prefixCorrect := ccA.Prefix
	if ccResult.Prefix != prefixCorrect {
		t.Errorf("Config merging failed. Prefix should be %q but was %q\n",
			prefixCorrect, ccResult.Prefix)
	}

	ttlCorrect := ccB.TTL
	if ccResult.TTL != ttlCorrect {
		t.Errorf("Config merging failed. TTL should be %q but was %q\n",
			ttlCorrect, ccResult.TTL)
	}
}
