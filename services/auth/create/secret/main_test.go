package secret

import (
	"os"
	"time"

	miniredis "github.com/alicebob/miniredis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret", func() {
	var mr *miniredis.Miniredis
	var err error
	JustBeforeEach(func() {
		mr, err = miniredis.Run()
		Expect(err).To(BeNil())
		os.Setenv("REDIS_ENDPOINT", mr.Addr())
	})

	Context("Given the publicKey is not empty", func() {
		Context("When getting a secret", func() {
			It(`Then it will return Secret`, func() {
				secret, err := GetSecret("PUBLIC_KEY_TEST")
				Expect(err).To(BeNil())
				Expect(secret).NotTo(BeNil())
			})
		})
	})

	Context("Given the publicKey is empty", func() {
		Context("When getting a secret", func() {
			It(`Then it will return no secret`, func() {
				secret, err := GetSecret("")
				Expect(secret).To(BeNil())
				Expect(err.Error()).To(Equal("Invalid Public Key"))
			})
		})
	})

	Context("Given the same publicKey is used", func() {
		Context("When getting two secrets in an interval less or equal 60 seconds", func() {
			It(`Then it will return the same secret`, func() {
				secret1, err := GetSecret("PUBLIC_KEY_TEST")
				mr.FastForward(30 * time.Second)
				secret2, err := GetSecret("PUBLIC_KEY_TEST")
				Expect(err).To(BeNil())
				Expect(secret1).NotTo(BeNil())
				Expect(secret2).NotTo(BeNil())
				Expect(secret1).To(Equal(secret2))
			})
		})
	})

	Context("Given the same publicKey is used", func() {
		Context("When getting two secrets in an interval greater than 60 seconds", func() {
			It(`Then it will return the same secret`, func() {
				secret1, err := GetSecret("PUBLIC_KEY_TEST")
				mr.FastForward(61 * time.Second)
				secret2, err := GetSecret("PUBLIC_KEY_TEST")
				Expect(err).To(BeNil())
				Expect(secret1).NotTo(BeNil())
				Expect(secret2).NotTo(BeNil())
				Expect(secret1).NotTo(Equal(secret2))
			})
		})
	})
})
