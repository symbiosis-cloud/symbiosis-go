# Symbiosis Golang SDK
---

## Installation ##

symbiosis-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/symbiosis-cloud/symbiosis-go
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/symbiosis-cloud/symbiosis-go"
```

and run `go get` without parameters.


## Usage ##

## Creating the client ###

You can easily instantiate the API client by providing it with a valid API key:

```go
client, err := symbiosis.NewClient(os.Getenv("SYMBIOSIS_API_KEY"))
```

### Customizing client ###
ClientOptions can be passed to symbiosis.NewClient()

for example:

```go
// Changing the symbiosis API endpoint
client, err := symbiosis.NewClient(os.Getenv("SYMBIOSIS_API_KEY"), symbiosis.WithEndpoint("https://some-other-url"))

// Setting a default timeout
client, err := symbiosis.NewClient(os.Getenv("SYMBIOSIS_API_KEY"), symbiosis.WithTimeout(time.Second * 30)))
```

### Inviting team members:

```go
members, err := client.InviteTeamMembers([]string{"test1@symbiosis.host", "test2@symbiosis.host"}, symbiosis.RoleAdmin)
```

Valid roles can be obtained by using the `symbiosis.GetValidRoles()` function which returns a map[string]bool of valid roles.

## Running tests ##
```bash
go test
```


## TODO:

* Expand readme
* Add more unit tests
* Public docs
* Add integration tests