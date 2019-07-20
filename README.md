# go-rachio
Golang API for Rachio sprinkler systems

Based on API definition: https://rachio.readme.io/v1.0/docs
 
```
r,err := rachio.NewClient("<token>")
if err != nil {
    log.Fatal(err)
}

person err := r.Self()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("       ID:%s\n", person.ID)
fmt.Printf("User Name:%s\n", person.Username)
fmt.Printf("Full Name:%s\n", person.FullName)
fmt.Printf("   E-Mail:%s\n", person.Email)
``` 

## Example Cli in

cmd/rachio-cli


