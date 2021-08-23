1. URL for Search = GET http://127.0.0.1:8089/search/(movieName)?page=2 e.g: http://127.0.0.1:8089/search/Batman?page=2
2. URL for Detail = GET http://127.0.0.1:8089/detail?movieId=(movieID) e.g: http://127.0.0.1:8089/detail?movieId=tt4853102 

GRPC in working progress, still new with GRPC, TO DO: Adjust GRPC implementation into current design architecture

Project Layout 
├── cmd
│   └── {service_name}  // service entrypoint
├── internal
│   ├── application
│   │   └── {application_name} // application logic
│   ├── domain
│   │   └── model  // data model
│   ├── interactor
│   │   ├── adapter.go // interactor interface contract
│   │   ├── grpc // grpc implementation
│   │   └── rest //rest implementation
│   └── repository
│       └── implementor
│           ├── mysql // mysql implementation
│           └── psql //psql implementation
├── pkg
│   └── {dependency_package}
├── README.md
└── script // bash script directory
