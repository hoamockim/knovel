package jwt

import (
	"crypto/rsa"
	"knovel/userprofile/presentation/util"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	RS256 = "RS256"
)

type ClaimInfo struct {
	*jwt.StandardClaims
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      []string `json:"roles"`
	Code      string   `json:"-"`
}

type jwtParse struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	parser     *jwt.Parser
}

var jwp *jwtParse

func InitJWT(publicPath, privatePath string) {
	if jwp != nil {
		return
	}

	priBytes, pubBytes := readKeys(privatePath, publicPath)
	priKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priBytes)
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubBytes)

	jwp = &jwtParse{
		privateKey: priKey,
		publicKey:  pubKey,
		parser: &jwt.Parser{
			ValidMethods:         []string{RS256},
			SkipClaimsValidation: false,
		},
	}
}

func readKeys(priKeyPath, pubKeyPath string) (priBytes, pubBytes []byte) {
	priBytes, _ = os.ReadFile(priKeyPath)
	pubBytes, _ = os.ReadFile(pubKeyPath)

	if len(priBytes) == 0 || len(pubBytes) == 0 {
		panic("jwt is invalid")
	}
	return
}

func GenerateToken(claim ClaimInfo) string {

	claim.StandardClaims = &jwt.StandardClaims{
		ExpiresAt: util.ConvertTimestampToMilliSecond(time.Now().Add(24 * time.Hour).Unix()),
		IssuedAt:  int64(util.MakeCurrentTimestampMilliSecond()),
	}

	tkn := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claim)
	var token string
	var err error
	if token, err = tkn.SignedString(jwp.privateKey); err != nil {
		log.Fatal("Cannout generate jwt token: ", err)
	}

	return token
}

func ValidateToken() bool {
	return true
}
