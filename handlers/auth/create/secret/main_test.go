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

	JustBeforeEach(func() {
		mr, err := miniredis.Run()
		Expect(err).To(BeNil())
		os.Setenv("REDIS_ENDPOINT", mr.Addr())
		response, err = handler(request)
		Expect(err).To(BeNil())
	})

	Context("Given the publicKey is set", func() {
		BeforeEach(func() {
			request = events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"publicKey": "TEST_PUBLIC_KEY"},
			}
		})
		AfterEach(func() {
			request = events.APIGatewayProxyRequest{}
		})
		Context("When request is made", func() {
			It(`Then it will return Secret and an OK status`, func() {
				Expect(response.Body).To(ContainSubstring("{\"secret\":"))
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})

	Context("Given the publicKey is not set", func() {
		Context("When request is made", func() {
			It(`Then it will return no Secret and a Bad Request Error status`, func() {
				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
