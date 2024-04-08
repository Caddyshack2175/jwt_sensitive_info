# JSON Web Token (JWT) contains sensitive information.
##### A JSON Web Token (JWT) should not contain sensitive information. JWTs are encoded using Base64, which is not an encryption method, meaning anyone can decode and read the information within the token. Encryption is not a built-in feature of JWT, but you can use another standard called JSON Web Encryption (JWE) to encrypt the data in JWT. This provides an extra layer of security to protect sensitive data.
##### This tool looks for senitive information (username, email, etc.) stored in these tokens.
