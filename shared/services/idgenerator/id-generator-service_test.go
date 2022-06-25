package idgenerator_test

import (
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/sepisoad/robot-challange/shared/services/idgenerator"
	. "github.com/onsi/gomega"
)

func Test_Generate_Should_Return_Non_Zero_Id(t *testing.T) {
	g := NewGomegaWithT(t)

	snowflakeNode, err := snowflake.NewNode(1)
	g.Expect(err).Should(BeNil())

	sut, err := idgenerator.NewIdGeneratorService(snowflakeNode)
	g.Expect(err).Should(BeNil())

	id := sut.Generate()
	g.Expect(id).Should(BeNumerically(">", 0))
}
