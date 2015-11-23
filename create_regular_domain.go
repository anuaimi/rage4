package rage4

import (
  "fmt"
  // "net/url"
)

func (c *Client) CreateRegularDomain(Name string, Email string) (status Status, err error) {

  // create http request
  parameters := fmt.Sprintf("name=%s&email=%s", Name, Email)
  endpoint := fmt.Sprintf("createregulardomain/?%s", parameters)
  req, err := c.NewRequest(nil, "GET", endpoint)
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

type Status struct {
  Status      bool      `json:"status"`
  Id          int       `json:"id"`
  Error       string    `json:"error"`
}





