/**
 * Copyright 2017 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */
package routing

import (
	"crypto/rsa"
	"errors"
	"strings"

	"encoding/base64"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
)

var EXPIRED_TOKEN_ERR error = errors.New("expired")
var MISSING_TOKEN_ERR error = errors.New("missing")
var WITHHELD_TOKEN_ERR error = errors.New("withheld")

const RSA_PUB_BEG_COMMENT = "-----BEGIN PUBLIC KEY-----"
const RSA_PUB_END_COMMENT = "-----END PUBLIC KEY-----"
const RSA_PRIVATE_BEG_COMMENT = "-----BEGIN RSA PRIVATE KEY-----"
const RSA_PRIVATE_END_COMMENT = "-----END RSA PRIVATE KEY-----"

type IJwtAuthenticator interface {
	MakeJWT(claims jwtgo.Claims) (string, error)
	MakeJWTWithIssuer(claims jwtgo.Claims, issuer string) (string, error)
	MakeJWTWithHeaders(claims jwtgo.Claims, headers map[string]string) (string, error)
	MakeUnsignedJWTWithIssuer(claims jwtgo.Claims, issuer string) (string, error)
	MakeUnsignedJWTWithIssuerBase64Encoded(claims jwtgo.Claims, issuer string) (string, error)
	DecodeJwt(fullTokenString string, claims jwtgo.Claims) error
	DecodeJwtWithKeyFunc(fullTokenString string, claims jwtgo.Claims, keyFunc jwtgo.Keyfunc) error
}

//Implements IJwtAuthenticator interface
type jwtAuthenticator struct {
	/*
		public_rsa_key is the string value of the PEM encoded public RSA key.
		This is injected in the constructor GetJwtAuthenticator.
	*/
	public_rsa_key string

	/*
		private_rsa_key is the string value of the PEM encoded private RSA key.
		This is injected in the constructor GetJwtAuthenticator.
	*/
	private_rsa_key string
}

/*
	GetJwtAuthenticator is the factory method for the jwtAuthenticator class.
	@params
		public_rsa_key string PEM encoded public RSA key
		private_rsa_key string PEM encoded private RSA key
*/
func GetJwtAuthenticator(public_rsa_key string, private_rsa_key string) IJwtAuthenticator {
	return &jwtAuthenticator{
		public_rsa_key:  public_rsa_key,
		private_rsa_key: private_rsa_key,
	}
}

/*
	MakeJWT takes an implementation of the jwt-go.claims interface and signs it with the private_rsa_key.
	It then returns the jwt string.
	@params
		claims jwtgo.Claims The Jwt claims to be encoded
	@returns
		string The jwt string
		error nil if all went well
*/
func (this jwtAuthenticator) MakeJWT(claims jwtgo.Claims) (string, error) {
	new_token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)

	private_key, key_err := parsePrivateKey(this.private_rsa_key)
	if key_err != nil {
		return "", errors.New("error_trying_to_create_token")
	}

	tokenString, sign_err := new_token.SignedString(private_key)
	if sign_err != nil {
		return "", errors.New("error_trying_to_sign_token")
	}

	return tokenString, nil
}

/*
	GetJWT takes an implementation of the jwt-go.claims interface and signs it with the private_rsa_key.
	It then returns the jwt string.
	@params
		claims jwtgo.Claims The Jwt claims to be encoded
	@returns
		string The jwt string
		error nil if all went well
*/
func (this jwtAuthenticator) MakeJWTWithIssuer(claims jwtgo.Claims, issuer string) (string, error) {
	new_token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	new_token.Header["iss"] = issuer

	private_key, key_err := parsePrivateKey(this.private_rsa_key)
	if key_err != nil {
		return "", errors.New("error_trying_to_create_token")
	}

	tokenString, sign_err := new_token.SignedString(private_key)
	if sign_err != nil {
		return "", errors.New("error_trying_to_sign_token")
	}

	return tokenString, nil
}

/*
	MakeJWTWithHeaders takes an implementation of the jwt-go.claims interface and a map of headers for the token.
	It signs it with the private_rsa_key. It then returns the jwt string.
	@params
		claims jwtgo.Claims The Jwt claims to be encoded
		headers map[string]interface{} The Jwt headers to be encoded
	@returns
		string The jwt string
		error nil if all went well
*/
func (this jwtAuthenticator) MakeJWTWithHeaders(claims jwtgo.Claims, headers map[string]string) (string, error) {
	new_token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	for key, value := range headers {
		new_token.Header[key] = value
	}

	private_key, key_err := parsePrivateKey(this.private_rsa_key)
	if key_err != nil {
		return "", errors.New("error_trying_to_create_token")
	}

	tokenString, sign_err := new_token.SignedString(private_key)
	if sign_err != nil {
		return "", errors.New("error_trying_to_sign_token")
	}

	return tokenString, nil
}

