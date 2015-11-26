package rage4

import (
  "fmt"
  // "net/url"
)

// NOTE: NOT YET WORKING!

func (c *Client) UpdateDomain(DomainId int, Email string) (status Status, err error) {

  // create http request
  endpoint := fmt.Sprintf("updatedomain/%d", DomainId)
  parameters := map[string]string {
    "email" : Email,
  }
  req, err := c.NewRequest(nil, "GET", endpoint, parameters)
  if err != nil {
    return Status{}, err
  }

  // issue the API request
  resp, err := c.Http.Do(req)
  if err != nil {
    return Status{}, err
  }
  defer resp.Body.Close()

  // parse the response
  getStatusResponse := Status{}
  err = decode(resp.Body, &getStatusResponse)
  if err != nil {
    return Status{}, err
  }

  status = getStatusResponse
  
  return status, nil
}





