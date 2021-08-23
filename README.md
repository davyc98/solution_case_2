1. URL for Search = GET http://127.0.0.1:8089/search/(movieName)?page=2 e.g: http://127.0.0.1:8089/search/Batman?page=2
2. URL for Detail = GET http://127.0.0.1:8089/detail?movieId=(movieID) e.g: http://127.0.0.1:8089/detail?movieId=tt4853102 

GRPC in working progress, still new with GRPC, TO DO: Adjust GRPC implementation into current design architecture

## Architectural Approach
### Hexagonal Architecture
Using Clean Code Architectural Approach, specifically addopting [netflix's Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) 
![Netflix's Hexagonal Architecture](https://miro.medium.com/max/700/1*NfFzI7Z-E3ypn8ahESbDzw.png)

### Project Layout
```
├── cmd
│   └── {service_name}  // service entrypoint
├── docker  // dockerfiles directory
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
```
