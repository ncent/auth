package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"

	miniredis "github.com/alicebob/miniredis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret", func() {
	var (
		response events.APIGatewayProxyResponse
		request  events.APIGatewayProxyRequest
	)

	publicKey := "045da045a8caee5819724949a3c8121931137cc355009e1f260611b4a6701c0ff729b183ecb0e42a514783b8b1d76029f1fdd690bb4b9dfabf24f703a82b1163f3"
	validSignature := "3044022052d0e24c4ef6171b4af635a8be29be3132887543c95b207eaf7be0ba6558a93c02205e7f4995f728c99db47f33c62b37975a0896814344ec8c798f5084ee3c332731"
	secret := "test"
	email := "test@test.com"

	JustBeforeEach(func() {
		mr, err := miniredis.Run()
		Expect(err).To(BeNil())
		os.Setenv("REDIS_ENDPOINT", mr.Addr())
		os.Setenv("JWT_SECRET", "Hey")
		mr.Set(publicKey, secret)
		response, err = handler(request)
		Expect(err).To(BeNil())
	})

	Context("Given the publicKey and signature are set", func() {
		BeforeEach(func() {
			request = events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"email": email, "publicKey": publicKey, "signature": validSignature},
			}
		})
		AfterEach(func() {
			request = events.APIGatewayProxyRequest{}
		})
		Context("When request is made", func() {
			It(`Then it will return a JWT and an OK status`, func() {
				Expect(response.Body).To(ContainSubstring("{\"JWT\":"))
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})

	Context("Given the publicKey is not set", func() {
		Context("When request is made", func() {
			It(`Then it will return no JWT and a Bad Request Error status`, func() {
				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Context("Given the signature is not set", func() {
		Context("When request is made", func() {
			It(`Then it will return no JWT and a Bad Request Error status`, func() {
				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