/*
	MakeUnsignedJWTWithIssuer takes an implementation of the jwt-go.claims interface and creates an unsigned jwt token
	It then returns the jwt string in all three parts with an identifiable fake signature xxxx.yyyy.UnsignedToken
	@params
		claims jwtgo.Claims The Jwt claims to be encoded
		issuer string The issuer string to put in the header of the token
	@returns
		string The jwt signing string
		error nil if all went well
*/
func (this jwtAuthenticator) MakeUnsignedJWTWithIssuer(claims jwtgo.Claims, issuer string) (string, error) {
	new_token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	new_token.Header["iss"] = issuer

	tokenString, sign_err := new_token.SigningString()
	if sign_err != nil {
		return "", errors.New("error_trying_to_create_unsigned_token")
	}

	return tokenString + ".UnsignedToken", nil
}

/*
	MakeUnsignedJWTWithIssuerBase64Encoded Is the same as MakeUnsignedJWTWithIssuer, but base64 encodes the signing string
*/
func (this jwtAuthenticator) MakeUnsignedJWTWithIssuerBase64Encoded(claims jwtgo.Claims, issuer string) (string, error) {
	new_token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	new_token.Header["iss"] = issuer

	tokenString, sign_err := new_token.SigningString()
	if sign_err != nil {
		return "", errors.New("error_trying_to_create_unsigned_token")
	}

	signing_string := ".VW5zaWduZWRUb2tlbg==" //Base64 encoded "UnsignedToken"
	return tokenString + signing_string, nil
}

/*
	DecodeJwt takes a jwt token as a string (with or without "Bearer " prefix) and validates it against the
	public_rsa_key. It will parse the claims into the passed in struct and/or an error if something goes wrong
	@params
		fullTokenString string The Jwt token to be validated
		claims *jwtgo.Claims A pointer to the claims struct that implements jwtgo's claims interface.
	@returns
		error
*/
func (this jwtAuthenticator) DecodeJwt(fullTokenString string, claims jwtgo.Claims) error {
	if fullTokenString == "" {
		return MISSING_TOKEN_ERR
	}

	jwt := strings.TrimPrefix(fullTokenString, "Bearer ")

	token, err := jwtgo.ParseWithClaims(
		jwt,
		claims,
		func(token *jwtgo.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwtgo.SigningMethodRSA); !ok {
				return jwtgo.Token{}, errors.New("Unexpected signing method")
			}
			return ParsePublicKey(this.public_rsa_key)
		},
	)

	if err == nil && token.Valid {
		return nil
	} else if ve, ok := err.(*jwtgo.ValidationError); ok {
		logToken := fullTokenString
		if splitToken := strings.Split(fullTokenString, "."); len(splitToken) >= 2 {
			logToken = splitToken[0] + "." + splitToken[1] + ".SIGNATURE_REDACTED"
		}
		if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
			fmt.Printf("[WARN] Malformed Token : %s token: %s\n", err.Error(), logToken)
			return WITHHELD_TOKEN_ERR
		} else if ve.Errors&(jwtgo.ValidationErrorExpired) != 0 {
			fmt.Printf("[WARN] Expired Token : %s\n", err.Error())
			return EXPIRED_TOKEN_ERR
		} else if ve.Errors&(jwtgo.ValidationErrorClaimsInvalid) != 0 {
			fmt.Printf("[WARN] Invalid Claims : %s token: %s\n", err.Error(), logToken)
			return WITHHELD_TOKEN_ERR
		} else {
			fmt.Printf("[WARN] Invalid Token : %s token: %s\n", err.Error(), logToken)
			return WITHHELD_TOKEN_ERR
		}
	} else {
		fmt.Println("[ERROR] Token not returned from Parse!")
		return WITHHELD_TOKEN_ERR
	}
}

