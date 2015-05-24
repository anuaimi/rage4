This is a package to allow Go apps to use the [Rage4 DNS API](https://gbshouse.uservoice.com/knowledgebase/articles/109834-rage4-dns-developers-api)

Using the API is very straight forward.  Create a client object and then you can start calling the various Rage4 API commands.

Example usage:

```
	import (
		"github.com/anuaimi/rage4"
	)
	
	client := rage4.Client{
		AccountKey: "account_key",
		Email: "rage_login",
		URL:   "https://secure.rage4.com/rapi/",
		Http:  http.DefaultClient,
  }

  domains, err := client.GetDomains()
```
  