package rage4

import (
  "fmt"
)

func (c *Client) SetRecordFailover( recordId int, failover bool) (status Status, err error) {

  // create http request
  parameters := fmt.Sprintf("/%d?active=%t", recordId, failover)
  endpoint := fmt.Sprintf("/setrecordfailover/%d?%s", recordId, parameters)
 
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


