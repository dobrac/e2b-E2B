// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xZXW/bOg/+K4Le99Kts60HGHLXrj1bcba1aLbdDMGg2EyszZY0SU5PUPi/H1CyHX+m",
	"6fqxnnPT1DZNSs9DUiR9QyOZKSlAWEOnN1QxzTKwoN0VU/wv2Fwoixdc0Cn9mYPe0IAKlgGdosC3H4A3",
	"TJRAxlDObhQ+MlZzsaJFEZRqruDn/dREMoaZ4EqBPT+tVSlmk62mtkxANfzMuYaYTq3OYbd+A8ZwKUZ1",
	"b5/fRW+BwkZJYcBhejSZ4E8khQVhPcoq5RGzXIrwu5EC7231/V/Dkk7p/8ItUaF/asIzraX2NmIwkeYK",
	"ldApPWExwSWCsbQI6NHkxePbPM5tAsKWWgl4uYD+8RQbnoFeg66MFhUhDvEzseZaigyEnVlmAe+ByDM6",
	"/UpPcp7GSFRA/2Q8hZgG9FQKoPOgy2TQU/RZxaU6paUCbbmn2FRWdm6luyrvKZVbfS21zNGu21XPDPo6",
	"/raRcMLEPQvoUuqMWTqlXNhXL2m9Jy4srMDRk4ExbDWmiAYDgdJcZWmo0oLLfZcvTpllC2YGsIkXn6/e",
	"9429kUJAhBcr4i0RK4lNgMSlqsORpZQrKRUXAf0I1w1sB1YAyv1yC5kZiNjaDNOabdw1ZCrdg9JPlVwX",
	"pFpB4K3P/TJnPqEMM9vKdG2wzmOMsyUHTeSSMEc2MV6eXCc8Ssq/3DgIYQsHYcbIiDPrPL23c4i5PRNs",
	"gYHgzC5Znlo6XbLUQDfmLtw/JMbjIuMCOeNLZ7DMlGifEZMwDTFRoA03mAIIWqlktqtYSJkCE4MetgUD",
	"oRvHLeUg9oAM1+hlh0B4Pui3zAofzA7Xh4C5c+DdhldP0UhCaJ6SNR/tvQUDnH5qxFiVnD/KGL4bGtC3",
	"kgb0hBk8jK9yg6xdbmwixSsa0GNhNij0aaPKLQwkb7fbKNfcbmYYrd5fjl1hggfX7ZVJlRTcK/4I4mIp",
	"+9Cdwnoh5Q9yfHmOr3GbwvYuDegaGXKSLw4nhxMkQioQTHE6pa/crcBVHm6JIYi1CW9agBXeaAoWhszj",
	"fe/iTc9suCF1FrXzpfO4fudMrM2bTv3UrAe/Dme/rUi4LfSK4Fbhtg8U806d9HJy1N/cLI8iMGaZp+mG",
	"eATibpD5gucJao43TAhpy2UMr+LFmPJ6qyEKbauk3bIohOuwbIV8UPQNFz2K2Sjpw+ULlF4WqjJg0z2s",
	"JBkq8cKYOXY5zCWK/gv8xQMQD26KLMs6qcKg4Ti3UTZ5YnqlsQM1kwZklxEB17gxt6H2edQnThr7DHhz",
	"ncmJjDcPFqWdqq9oH0vYpBV3TjCRAzgmTIzkl2flJkUwfFiEdTuicrszRzjBKjvsDP+870S+ifkPeNJI",
	"l/erHnWZL1JuEtfR/N4sNOwyCbDU1z8rGHCPd+4xiRKIfvT84C1Y/5z2kJj0VV15ssg1M8TUcXbHPbg1",
	"K81Nxg6SfBHGC+fZOxKkKTNkki/qZnIwL146rdi9Lvr7ebgBSrM9HppkDGag/g7uli0cbGVpbkbJfs+N",
	"JSxNSS05wPhs++yXY30+7C9741v37ixNL5ajxivIq4axmPe6+1sI0GBzLZCBJipPkdVrY3sUAFUjKEXZ",
	"quk16EEfvyd5F8o+5gle87RPrn24iGyZ3Ssaq2b4uZUCDadpBnx4U3fl+/SO+zqUF69catZo/B+xCNjO",
	"F+7ZMrY4/P28hBqWGkwyfpxdeYHW6Af+tiBiN3qzhlieAXZyKV/DzvCvuSp13icdPDJlJSz3I+3Im3nc",
	"eYCfvJfrRUoqkg7q/4TEgi8X8ah3uEEVRpunofOJ5fKcgIiV5K4oz3VKpzSxVplpGEoR+xnTYSSz0PHS",
	"/UZS+kypgVxzm8jcEqMg4ksOMVFS9xU3nPTgphrpFYcdewFdM83ZIh0axZZjZPqtMRHsDBxPO4PZ/nyy",
	"VtIcMY5pMY2cvhcQFQrRKAx4tzj4bWi4RbXefz15Pem9eim1m+ZE/osKsfIBoZzXTtvF9AMTbIVO/+VD",
	"s3psfzo1A1551umFvnzYvubalGJe/BMAAP//1gxLwaYeAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
