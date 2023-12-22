# Generate a private key for the client
openssl genpkey -algorithm RSA -out client-key.pem

# Generate a certificate signing request (CSR)
openssl req -new -key client-key.pem -out client-csr.pem

# Self-sign the CSR to generate a certificate
openssl x509 -req -in client-csr.pem -signkey client-key.pem -out client-cert.pem
