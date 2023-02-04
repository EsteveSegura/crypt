# Crypt
crypt encrypts and decrypts .env files for safe transmission

# Basic usage
Encoding and decoding examples.
```bash
# Using binary
crypt -o .env -e # Encode .env file
crypt -o .env -d # Decode .env file and outputs to: .env.dec
```

.env example file (Not encoded)
```env
SAMPLE_ENV_VAR=LONGkeyNEEDStoBEencodedANDnotREADBLE
```
.env example file (Encoded)
```env
SAMPLE_ENV_VAR=[ENC--ZXlKaGJHY2lJWFZDSlNNZVNVwif09NMeo12Fzc3c1Yw==--ENC]
```

# Build binaries
It creates binaries for all major operating systems and architectures.
```bash
bash build.sh ./cmd/main.go
```