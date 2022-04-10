package http

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"lh/fabric"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var doc_count uint64 = 0
var file_path = "/tmp"

func initial(c *gin.Context) {
	setCorsHeader(c)

	req := NewRequest(c.Request, nil)

	rsp := HandleinitialRequest(req)

	responseBody, err := json.Marshal(rsp.Body)
	if err != nil {
		log.Println(err)
		problemDetails := ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

func HandleinitialRequest(request *Request) *Response {

	name := request.Params["name"]
	id := request.Params["id"]

	problemDetails := initialProcedure(name, id)

	if problemDetails == nil {
		return NewResponse(http.StatusOK, nil, nil)
	} else if problemDetails != nil {
		return NewResponse(int(problemDetails.Status), nil, problemDetails)
	}

	problemDetails = &ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  "UNSPECIFIED",
	}

	return NewResponse(http.StatusForbidden, nil, problemDetails)
}
func initialProcedure(name, id string) *ProblemDetails {
	fmt.Printf("Handle initialProcedure\n")
	err := fabric.InitUser(fmt.Sprintf("%s-%s", name, id), "User1_org1", "User1_org2")
	if err != nil {

		return &ProblemDetails{
			Status: 500,
			Cause:  err.Error(),
		}
	}
	fmt.Printf("Handle initialProcedure end\n")
	return nil
}

func uploadRightdoc(c *gin.Context) {
	setCorsHeader(c)

	req := NewRequest(c.Request, c.Request.Body)

	rsp := HandleuploadRightdocRequest(req)

	responseBody, err := json.Marshal(rsp.Body)
	if err != nil {
		log.Println(err)
		problemDetails := ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

func HandleuploadRightdocRequest(request *Request) *Response {

	log.Printf("test")
	name := request.Params["name"]
	id := request.Params["id"]
	docname := request.Params["docname"]
	problemDetails := uploadRightdocProcedure(name, id, docname, request.Body)

	if problemDetails == nil {
		return NewResponse(http.StatusOK, nil, nil)
	} else if problemDetails != nil {
		return NewResponse(int(problemDetails.Status), nil, problemDetails)
	}

	problemDetails = &ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  "UNSPECIFIED",
	}

	return NewResponse(http.StatusForbidden, nil, problemDetails)
}
func uploadRightdocProcedure(name, id, docname string, Body interface{}) *ProblemDetails {
	fmt.Printf("Handle uploadRightdocProcedure")
	b, ok := Body.(io.ReadCloser)
	if !ok {
		return &ProblemDetails{
			Status: 500,
		}
	}
	filecontent, err := ioutil.ReadAll(b)
	if err != nil {

		return &ProblemDetails{
			Status: 500,
		}
	}
	key := fmt.Sprintf("%s-%s", name, id)
	savepath := fmt.Sprintf("%s/%s/%s", file_path, key, docname)
	file, err := os.OpenFile(savepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return &ProblemDetails{
			Status: 500,
		}
	}
	defer file.Close()
	file.Write(filecontent)

	h := sha256.New()
	h.Write(filecontent)
	hash := hex.EncodeToString(h.Sum(nil))
	err = fabric.UpdateRightDoc(key, name, docname, fmt.Sprintf("%d", doc_count), hash, "User1")
	if err != nil {

		return &ProblemDetails{
			Status: 500,
			Cause:  "fabric fail",
		}
	}
	doc_count++
	return nil
}
