package avro

import (
	"encoding/json"
	"log"

	"github.com/asyncapi/parser/pkg/errs"
)

// ArrayItems Items of the array
type ArrayItems struct {
	Type                 string                   `json:"type,omitempty"`
	AdditionalProperties AdditionalPropertiesItem `json:"additionalProperties,omitempty"`
}

// SimpleArrayItems Items of the array
type SimpleArrayItems struct {
	Type string `json:"type,omitempty"`
}

// AdditionalPropertiesItem maps additional properties object
type AdditionalPropertiesItem struct {
	Type    string  `json:"type"`
	Pattern string  `json:"pattern,omitempty"`
	Min     float64 `json:"minimum,omitempty"`
	Max     float64 `json:"maximum,omitempty"`
}

// ArrayAvro maps avro array scheme
type ArrayAvro struct {
	Type        string          `json:"type,omitempty"`
	Items       ArrayItems      `json:"items,omitempty"`
	Definitions json.RawMessage `json:"definitions,omitempty"`
}

// SimpleArrayAvro maps simple array scheme
type SimpleArrayAvro struct {
	Type  string           `json:"type"`
	Items SimpleArrayItems `json:"items"`
}

// Convert transforms avro formatted message to JSONSchema
func (ra *ArrayAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	var aA ArrayAvro
	var sAa SimpleArrayAvro
	var aAbytes []byte
	switch message["items"].(type) {
	// Simple objects
	case string:
		log.Printf("String")
		aI := SimpleArrayItems{Type: message["items"].(string)}
		sAa = SimpleArrayAvro{Type: "array", Items: aI}
		aAbytes, err := json.Marshal(sAa)
		if err != nil {
			return "", errs.New(err.Error())
		}
		return string(aAbytes), nil
		// Complex objects
	case map[string]interface{}:
		log.Printf("Map")
		itemMap := message["items"].(map[string]interface{})
		aI := ArrayItems{Type: convertType(itemMap["type"].(string)), AdditionalProperties: convertValues(itemMap["values"].(string))}
		aA = ArrayAvro{Type: "array", Items: aI}
		aAbytes, err := json.Marshal(aA)
		if err != nil {
			return "", errs.New(err.Error())
		}
		return string(aAbytes), nil
	default:
		log.Printf("I don't know about type %T!\n", message)
	}

	return string(aAbytes), nil
}
