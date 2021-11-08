package proto

import "context"

// RI_SDK_Exec_RGB_LED_Extend -  расширяет исполнитель до светодиода
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_Extend(context.Background(), &RI_SDK_Exec_RGB_LED_ExtendParams{
		Exec: exec,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_ExtendToModel - расширяет светодиод до модели
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_ExtendToModel(context.Background(), &RI_SDK_Exec_RGB_LED_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_SinglePulse - Одиночный световой импульса (непрерывное свечение)
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_SinglePulse(descriptor, r, g, b, duration int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_SinglePulse(context.Background(), &RI_SDK_Exec_RGB_LED_SinglePulseParams{
		Descriptor_: descriptor,
		R:           r,
		G:           g,
		B:           b,
		Duration:    duration,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_Stop - Прекращает подачу сигнала к светодиоду
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_Stop(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_Stop(context.Background(), &RI_SDK_Exec_RGB_LED_StopParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_GetState - Получение состояния светодиода
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_GetState(descriptor int64) (state int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_GetState(context.Background(), &RI_SDK_Exec_RGB_LED_GetStateParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.State, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_GetColor - Получение цвета светодиода
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_GetColor(descriptor int64) (r, g, b int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_GetColor(context.Background(), &RI_SDK_Exec_RGB_LED_GetColorParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.R, resp.B, resp.G, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_FlashingWithFrequency - Мигание с заданной частотой
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_FlashingWithFrequency(descriptor, r, g, b, frequency, qty int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_FlashingWithFrequency(context.Background(), &RI_SDK_Exec_RGB_LED_FlashingWithFrequencyParams{
		Descriptor_: descriptor,
		R:           r,
		G:           g,
		B:           b,
		Frequency:   frequency,
		Qty:         qty,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_FlashingWithPause - Мигание с заданной продолжительностью и паузой
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_FlashingWithPause(descriptor, r, g, b, duration, pause, qty int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_FlashingWithPause(context.Background(), &RI_SDK_Exec_RGB_LED_FlashingWithPauseParams{
		Descriptor_: descriptor,
		R:           r,
		G:           g,
		B:           b,
		Duration:    duration,
		Pause:       pause,
		Qty:         qty,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_Flicker - Мерцание
func (m *GRPCClient) RI_SDK_Exec_RGB_LED_Flicker(descriptor, r, g, b, duration, qty int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RGB_LED_Flicker(context.Background(), &RI_SDK_Exec_RGB_LED_FlickerParams{
		Descriptor_: descriptor,
		R:           r,
		G:           g,
		B:           b,
		Duration:    duration,
		Qty:         qty,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RGB_LED_Extend - расширяет исполнитель до светодиода
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_Extend(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_ExtendParams) (*RI_SDK_Exec_RGB_LED_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_Extend(req.Exec)
	return &RI_SDK_Exec_RGB_LED_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_ExtendToModel - расширяет светодиод до модели
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_ExtendToModelParams) (*RI_SDK_Exec_RGB_LED_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Exec_RGB_LED_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_SinglePulse - Одиночный световой импульса (непрерывное свечение)
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_SinglePulse(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_SinglePulseParams) (*RI_SDK_Exec_RGB_LED_SinglePulseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_SinglePulse(req.Descriptor_, req.R, req.G, req.B, req.Duration, req.Async)
	return &RI_SDK_Exec_RGB_LED_SinglePulseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_Stop - Прекращает подачу сигнала к светодиоду
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_Stop(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_StopParams) (*RI_SDK_Exec_RGB_LED_StopReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_Stop(req.Descriptor_)
	return &RI_SDK_Exec_RGB_LED_StopReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_GetState - Получение состояния светодиода
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_GetState(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_GetStateParams) (*RI_SDK_Exec_RGB_LED_GetStateReturn, error) {
	state, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_GetState(req.Descriptor_)
	return &RI_SDK_Exec_RGB_LED_GetStateReturn{
		State:     state,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_GetColor - Получение цвета светодиода
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_GetColor(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_GetColorParams) (*RI_SDK_Exec_RGB_LED_GetColorReturn, error) {
	r, g, b, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_GetColor(req.Descriptor_)
	return &RI_SDK_Exec_RGB_LED_GetColorReturn{
		R:         r,
		G:         g,
		B:         b,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_FlashingWithFrequency - Мигание с заданной частотой
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_FlashingWithFrequency(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_FlashingWithFrequencyParams) (*RI_SDK_Exec_RGB_LED_FlashingWithFrequencyReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_FlashingWithFrequency(req.Descriptor_, req.R, req.G, req.B, req.Frequency, req.Qty, req.Async)
	return &RI_SDK_Exec_RGB_LED_FlashingWithFrequencyReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_FlashingWithPause - Мигание с заданной продолжительностью и паузой
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_FlashingWithPause(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_FlashingWithPauseParams) (*RI_SDK_Exec_RGB_LED_FlashingWithPauseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_FlashingWithPause(req.Descriptor_, req.R, req.G, req.B, req.Duration, req.Pause, req.Qty, req.Async)
	return &RI_SDK_Exec_RGB_LED_FlashingWithPauseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RGB_LED_Flicker - Мерцание
func (m *GRPCServer) RI_SDK_Exec_RGB_LED_Flicker(
	ctx context.Context,
	req *RI_SDK_Exec_RGB_LED_FlickerParams) (*RI_SDK_Exec_RGB_LED_FlickerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RGB_LED_Flicker(req.Descriptor_, req.R, req.G, req.B, req.Duration, req.Qty, req.Async)
	return &RI_SDK_Exec_RGB_LED_FlickerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
