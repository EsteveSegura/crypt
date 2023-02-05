package env

import (
	"testing"
)

// TODO: Mock Outputs and inputs
func TestEncodeEnv(t *testing.T) {
	envs := make(map[string]string)
	envs["SAMPLE"] = "KEY-token-superSecret"
	envs["TOKEN_SAMPLE"] = "token=secret=yes"
	EncodeEnv(envs)

}

func TestDecodeEnv(t *testing.T) {
	envs := make(map[string]string)
	envs["SAMPLE"] = "[ENC--S0VZLXRva2VuLXN1cGVyU2VjcmV0--ENC]"
	envs["TOKEN_SAMPLE"] = "[ENC--dG9rZW49c2VjcmV0PXllcw==--ENC]"
	DecodeEnv(envs)
}
