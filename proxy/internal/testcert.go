package internal

import "crypto/tls"

// LocalhostCert is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at Jan 29 16:00:00 2084 GMT.
// generated from src/crypto/tls:
// go run generate_cert.go  --rsa-bits 1024 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIICEzCCAXygAwIBAgIQS3cofn+2H4NxFntgaMRAPTANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCB
iQKBgQDp+sQVBNYwZ4YSskddAtTYq2NPdWYawNw9YQDBU9ft3fIm1r9UoyL/57bo
gCgFAkglXo06sAfuk+W6OXRPplEwxCU/mAiAjMLKES1V3oZnI42sTeiskdvb8j6E
47EpbWSA2OU4Nqulbh6vkGrzYzUdlmwwz+rGvfmHp1EOjMVzvQIDAQABo2gwZjAO
BgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUw
AwEB/zAuBgNVHREEJzAlggtleGFtcGxlLmNvbYcEfwAAAYcQAAAAAAAAAAAAAAAA
AAAAATANBgkqhkiG9w0BAQsFAAOBgQBChdgkaHaw83GFx8aDWoE3K4+h9YqXuvEP
b2OWAYlzY/U99BA9P0lE4vGpaIAeCFxalJ2AK3yHjt+eezy3sw0bMeG8ZNYcOyIV
exS95UdAKFt93a5zIWrkYQvhuzln1IOxPJQZ4rkq4nikLj2WuyGR7QnuVBdgPqP7
RN4BPb5Sog==
-----END CERTIFICATE-----`)

// LocalhostKey is the private key for LocalhostCert.
var LocalhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDp+sQVBNYwZ4YSskddAtTYq2NPdWYawNw9YQDBU9ft3fIm1r9U
oyL/57bogCgFAkglXo06sAfuk+W6OXRPplEwxCU/mAiAjMLKES1V3oZnI42sTeis
kdvb8j6E47EpbWSA2OU4Nqulbh6vkGrzYzUdlmwwz+rGvfmHp1EOjMVzvQIDAQAB
AoGAGVoXduOQRaxh5ZK1kslkwJlJaGmjB5EQDAJ/r3LjOZ3LyBOKpaQLfcjgk66X
J3vIz2vAR7SdF2elA5mIFb1CnJ4HW4cWHzgFQdUnUtoUNuMPy/9QREFfeag9GMPx
dZNiypiKqHDSY5ovUL92gtv5W0/w00lYpFiBaYLl+WHvQ6ECQQDvZpULZCEmZHwL
hZun4ObzLwFNZ9sNPgwJybnxVYaolXACeh4Ewur0kZlY9DJMqo7Rz82JWuFarkgU
GQK/L231AkEA+jP0+q7jfI8NJqwpWFDjwKiI7fadClcdUgXvW2c5wc2pEe4KiAqs
ZOWPGsH7SxigGRLzw01SCoInX5yw689JqQJARIOTPENXyWkQpyuBtLYE4qwdL039
vvh28YYuFQdpFm5ONCdG2A4AuCXDQVYB3zcg0KMsK5c6z3z5W+cchiLI0QJBAIDS
ZYz4pNoKEVxbAgKdy1XzsGTNN/gN+GO1+JJYKK23RRidNkDrNe3RIAhH3inBKRUf
4/AnjFkqwDkDRTh0htkCQQDfrRZr+gazwzDTSp23+l6MEbqBbc+TTC3c40zpNj4a
egxjd5+SkMj6zXEJxAOgo+LmQDGWsu1YQ+XXL87VPwIP
-----END RSA PRIVATE KEY-----`)

var LocalhostTLSCert tls.Certificate

func init() {
	cert, err := tls.X509KeyPair(LocalhostCert, LocalhostKey)
	if err != nil {
		panic(err)
	}
	LocalhostTLSCert = cert
}
