// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
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

	"H4sIAAAAAAAC/+xaW08bPxb/KpZ3HwOTUnYf8gb9s92o3RYB0q6EUOWMz2RcPPbUPhOIUL77yvZM5pob",
	"haqVeCqJj8/98vNJn2iss1wrUGjp5InmzLAMEIz/NCuE5NO/3J9C0QnNGaZ0RBXLgE7WpyNq4EchDHA6",
	"QVPAiNo4hYy5a7jMHalFI9ScrlYjCmqxkWM4O4yfUBaZimEj0wbBIZxXjtjmWlnwvjgdj90/sVYICt2f",
	"LM+liBkKraLvViv3Xc3v7wYSOqF/i2oHR+HURhfGaBNkcLCxEbljQif0nHHiVASLdDWip+N3ry/zrMAU",
	"FJZcCQQ6J/z09YV/0UgSXSjuJP7jV7j4GswCTGXmqkoBH+MLtRBGq6yUnhudg0EREoBJwcpc6DgwHBCd",
	"EEyBQIPJiAqEzA7k16j6ghnDlu5zo9ja/KfchScRYCoRklkktohjsDYpJPFXSaINmYsFqI4KPcnrEtwl",
	"ZwefvJhJEfcZ/TcFTMF0WRBhSbhCtCFaySVh3gQxk0BmS0+PwLJa1kxrCUzRUI1V6d6uG0Xdgkpd7laj",
	"ZhTP3Xk/lAe42pO+phOlnpc5lbBCIp3c3vW6gg+vJzwkoSwyLAby9dp/P6weqCJzDvZmO77O74wv3ZGv",
	"mLvRQA/eHppEKGFTcF70NvgYeWa9wMSaQ19hT0z82Ygm2mQMfV/H9ye1R4VCmIPvXBlYy+abGNFdFpSC",
	"Ki5O3Wk5QwY0lgIU7pcJgfalkokkRmfkIRVx6irLnVajjsQGGMJg3rbn5S55FTXdN+ytabt2TtOH/wFk",
	"nKHv4Ixz4WQzednyalut6sZeSn2Bh83R2tvPlWm7qjdrGLNtHPWMH/af89NVoZRQ87eUe1F3+3ZoEPg3",
	"hn0tbkQG5CEF1bbpgVlSXmt2Hs4QjlBkP1sVLZ3uPBaBuDACl9fOjhDtMz8ib/Q9KIfU/PgCZsD8q1In",
	"DNFv6EhoiWf88PRktY4pYu4ccZaLT7CsmHnEnALjnrTEzP87OrucHn2CZX2b+VsBUQmVaD+CBEp3dnFy",
	"Ts4up3REF2BscOn4+N3x2InTOSiWCzqh74/Hx2M3qxmm3rYI1ML/MYeBoHwWFgmTspmBbgK6WvCQcMrp",
	"hH4EvHBcOnD95EAsuR6rTMqvCZ3c7oCXjbawuusN3wHYuUZrckkMYGEU8L51Negfkr62MHJENWLeTuuI",
	"msnljeul1e3dys125tCIT9wwqHNtB0LzwRc7YUTBQ6dFtqNzqW0dHv+uOdd82YlMVkgUOTMYuRI7qgoc",
	"VKw9BqkQ2wd35REbt2+q3KxDq2MEPLJoHJCs8f0GPL8Bze+J3bpatTn9peN7MCU4j0uqRhuZCcXMcogv",
	"9zcTIWETV3dGKhcOcPCN5UPGB9GfQRLrLGOKE9QEHiEu0HW9EoBnufTBTbDE8BswcKfXtZzRsqHGjHr2",
	"HWKsXtfNp/iqV78nL/cWbBZrvzRvKhNJ2uj3v1klrkahW0ZPfqasQmAl4FCK+O8JU1tLM1C54rwop1Rz",
	"+bOh/9UkUZhtTtVO4E4H5msHR7jBGrTnjXesXP4p/e8KQsLscHHV/X7ewW+t8611vrXOl2mdkdfYRk/l",
	"imK1EYF+BGy1rWCqR78bYKivdL+vsef1buo5ZT/aSVhtWBxp51GlODxWheijItS81F7qOcGUIbGpLiQn",
	"M6jh6IPAdGiH6Xj+KMBXXPlCkHpuvyaJBaTNPrHeYo37i5mBUTF+jUQNS7+9MXjX3gOzt1yT76I9/S0y",
	"Paq2jcNj7Yxzwuo82T7PWln+OSwnXzvTt87C7cnTmWC5uIbYDBX92eWU2HC2ZV277xq20+lrudUy9FkN",
	"/t3AcGomN+McQghfLZtXg9mXApNhrzDYUP/tj0mcQnw/1ELD+Ya3fBeC+TwIO5q17YfZ630dVcuZPTYR",
	"JuzmSH1lwIhp4/BX7iS6i8Of3Ev0bfW+He/j2/EvxACNbVa7KdaK715kqHrp53ec/SHY74TtOD+vK22L",
	"Z3OTvX9XeBHRbbl9zNfakJab34GH3J+ZK62eED3Vu9tVZCAxYFPYMkOvAkl7kQyPCIr7akJLUGTgng9S",
	"LGB7Zk3Xsq/Wkg+dso3d88vNT14ElQceWeWJ/zk6/FTQckUNO+8hd11VLPwDykKsFXfmZexRZEVGJ+//",
	"OR6PaCZU+DgEKQem52qfhUSn7QXf8vbvD685Op+Tlv6aWVRhL4wsl/p2EkUsF8dwMjvmsKANDk/d/wVj",
	"Pd6q/7+NdWPi/wEAAP//QLRg2QYkAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
