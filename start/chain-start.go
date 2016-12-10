func (t *SimpleChaincode) Init (stub shim.ChaincoedSubInterface, function string, args [string]) ([]byte, error) {
	if len(args)!= 1 {
		return nil, errors.NEw("Incorrect number of arguments. Expecting 1")

	}

	err := stub.PutState ("hello_world", []byte(args[0]))
	if err §= nil {
		return nil, err
	}

	return nil, nil

}
