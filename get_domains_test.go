package rage4 

import (
  "fmt"
  "testing"
)

func TestGetDomains(t *testing.T) {

  fmt.Print("testing getdomains")
  // domains, err := client.GetDomains()
  domains, err:= client.GetDomains()
  if (err != nil) {
    t.Error("Error calling GetDomains() - ", err)
  }
  if (len(domains) == 0) {
    t.Error("Error no domains returned by GetDomains - ", err)
  }
}


