// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.go-test-bench.v1+json"
  ],
  "produces": [
    "application/io.goswagger.go-test-bench.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API built with go-swagger to generate intentionally vulnerable endpoints",
    "title": "swagger-bench",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "produces": [
          "text/html"
        ],
        "tags": [
          "swagger-server"
        ],
        "operationId": "root",
        "responses": {
          "200": {
            "description": "serves to display the root of the swagger API for the test bench",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/cmdInjection": {
      "get": {
        "produces": [
          "text/html"
        ],
        "tags": [
          "cmd-injection"
        ],
        "summary": "supposed to serve the frontend for the query or cookie vulns",
        "operationId": "cmdInjectionFront",
        "responses": {
          "200": {
            "description": "served front end for command injection page of Swagger API",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "error occured"
          }
        }
      }
    },
    "/cmdInjection/{command}/cookies/{safety}/": {
      "post": {
        "description": "Demonstrates the command injection vulnerability with input through Cookie\n",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "cmd-injection"
        ],
        "operationId": "postCookiesExploit",
        "parameters": [
          {
            "$ref": "#/parameters/commandParam"
          },
          {
            "$ref": "#/parameters/safetyParam"
          },
          {
            "description": "the user provided input for the query vulnerability",
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the rendered response as a string",
            "schema": {
              "description": "The response when succesful query happens",
              "type": "string"
            }
          },
          "default": {
            "description": "Error occured"
          }
        }
      }
    },
    "/cmdInjection/{command}/query/{safety}": {
      "get": {
        "description": "Demonstrates the command injection through user query\n",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "cmd-injection"
        ],
        "operationId": "getQueryExploit",
        "parameters": [
          {
            "$ref": "#/parameters/commandParam"
          },
          {
            "$ref": "#/parameters/safetyParam"
          },
          {
            "type": "string",
            "description": "the user provided input for the query vulnerability",
            "name": "input",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the rendered response as a string",
            "schema": {
              "description": "The response when succesful query happens",
              "type": "string"
            }
          },
          "default": {
            "description": "Error occured"
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  },
  "parameters": {
    "commandParam": {
      "enum": [
        "exec.Command",
        "exec.CommandContext"
      ],
      "type": "string",
      "description": "specify if exec.Command or exec.CommandContext should be invoked",
      "name": "command",
      "in": "path",
      "required": true
    },
    "safetyParam": {
      "enum": [
        "safe",
        "unsafe"
      ],
      "type": "string",
      "description": "safety qualifier",
      "name": "safety",
      "in": "path",
      "required": true
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.go-test-bench.v1+json"
  ],
  "produces": [
    "application/io.goswagger.go-test-bench.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API built with go-swagger to generate intentionally vulnerable endpoints",
    "title": "swagger-bench",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "produces": [
          "text/html"
        ],
        "tags": [
          "swagger-server"
        ],
        "operationId": "root",
        "responses": {
          "200": {
            "description": "serves to display the root of the swagger API for the test bench",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/cmdInjection": {
      "get": {
        "produces": [
          "text/html"
        ],
        "tags": [
          "cmd-injection"
        ],
        "summary": "supposed to serve the frontend for the query or cookie vulns",
        "operationId": "cmdInjectionFront",
        "responses": {
          "200": {
            "description": "served front end for command injection page of Swagger API",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "error occured"
          }
        }
      }
    },
    "/cmdInjection/{command}/cookies/{safety}/": {
      "post": {
        "description": "Demonstrates the command injection vulnerability with input through Cookie\n",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "cmd-injection"
        ],
        "operationId": "postCookiesExploit",
        "parameters": [
          {
            "enum": [
              "exec.Command",
              "exec.CommandContext"
            ],
            "type": "string",
            "description": "specify if exec.Command or exec.CommandContext should be invoked",
            "name": "command",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "safe",
              "unsafe"
            ],
            "type": "string",
            "description": "safety qualifier",
            "name": "safety",
            "in": "path",
            "required": true
          },
          {
            "description": "the user provided input for the query vulnerability",
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the rendered response as a string",
            "schema": {
              "description": "The response when succesful query happens",
              "type": "string"
            }
          },
          "default": {
            "description": "Error occured"
          }
        }
      }
    },
    "/cmdInjection/{command}/query/{safety}": {
      "get": {
        "description": "Demonstrates the command injection through user query\n",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "cmd-injection"
        ],
        "operationId": "getQueryExploit",
        "parameters": [
          {
            "enum": [
              "exec.Command",
              "exec.CommandContext"
            ],
            "type": "string",
            "description": "specify if exec.Command or exec.CommandContext should be invoked",
            "name": "command",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "safe",
              "unsafe"
            ],
            "type": "string",
            "description": "safety qualifier",
            "name": "safety",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the user provided input for the query vulnerability",
            "name": "input",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the rendered response as a string",
            "schema": {
              "description": "The response when succesful query happens",
              "type": "string"
            }
          },
          "default": {
            "description": "Error occured"
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  },
  "parameters": {
    "commandParam": {
      "enum": [
        "exec.Command",
        "exec.CommandContext"
      ],
      "type": "string",
      "description": "specify if exec.Command or exec.CommandContext should be invoked",
      "name": "command",
      "in": "path",
      "required": true
    },
    "safetyParam": {
      "enum": [
        "safe",
        "unsafe"
      ],
      "type": "string",
      "description": "safety qualifier",
      "name": "safety",
      "in": "path",
      "required": true
    }
  }
}`))
}
