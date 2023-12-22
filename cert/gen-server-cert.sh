# Generate a private key for the server
openssl genpkey -algorithm RSA -out server-key.pem

# Generate a certificate signing request (CSR)
openssl req -new -key server-key.pem -out server-csr.pem

# Self-sign the CSR to generate a certificate
openssl x509 -req -in server-csr.pem -signkey server-key.pem -out server-cert.pem
