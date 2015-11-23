package rage4

import (
  "fmt"
  // "net/url"
)

// NOTE: NOT YET WORKING!

func (c *Client) UpdateRecord( recordId int, record Record) (status Status, err error) {

  // create http request
  parameters := fmt.Sprintf("email=%s", record)

  endpoint := fmt.Sprintf("updaterecord/%d?%s", recordId, parameters)
  fmt.Printf("%s\n",endpoint)
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





