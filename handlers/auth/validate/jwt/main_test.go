package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	createJWTHelper "gitlab.com/ncent/auth/services/auth/create/jwt"
)

var _ = Describe("the auth function", func() {
	publicKey := "045da045a8caee5819724949a3c8121931137cc355009e1f260611b4a6701c0ff729b183ecb0e42a514783b8b1d76029f1fdd690bb4b9dfabf24f703a82b1163f3"
	email := "test@test.com"
	var (
		response    events.APIGatewayCustomAuthorizerResponse
		request     events.APIGatewayCustomAuthorizerRequest
		err         error
		tokenString *string
	)

	JustBeforeEach(func() {
		os.Setenv("JWT_SECRET", "Hey")
		tokenString, err = createJWTHelper.CreateSignedToken(email, publicKey)
		response, err = handler(request)
	})

	AfterEach(func() {
		request = events.APIGatewayCustomAuthorizerRequest{}
		response = events.APIGatewayCustomAuthorizerResponse{}
	})

	Context("When the auth bearer is not set", func() {
		It("Fails auth", func() {
			Expect(err).To(MatchError("Unauthorized"))
			Expect(response).To(Equal(events.APIGatewayCustomAuthorizerResponse{}))
		})
	})

	Context("When the auth bearer is set", func() {
		Context("and auth fails", func() {
			BeforeEach(func() {
				request = events.APIGatewayCustomAuthorizerRequest{
					AuthorizationToken: "bearer token",
					MethodArn:          "testARN",
				}
			})

			It("Fails auth", func() {
				Expect(err).To(MatchError("Unauthorized"))
				Expect(response).To(Equal(events.APIGatewayCustomAuthorizerResponse{}))
			})
		})

		Context("and auth succeeds", func() {
			BeforeEach(func() {
				request = events.APIGatewayCustomAuthorizerRequest{
					AuthorizationToken: "bearer " + *tokenString,
					MethodArn:          "testARN",
				}
			})

			It("authorizes", func() {
				Expect(err).To(BeNil())
				Expect(response.PolicyDocument.Version).To(Equal("2012-10-17"))
				Expect(response.PolicyDocument.Statement).To(Equal([]events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Allow",
						Resource: []string{"testARN"},
					},
				}))
			})
		})
	})
})
