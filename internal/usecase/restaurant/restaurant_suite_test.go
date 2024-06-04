package restaurant_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRestaurant(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Restaurant Suite")
}