/*
	DecodeJwtWithKeyFunc is the same as DecodeJwt but takes a keyFunc as a parameter. See jwt-go's docs on keyfunc's for
	more information on how to use it.
	@params
		fullTokenString string The Jwt token to be validated
		claims *jwtgo.Claims A pointer to the claims struct that implements jwtgo's claims interface.
		keyFunc jwtgo.Keyfunc A function that finds the public key to use to decode the jwt
	@returns
		error
*/
func (this jwtAuthenticator) DecodeJwtWithKeyFunc(fullTokenString string, claims jwtgo.Claims, keyFunc jwtgo.Keyfunc) error {
	if fullTokenString == "" {
		return MISSING_TOKEN_ERR
	}

	jwt := strings.TrimPrefix(fullTokenString, "Bearer ")

	token, err := jwtgo.ParseWithClaims(
		jwt,
		claims,
		keyFunc,
	)

	if err == nil && token.Valid {
		return nil
	} else if ve, ok := err.(*jwtgo.ValidationError); ok {
		if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
			fmt.Printf("[WARN] Malformed Token : %v\n", err)
			return WITHHELD_TOKEN_ERR
		} else if ve.Errors&(jwtgo.ValidationErrorExpired) != 0 {
			fmt.Printf("[WARN] Expired Token : %v\n", err)
			return EXPIRED_TOKEN_ERR
		} else if ve.Errors&(jwtgo.ValidationErrorClaimsInvalid) != 0 {
			fmt.Printf("[WARN] Invalid Claims : %v\n", err)
			return WITHHELD_TOKEN_ERR
		} else {
			fmt.Printf("[WARN] Invalid Token : %v\n", err)
			return WITHHELD_TOKEN_ERR
		}
	} else {
		fmt.Println("[ERROR] Token not returned from Parse!")
		return WITHHELD_TOKEN_ERR
	}
}

/*
	ParsePublicKey is a helper function that attempts to PEM decode a given string that is Base64 encoded
	into an rsa.PublicKey.
	@params
		publicKeyValue string The string version of the public key to parse
	@returns
		*rsa.PublicKey
		error nil if all goes well
*/
func ParsePublicKey(publicKeyValue string) (*rsa.PublicKey, error) {
	if !strings.HasPrefix(publicKeyValue, RSA_PUB_BEG_COMMENT) {
		publicKeyValue = RSA_PUB_BEG_COMMENT + "\n" + publicKeyValue
	}
	if !strings.HasSuffix(publicKeyValue, RSA_PUB_END_COMMENT) {
		publicKeyValue += "\n" + RSA_PUB_END_COMMENT
	}

	publicKey, parse_err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(publicKeyValue))
	if parse_err != nil {
		return nil, errors.New("Could not parse public key from PEM")
	}

	return publicKey, nil
}

/*
	Base64ParsePublicKey is a helper function that attempts to PEM decode a given string that is Base64 encoded
	into an rsa.PublicKey.
	@params
		publicKeyValue string The string version of the public key to parse
	@returns
		*rsa.PublicKey
		error nil if all goes well
*/
func Base64ParsePublicKey(publicKeyValue string) (*rsa.PublicKey, error) {
	publicKeyStr, decode_err := base64.StdEncoding.DecodeString(publicKeyValue)
	if decode_err != nil {
		return nil, errors.New("Could not decode public key. err: " + decode_err.Error())
	}

	publicKey, parse_err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	if parse_err != nil {
		return nil, errors.New("Could not parse public key from PEM. err: " + parse_err.Error())
	}

	return publicKey, nil
}

/*
	parsePrivateKey is a helper function that attempts to PEM decode a given string that is Base64 encoded
	into an rsa.PrivateKey.
	@params
		privateKeyValue string The string version of the private key to parse
	@returns
		*rsa.PrivateKey
		error nil if all goes well
*/
func parsePrivateKey(privateKeyValue string) (*rsa.PrivateKey, error) {
	if !strings.HasPrefix(privateKeyValue, RSA_PRIVATE_BEG_COMMENT) {
		privateKeyValue = RSA_PRIVATE_BEG_COMMENT + "\n" + privateKeyValue
	}
	if !strings.HasSuffix(privateKeyValue, RSA_PRIVATE_END_COMMENT) {
		privateKeyValue += "\n" + RSA_PRIVATE_END_COMMENT
	}

	privateKey, parse_err := jwtgo.ParseRSAPrivateKeyFromPEM([]byte(privateKeyValue))
	if parse_err != nil {
		return nil, errors.New("Could not parse private key from PEM")
	}

	return privateKey, nil
}
