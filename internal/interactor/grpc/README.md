## Regenerate Protobufs: 

From internal/interactor directory in a terminal, run:
`protoc grpc/movie.proto --go_out=plugins=grpc:.`

We currently keep proto files in the separate repository, but I wonder if it's not better to just keep them in the service repo, like in this example. Are there any advantages to keeping them together?
