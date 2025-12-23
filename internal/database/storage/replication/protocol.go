package replication

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

const (
	TypeMaster = "master"
	TypeSlave  = "slave"
)

// Request represents a replication request sent from a slave to a master node.
// It contains the name of the last WAL segment the slave has received.
type Request struct {
	LastSegmentName string
}

// Response represents a replication response sent from a master to a slave node.
// It contains the success status, segment name, and segment data if available.
type Response struct {
	Succeed     bool
	SegmentName string
	SegmentData []byte
}

// Encode serializes a protocol object (Request or Response) into a byte array using gob encoding.
func Encode[ProtocolObject Request | Response](object *ProtocolObject) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(object); err != nil {
		return nil, fmt.Errorf("failed to encode object: %w", err)
	}

	return buffer.Bytes(), nil
}

// Decode deserializes a byte array into a protocol object (Request or Response) using gob decoding.
func Decode[ProtocolObject Request | Response](object *ProtocolObject, data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(&object); err != nil {
		return fmt.Errorf("failed to decode object: %w", err)
	}

	return nil
}
