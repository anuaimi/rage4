package rage4 

import (
  "testing"
)

func TestGetDomains(t *testing.T) {

  domains, err:= client.GetDomains()
  if (err != nil) {
    t.Error("Error calling GetDomains() - ", err)
  }

  // should only have one domain
  if (len(domains) == 0) {
    t.Error("Error no domains returned by GetDomains")
  }
  if (len(domains) != 1) {
    t.Error("Incorrect number of domains returned")
  }

}


