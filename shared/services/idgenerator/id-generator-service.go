package idgenerator

import (
	"github.com/bwmarrin/snowflake"
)

type idGeneratorService struct {
	node *snowflake.Node
}

// NewIdGeneratorService creates a concrete implementation of the IdGeneratorInterface interface
func NewIdGeneratorService(node *snowflake.Node) (IdGeneratorInterface, error) {
	return &idGeneratorService{
		node: node,
	}, nil
}

// Generate genertaes an id
func (s *idGeneratorService) Generate() int64 {
	return s.node.Generate().Int64()
}
