package jwt

import (
	"fmt"
	"os"

	miniredis "github.com/alicebob/miniredis"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jwt", func() {
	var mr *miniredis.Miniredis
	var err error
	publicKey := "045da045a8caee5819724949a3c8121931137cc355009e1f260611b4a6701c0ff729b183ecb0e42a514783b8b1d76029f1fdd690bb4b9dfabf24f703a82b1163f3"
	validSignature := "3044022052d0e24c4ef6171b4af635a8be29be3132887543c95b207eaf7be0ba6558a93c02205e7f4995f728c99db47f33c62b37975a0896814344ec8c798f5084ee3c332731"
	invalidSignature := "30460221009414138dda2e406d5199d7d8aa07bfe4791078341f0a063c0a2b8af6c6233341022100dfac62c8e82b1545b31afab64eb428f41fa9be97cd3ed9063000869d6bcee511"
	secret := "test"
	email := "test@test.com"
	JustBeforeEach(func() {
		mr, err = miniredis.Run()
		Expect(err).To(BeNil())
		os.Setenv("REDIS_ENDPOINT", mr.Addr())
		os.Setenv("JWT_SECRET", "Hey")
		mr.Set(publicKey, secret)
	})

	Context("Given the publicKey and a valid signature", func() {
		Context("When getting a JWT", func() {
			It(`Then it will return a valid JWT`, func() {
				jwtString, err := GetJWT(email, publicKey, validSignature)
				Expect(err).To(BeNil())
				token, err := jwt.Parse(*jwtString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					return []byte("Hey"), nil
				})
				Expect(err).To(BeNil())
				Expect(token.Valid).To(BeTrue())
			})
		})
	})

	Context("Given the publicKey and an invalid signature", func() {
		Context("When getting a JWT", func() {
			It(`Then it will return a NIL JWT`, func() {
				jwt, err := GetJWT(email, publicKey, invalidSignature)
				Expect(err).NotTo(BeNil())
				Expect(jwt).To(BeNil())
			})
		})
	})
})
