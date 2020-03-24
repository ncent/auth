package crypto

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crypto", func() {
	var publicKey string
	var secret string
	JustBeforeEach(func() {
		// Key Pair generated and signature created on https://kjur.github.io/jsrsasign/sample/sample-ecdsa.html
		publicKey = "045da045a8caee5819724949a3c8121931137cc355009e1f260611b4a6701c0ff729b183ecb0e42a514783b8b1d76029f1fdd690bb4b9dfabf24f703a82b1163f3"
		secret = "test"
	})

	Context("Given a publicKey, signature and secret", func() {
		Context("When validating a valid signature", func() {
			It(`Then it will return True`, func() {
				validSignature := "3044022052d0e24c4ef6171b4af635a8be29be3132887543c95b207eaf7be0ba6558a93c02205e7f4995f728c99db47f33c62b37975a0896814344ec8c798f5084ee3c332731"
				valid, err := IsValidSignature(publicKey, validSignature, secret)
				Expect(err).To(BeNil())
				Expect(valid).To(BeTrue())
			})
		})
	})

	Context("Given a publicKey, signature and secret", func() {
		Context("When validating a non valid signature", func() {
			It(`Then it will return True`, func() {
				invalidSignature := "30460221009414138dda2e406d5199d7d8aa07bfe4791078341f0a063c0a2b8af6c6233341022100dfac62c8e82b1545b31afab64eb428f41fa9be97cd3ed9063000869d6bcee511"
				valid, err := IsValidSignature(publicKey, invalidSignature, secret)
				Expect(err).To(BeNil())
				Expect(valid).To(BeFalse())
			})
		})
	})
})
