package idgenerator

// IdGeneratorInterface defines contract for an id generator
type IdGeneratorInterface interface {
	Generate() int64
}
