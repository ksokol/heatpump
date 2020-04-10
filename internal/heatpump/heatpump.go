package heatpump

import (
	"fmt"
	"heatpump/internal/cli"
	"log"
)

const operationParameters = 3003
const operationValues = 3004

type Data struct {
	Parameters *[]int32
	Values     *[]int32
}

var address string

func init() {
	cliOptions := cli.Options()

	address = fmt.Sprintf("%v:8888", cliOptions.HeatpumpIp)
}

func ReadData() (*Data, error) {
	log.Printf("reading data from %v", address)

	connection, err := NewSocketConnection(address)
	defer connection.Close()

	if err != nil {
		return &Data{}, err
	}

	parameters, err := readData(operationParameters, &connection)
	var values *[]int32

	if err == nil {
		values, err = readData(operationValues, &connection)
	}

	return &Data{parameters, values}, err
}

func readData(operationCommand int32, connection *SocketConnection) (*[]int32, error) {
	log.Printf("start reading operation %v", operationCommand)

	if _, err := connection.Write(operationCommand, 0); err != nil {
		return nil, err
	}

	if operation, err := connection.Read(); err != nil {
		return nil, err
	} else if operation != operationCommand {
		return nil, fmt.Errorf("operation command %v received. expected %v", operation, operationCommand)
	}

	if operationCommand == operationValues {
		if v, err := connection.Read(); err != nil {
			return nil, err
		} else if v != 0 {
			return nil, fmt.Errorf("command %v expected 0 value got %v", operationCommand, v)
		}
	}

	dataSize, err := connection.Read()

	if err != nil {
		return nil, err
	}

	values := make([]int32, dataSize)

	for i := int32(0); i < dataSize; i++ {
		v, err := connection.Read()

		if err != nil {
			values = nil
			break
		}

		values[i] = v
	}

	log.Printf("end reading operation %v. value size %v", operationCommand, len(values))
	return &values, err
}
