package proto

import "context"

// RI_SDK_Sigmod_PWM_Extend - расширяет компонент группы
func (m *GRPCClient) RI_SDK_Sigmod_PWM_Extend(con int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_Extend(context.Background(), &RI_SDK_Sigmod_PWM_ExtendParams{
		Con: con,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_ExtendToModel - расширяет pwm до модели
func (m *GRPCClient) RI_SDK_Sigmod_PWM_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_ExtendToModel(context.Background(), &RI_SDK_Sigmod_PWM_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_GetResolution - функция получения разрешения ШИМ
func (m *GRPCClient) RI_SDK_Sigmod_PWM_GetResolution(descriptor int64) (resolution int, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_GetResolution(context.Background(), &RI_SDK_Sigmod_PWM_GetResolutionParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return int(resp.Resolution), resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_GetFreq - функция получения частоты ШИМ
func (m *GRPCClient) RI_SDK_Sigmod_PWM_GetFreq(descriptor int64) (freq int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_GetFreq(context.Background(), &RI_SDK_Sigmod_PWM_GetFreqParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Freq, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_SetFreq - функция установки частоты ШИМ
func (m *GRPCClient) RI_SDK_Sigmod_PWM_SetFreq(descriptor int64, freq int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_SetFreq(context.Background(), &RI_SDK_Sigmod_PWM_SetFreqParams{
		Descriptor_: descriptor,
		Freq:        freq,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_WriteRegBytes - Запись в указанный регистр n-кол-во байт
func (m *GRPCClient) RI_SDK_Sigmod_PWM_WriteRegBytes(descriptor int64, reg byte, buf []byte) (wroteBytesLen int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_WriteRegBytes(context.Background(), &RI_SDK_Sigmod_PWM_WriteRegBytesParams{
		Descriptor_: descriptor,
		Reg:         []byte{reg},
		Buf:         buf,
	})
	if err != nil {
		return
	}
	return resp.WroteBytesLen, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_ReadRegBytes - Читает с указанного регистра n-кол-во байт
func (m *GRPCClient) RI_SDK_Sigmod_PWM_ReadRegBytes(descriptor int64, reg byte, buf []byte) (readedBytesLen int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_ReadRegBytes(context.Background(), &RI_SDK_Sigmod_PWM_ReadRegBytesParams{
		Descriptor_: descriptor,
		Reg:         []byte{reg},
		Buf:         buf,
	})
	if err != nil {
		return
	}
	return resp.ReadedBytesLen, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_WriteByte - Пишет байт в заданный регистр
func (m *GRPCClient) RI_SDK_Sigmod_PWM_WriteByte(descriptor int64, reg byte, value byte) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_WriteByte(context.Background(), &RI_SDK_Sigmod_PWM_WriteRegByteParams{
		Descriptor_: descriptor,
		Reg:         []byte{reg},
		Value:       []byte{value},
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_ReadByte - Читает байт из заданного регистра
func (m *GRPCClient) RI_SDK_Sigmod_PWM_ReadByte(descriptor int64, reg byte) (value byte, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_ReadByte(context.Background(), &RI_SDK_Sigmod_PWM_ReadRegByteParams{
		Descriptor_: descriptor,
		Reg:         []byte{reg},
	})
	if err != nil {
		return
	}
	return resp.Value[0], resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Sigmod_PWM_Extend(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ExtendParams) (*RI_SDK_Sigmod_PWM_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_Extend(req.Con)
	return &RI_SDK_Sigmod_PWM_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_ExtendToModel - расширяет pwm до модели
func (m *GRPCServer) RI_SDK_Sigmod_PWM_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ExtendToModelParams) (*RI_SDK_Sigmod_PWM_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Sigmod_PWM_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_GetPortFreq - функция получения частоты порта
func (m *GRPCClient) RI_SDK_Sigmod_PWM_GetPortFreq(descriptor int64, port int64) (freq int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_GetPortFreq(context.Background(), &RI_SDK_Sigmod_PWM_GetPortFreqParams{
		Descriptor_: descriptor,
		Port:        port,
	})
	if err != nil {
		return
	}
	return resp.Freq, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_SetPortFreq - функция установки частоты порта
func (m *GRPCClient) RI_SDK_Sigmod_PWM_SetPortFreq(descriptor int64, port int64, freq int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_SetPortFreq(context.Background(), &RI_SDK_Sigmod_PWM_SetPortFreqParams{
		Descriptor_: descriptor,
		Port:        port,
		Freq:        freq,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_ResetAll - Сброс значения всех портов
func (m *GRPCClient) RI_SDK_Sigmod_PWM_ResetAll(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_ResetAll(context.Background(), &RI_SDK_Sigmod_PWM_ResetAllParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_ResetPort - Сброс значения порта
func (m *GRPCClient) RI_SDK_Sigmod_PWM_ResetPort(descriptor int64, port int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_ResetPort(context.Background(), &RI_SDK_Sigmod_PWM_ResetPortParams{
		Descriptor_: descriptor,
		Port:        port,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_SetPortDutyCycle - функция установки значения порта
func (m *GRPCClient) RI_SDK_Sigmod_PWM_SetPortDutyCycle(descriptor int64, port int64, on int64, off int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_SetPortDutyCycle(context.Background(), &RI_SDK_Sigmod_PWM_SetPortDutyCycleParams{
		Descriptor_: descriptor,
		Port:        port,
		On:          on,
		Off:         off,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_GetPortDutyCycle - функция получения значения с порта
func (m *GRPCClient) RI_SDK_Sigmod_PWM_GetPortDutyCycle(descriptor int64, port int64) (on int64, off int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_GetPortDutyCycle(context.Background(), &RI_SDK_Sigmod_PWM_GetPortDutyCycleParams{
		Descriptor_: descriptor,
		Port:        port,
	})
	if err != nil {
		return
	}
	return resp.On, resp.Off, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_Close - функция закрытия соединения с i2c шиной
func (m *GRPCClient) RI_SDK_Sigmod_PWM_Close(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sigmod_PWM_Close(context.Background(), &RI_SDK_Sigmod_PWM_CloseParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sigmod_PWM_GetResolution - функция получения разрешения ШИМ
func (m *GRPCServer) RI_SDK_Sigmod_PWM_GetResolution(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_GetResolutionParams) (*RI_SDK_Sigmod_PWM_GetResolutioneReturn, error) {
	resolution, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_GetResolution(req.Descriptor_)
	return &RI_SDK_Sigmod_PWM_GetResolutioneReturn{
		Resolution: int64(resolution),
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_GetFreq - функция получения частоты ШИМ
func (m *GRPCServer) RI_SDK_Sigmod_PWM_GetFreq(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_GetFreqParams) (*RI_SDK_Sigmod_PWM_GetFreqReturn, error) {
	freq, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_GetFreq(req.Descriptor_)
	return &RI_SDK_Sigmod_PWM_GetFreqReturn{
		Freq:      freq,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_SetFreq - функция установки частоты ШИМ
func (m *GRPCServer) RI_SDK_Sigmod_PWM_SetFreq(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_SetFreqParams) (*RI_SDK_Sigmod_PWM_SetFreqReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_SetFreq(req.Descriptor_, req.Freq)
	return &RI_SDK_Sigmod_PWM_SetFreqReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_WriteRegBytes - Запись в указанный регистр n-кол-во байт
func (m *GRPCServer) RI_SDK_Sigmod_PWM_WriteRegBytes(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_WriteRegBytesParams) (*RI_SDK_Sigmod_PWM_WriteRegBytesReturn, error) {
	wroteBytesLen, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_WriteRegBytes(req.Descriptor_, req.Reg[0], req.Buf)
	return &RI_SDK_Sigmod_PWM_WriteRegBytesReturn{
		WroteBytesLen: wroteBytesLen,
		ErrorText:     errorText,
		ErrorCode:     errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_ReadRegBytes - Читает с указанного регистра n-кол-во байт
func (m *GRPCServer) RI_SDK_Sigmod_PWM_ReadRegBytes(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ReadRegBytesParams) (*RI_SDK_Sigmod_PWM_ReadRegBytesReturn, error) {
	readedBytesLen, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_ReadRegBytes(req.Descriptor_, req.Reg[0], req.Buf)
	return &RI_SDK_Sigmod_PWM_ReadRegBytesReturn{
		ReadedBytesLen: readedBytesLen,
		ErrorText:      errorText,
		ErrorCode:      errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_WriteByte - Пишет байт в заданный регистр
func (m *GRPCServer) RI_SDK_Sigmod_PWM_WriteByte(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_WriteRegByteParams) (*RI_SDK_Sigmod_PWM_WriteRegByteReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_WriteByte(req.Descriptor_, req.Reg[0], req.Value[0])
	return &RI_SDK_Sigmod_PWM_WriteRegByteReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_ReadByte - Читает байт из заданного регистра
func (m *GRPCServer) RI_SDK_Sigmod_PWM_ReadByte(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ReadRegByteParams) (*RI_SDK_Sigmod_PWM_ReadRegByteReturn, error) {
	value, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_ReadByte(req.Descriptor_, req.Reg[0])
	return &RI_SDK_Sigmod_PWM_ReadRegByteReturn{
		Value:     []byte{value},
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_GetPortFreq - функция получения частоты порта
func (m *GRPCServer) RI_SDK_Sigmod_PWM_GetPortFreq(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_GetPortFreqParams) (*RI_SDK_Sigmod_PWM_GetPortFreqReturn, error) {
	freq, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_GetPortFreq(req.Descriptor_, req.Port)
	return &RI_SDK_Sigmod_PWM_GetPortFreqReturn{
		Freq:      freq,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_SetPortFreq - функция установки частоты порта
func (m *GRPCServer) RI_SDK_Sigmod_PWM_SetPortFreq(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_SetPortFreqParams) (*RI_SDK_Sigmod_PWM_SetPortFreqReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_SetPortFreq(req.Descriptor_, req.Port, req.Freq)
	return &RI_SDK_Sigmod_PWM_SetPortFreqReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_ResetAll - Сброс значения всех портов
func (m *GRPCServer) RI_SDK_Sigmod_PWM_ResetAll(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ResetAllParams) (*RI_SDK_Sigmod_PWM_ResetAllReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_ResetAll(req.Descriptor_)
	return &RI_SDK_Sigmod_PWM_ResetAllReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_ResetPort - Сброс значения порта
func (m *GRPCServer) RI_SDK_Sigmod_PWM_ResetPort(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_ResetPortParams) (*RI_SDK_Sigmod_PWM_ResetPortReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_ResetPort(req.Descriptor_, req.Port)
	return &RI_SDK_Sigmod_PWM_ResetPortReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_SetPortDutyCycle - функция установки значения порта
func (m *GRPCServer) RI_SDK_Sigmod_PWM_SetPortDutyCycle(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_SetPortDutyCycleParams) (*RI_SDK_Sigmod_PWM_SetPortDutyCycleReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_SetPortDutyCycle(req.Descriptor_, req.Port, req.On, req.Off)
	return &RI_SDK_Sigmod_PWM_SetPortDutyCycleReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_GetPortDutyCycle - функция получения значения с порта
func (m *GRPCServer) RI_SDK_Sigmod_PWM_GetPortDutyCycle(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_GetPortDutyCycleParams) (*RI_SDK_Sigmod_PWM_GetPortDutyCycleReturn, error) {
	on, off, errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_GetPortDutyCycle(req.Descriptor_, req.Port)
	return &RI_SDK_Sigmod_PWM_GetPortDutyCycleReturn{
		On:        on,
		Off:       off,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sigmod_PWM_Close - функция закрытия соединения с i2c шиной
func (m *GRPCServer) RI_SDK_Sigmod_PWM_Close(
	ctx context.Context,
	req *RI_SDK_Sigmod_PWM_CloseParams) (*RI_SDK_Sigmod_PWM_CloseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sigmod_PWM_Close(req.Descriptor_)
	return &RI_SDK_Sigmod_PWM_CloseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
