package proto

import "context"

// RI_SDK_Connector_I2C_Check_Connection - проверка подключения устройства
func (m *GRPCClient) RI_SDK_Connector_I2C_Check_Connection(descriptor int64) (isConnected bool, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_Check_Connection(context.Background(), &RI_SDK_Connector_I2C_Check_ConnectionParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.IsConnected, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_Open - открывает соединение по указанному адресу
func (m *GRPCClient) RI_SDK_Connector_I2C_Open(descriptor int64, addr uint8) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_Open(context.Background(), &RI_SDK_Connector_I2C_OpenParams{
		Descriptor_: descriptor,
		Addr:        uint64(addr),
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_Extend - расширяет компонент группы
func (m *GRPCClient) RI_SDK_Connector_I2C_Extend(connectorDescriptor int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_Extend(context.Background(), &RI_SDK_Connector_I2C_ExtendParams{
		ConnectorDescriptor: connectorDescriptor,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_ExtendToModel - расширяет i2c адаптер до модели
func (m *GRPCClient) RI_SDK_Connector_I2C_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_ExtendToModel(context.Background(), &RI_SDK_Connector_I2C_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_State -  функция чтения состояния
func (m *GRPCClient) RI_SDK_Connector_I2C_State(descriptor int64) (state int, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_State(context.Background(), &RI_SDK_Connector_I2C_StateParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return int(resp.State), resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_CloseAll - функция закрытия всех соединений
func (m *GRPCClient) RI_SDK_Connector_I2C_CloseAll(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_CloseAll(context.Background(), &RI_SDK_Connector_I2C_CloseAllParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_Close - функция закрытия соединения
func (m *GRPCClient) RI_SDK_Connector_I2C_Close(descriptor int64, addr uint8) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_Close(context.Background(), &RI_SDK_Connector_I2C_CloseParams{
		Descriptor_: descriptor,
		Addr:        uint64(addr),
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_WriteBytes - функция для отправки байтов на I2C-устройства
func (m *GRPCClient) RI_SDK_Connector_I2C_WriteBytes(descriptor int64, addr uint8, buf []byte, length int) (wroteBytesLen int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_WriteBytes(context.Background(), &RI_SDK_Connector_I2C_WriteBytesParams{
		Descriptor_: descriptor,
		Addr:        uint64(addr),
		Buf:         buf,
		Len:         int64(length),
	})
	if err != nil {
		return
	}
	return resp.WroteBytesLen, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_ReadBytes - читает байты с I2C-устройства
func (m *GRPCClient) RI_SDK_Connector_I2C_ReadBytes(descriptor int64, addr uint8, readBytesLen int64) (buf []byte, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_ReadBytes(context.Background(), &RI_SDK_Connector_I2C_ReadBytesParams{
		Descriptor_:  descriptor,
		Addr:         uint64(addr),
		ReadBytesLen: readBytesLen,
	})
	if err != nil {
		return
	}
	return resp.Buf, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_WriteByte - отправляет байт на I2C-устройство
func (m *GRPCClient) RI_SDK_Connector_I2C_WriteByte(descriptor int64, addr uint8, value byte) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_WriteByte(context.Background(), &RI_SDK_Connector_I2C_WriteByteParams{
		Descriptor_: descriptor,
		Addr:        uint64(addr),
		Value:       []byte{value},
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_ReadByte - читает байт с I2C-устройства
func (m *GRPCClient) RI_SDK_Connector_I2C_ReadByte(descriptor int64, addr uint8) (readedByte byte, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_I2C_ReadByte(context.Background(), &RI_SDK_Connector_I2C_ReadByteParams{
		Descriptor_: descriptor,
		Addr:        uint64(addr),
	})
	if err != nil {
		return
	}
	return resp.Value[0], resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_Open - открывает соединение по указанному адресу
func (m *GRPCServer) RI_SDK_Connector_I2C_Check_Connection(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_Check_ConnectionParams) (*RI_SDK_Connector_I2C_Check_ConnectionReturn, error) {
	isConnected, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_Check_Connection(req.Descriptor_)
	return &RI_SDK_Connector_I2C_Check_ConnectionReturn{
		IsConnected: isConnected,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Connector_I2C_Open - открывает соединение по указанному адресу
func (m *GRPCServer) RI_SDK_Connector_I2C_Open(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_OpenParams) (*RI_SDK_Connector_I2C_OpenReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_Open(req.Descriptor_, uint8(req.Addr))
	return &RI_SDK_Connector_I2C_OpenReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Connector_I2C_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Connector_I2C_Extend(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_ExtendParams) (*RI_SDK_Connector_I2C_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_Extend(req.ConnectorDescriptor)
	return &RI_SDK_Connector_I2C_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Connector_I2C_ExtendToModel - расширяет i2c адаптер до модели
func (m *GRPCServer) RI_SDK_Connector_I2C_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_ExtendToModelParams) (*RI_SDK_Connector_I2C_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Connector_I2C_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Connector_I2C_State -  функция чтения состояния
func (m *GRPCServer) RI_SDK_Connector_I2C_State(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_StateParams) (*RI_SDK_Connector_I2C_StateReturn, error) {
	state, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_State(req.Descriptor_)
	return &RI_SDK_Connector_I2C_StateReturn{
		State:     int64(state),
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Connector_I2C_CloseAll - функция закрытия всех соединений
func (m *GRPCServer) RI_SDK_Connector_I2C_CloseAll(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_CloseAllParams) (*RI_SDK_Connector_I2C_CloseAllReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_CloseAll(req.Descriptor_)
	return &RI_SDK_Connector_I2C_CloseAllReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Connector_I2C_Close - функция закрытия соединения
func (m *GRPCServer) RI_SDK_Connector_I2C_Close(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_CloseParams) (*RI_SDK_Connector_I2C_CloseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_Close(req.Descriptor_, uint8(req.Addr))
	return &RI_SDK_Connector_I2C_CloseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Connector_I2C_WriteBytes - функция для отправки байтов на I2C-устройства
func (m *GRPCServer) RI_SDK_Connector_I2C_WriteBytes(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_WriteBytesParams) (*RI_SDK_Connector_I2C_WriteBytesReturn, error) {
	wroteBytesLen, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_WriteBytes(req.Descriptor_, uint8(req.Addr), req.Buf, int(req.Len))
	return &RI_SDK_Connector_I2C_WriteBytesReturn{
		ErrorText:     errorText,
		ErrorCode:     errorCode,
		WroteBytesLen: wroteBytesLen,
	}, err
}

// RI_SDK_Connector_I2C_ReadBytes - читает байты с I2C-устройства
func (m *GRPCServer) RI_SDK_Connector_I2C_ReadBytes(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_ReadBytesParams) (*RI_SDK_Connector_I2C_ReadBytesReturn, error) {
	buf, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_ReadBytes(req.Descriptor_, uint8(req.Addr), req.ReadBytesLen)
	return &RI_SDK_Connector_I2C_ReadBytesReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
		Buf:       buf,
	}, err
}

// RI_SDK_Connector_I2C_WriteByte -отправляет байт на I2C-устройство
func (m *GRPCServer) RI_SDK_Connector_I2C_WriteByte(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_WriteByteParams) (*RI_SDK_Connector_I2C_WriteByteReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_WriteByte(req.Descriptor_, uint8(req.Addr), req.Value[0])
	return &RI_SDK_Connector_I2C_WriteByteReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Connector_I2C_ReadByte - читает байт с I2C-устройства
func (m *GRPCServer) RI_SDK_Connector_I2C_ReadByte(
	ctx context.Context,
	req *RI_SDK_Connector_I2C_ReadByteParams) (*RI_SDK_Connector_I2C_ReadByteReturn, error) {
	readedByte, errorText, errorCode, err := m.Impl.RI_SDK_Connector_I2C_ReadByte(req.Descriptor_, uint8(req.Addr))
	return &RI_SDK_Connector_I2C_ReadByteReturn{
		Value:     []byte{readedByte},
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
