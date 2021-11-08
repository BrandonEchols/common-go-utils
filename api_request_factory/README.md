# API Request Factory

This package is meant to be an easy-to-use but still customizable api request framework for making requests from your
go server to other servers.

The entry point for this package is the APIRequestFactory.go.

Example usage:
```
import (
   	"ci-go-utils/api_request_factory"
   	"ci-go-utils/common"
   	"net/http"
)
   
func main() {
   	//Peripherals
   	var config common.IConfigGetter
   	hrf := hrf.GetAPIRequestFactory(&http.Client{}, config)
   	my_result := common.MakeCommonResult(config)
   
   
   	//The Simplest version
   	r := hrf.Delete("MyUrl").Do()
   	my_result.MergeWithResult(r)
   	if !r.WasSuccessful() {
   		//Do something with the error
   	}
   
   	//The async version
   	go hrf.Delete("MyUrl").DoAsync(my_result.GetChild())
   
   	//For more complex needs
   	req := hrf.Post(
   		"MyRules",
   		hrf.ValidResponses(map[int]interface{}{
   		    201: struct{}{}, 
   		    305: nil, 
   		    204: nil, 
   		    121: struct{}{},
   		}),
   		hrf.RequestBody(struct{}{}),
   		hrf.RequestFormatter(common.AuthFormatter{
   			Header: "Authorization",
   			Auth:   "1234xycasdflkn;lr...",
   		}),
   		hrf.Retry(17),
   	)
   	
   	//Do the result, but save it so we can get the http.Response
   	result := req.Do()
   	my_result.MergeWithResult(r)
   	if !result.WasSuccessful() {
   		//Do somthing with the error
   	}
    
    http_response := req.GetHttpResponse()
}
   ```
