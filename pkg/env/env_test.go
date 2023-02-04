package env

import (
	"reflect"
	"testing"
)

func TestEncodeEnv(t *testing.T) {
	envs := make(map[string]string)
	envs["SAMPLE"] = "KEY-token-superSecret"
	envs["TOKEN_SAMPLE"] = "token=secret=yes"

	expectedEnvs := make(map[string]string)
	expectedEnvs["SAMPLE"] = "[ENC--S0VZLXRva2VuLXN1cGVyU2VjcmV0--ENC]"
	expectedEnvs["TOKEN_SAMPLE"] = "[ENC--dG9rZW49c2VjcmV0PXllcw==--ENC]"

	encodedEnvs := EncodeEnv(envs)

	if !reflect.DeepEqual(encodedEnvs, expectedEnvs) {
		t.Error("Maps not matching")
	}
}

func TestDecodeEnv(t *testing.T) {
	envs := make(map[string]string)
	envs["SAMPLE"] = "[ENC--S0VZLXRva2VuLXN1cGVyU2VjcmV0--ENC]"
	envs["TOKEN_SAMPLE"] = "[ENC--dG9rZW49c2VjcmV0PXllcw==--ENC]"

	expectedEnvs := make(map[string]string)
	expectedEnvs["SAMPLE"] = "KEY-token-superSecret"
	expectedEnvs["TOKEN_SAMPLE"] = "token=secret=yes"

	encodedEnvs := DecodeEnv(envs)

	if !reflect.DeepEqual(encodedEnvs, expectedEnvs) {
		t.Error("Maps not matching")
	}
}
