package adapter

type GrpcToDomain interface {
	Adapt() interface{}
}

type DomainToGrpc interface {
	Adapt() interface{}
}
