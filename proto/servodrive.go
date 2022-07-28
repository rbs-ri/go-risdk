package proto

import "context"

// RI_SDK_Exec_ServoDrive_Extend -  расширяет исполнитель до сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_Extend(context.Background(), &RI_SDK_Exec_ServoDrive_ExtendParams{
		Exec: exec,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_ExtendToModel - расширяет сервопривод до модели
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_ExtendToModel(context.Background(), &RI_SDK_Exec_ServoDrive_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_CustomDeviceInit - Инициализация кастомного сервопривода - Инициализация кастомного сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_CustomDeviceInit(descriptor, maxImpulse, minImpulse, maxSpeed, rangeAngle int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_CustomDeviceInit(context.Background(), &RI_SDK_Exec_ServoDrive_CustomDeviceInitParams{
		Desrciptor: descriptor,
		MaxImpulse: maxImpulse,
		MinImpulse: minImpulse,
		MaxSpeed:   maxSpeed,
		RangeAngle: rangeAngle,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_TurnByDutyCycle - Абсолютный поворот. Угол задается через кол-во шагов сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_TurnByDutyCycle(descriptor, steps int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_TurnByDutyCycle(context.Background(), &RI_SDK_Exec_ServoDrive_TurnByDutyCycleParams{
		Desrciptor: descriptor,
		Steps:      steps,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_TurnByPulse - Абсолютный поворот. Угол задается через значение импульса
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_TurnByPulse(descriptor, pulse int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_TurnByPulse(context.Background(), &RI_SDK_Exec_ServoDrive_TurnByPulseParams{
		Desrciptor: descriptor,
		Pulse:      pulse,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_GetCurrentAngle - Возвращает текущий угол сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_GetCurrentAngle(descriptor int64) (angle int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_GetCurrentAngle(context.Background(), &RI_SDK_Exec_ServoDrive_GetCurrentAngleParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Angle, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_GetState - Получение состояния сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_GetState(descriptor int64) (state int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_GetState(context.Background(), &RI_SDK_Exec_ServoDrive_GetStateParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.State, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange - Устанавливает сервопривод в середину рабочего диапазона
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange(context.Background(), &RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRangeParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_MinStepRotate - Выполняет поворот сервопривода на минимальный шаг
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_MinStepRotate(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_MinStepRotate(context.Background(), &RI_SDK_Exec_ServoDrive_MinStepRotateParams{
		Descriptor_: descriptor,
		Direction:   direction,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_Turn - Выполняет поворот сервопривода на заданный угол с заданной угловой скоростью. Если угол положительный, поворот осуществляется по часовой стрелке, если отрицательный, то против часовой.
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_Turn(descriptor, angle, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_Turn(context.Background(), &RI_SDK_Exec_ServoDrive_TurnParams{
		Descriptor_: descriptor,
		Angle:       angle,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_Stop - Прекращает подачу сигнала к сервоприводу
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_Stop(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_Stop(context.Background(), &RI_SDK_Exec_ServoDrive_StopParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_exec_ServoDrive_Rotate - Выполняет вращение сервопривода с заданной угловой скоростью до
// тех пор, пока не будет достигнут максимальный или минимальный
// угол сервопривода.
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_Rotate(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_Rotate(context.Background(), &RI_SDK_Exec_ServoDrive_RotateParams{
		Descriptor_: descriptor,
		Direction:   direction,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed - Выполняет вращение сервопривода с заданным
// процентом от максимальной скорости до тех пор, пока не будет достигнут
// максимальный или минимальный угол сервопривода
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed(context.Background(), &RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeedParams{
		Descriptor_: descriptor,
		Direction:   direction,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed - Выполняет поворот сервопривода на заданный угол с заданным процентом
// от максимальной скорости. Если угол положительный, поворот
// осуществляется по часовой стрелке, если отрицательный, то против часовой.
func (m *GRPCClient) RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed(descriptor, angle, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed(context.Background(), &RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeedParams{
		Descriptor_: descriptor,
		Angle:       angle,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_I2C_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_Extend(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_ExtendParams) (*RI_SDK_Exec_ServoDrive_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_Extend(req.Exec)
	return &RI_SDK_Exec_ServoDrive_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_ExtendToModel - расширяет сервопривод до модели
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_ExtendToModelParams) (*RI_SDK_Exec_ServoDrive_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Exec_ServoDrive_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_CustomDeviceInit - Инициализация кастомного сервопривода - Инициализация кастомного сервопривода
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_CustomDeviceInit(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_CustomDeviceInitParams) (*RI_SDK_Exec_ServoDrive_CustomDeviceInitReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_CustomDeviceInit(req.Desrciptor, req.MaxImpulse, req.MinImpulse, req.MaxSpeed, req.RangeAngle)
	return &RI_SDK_Exec_ServoDrive_CustomDeviceInitReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_TurnByDutyCycle - Абсолютный поворот. Угол задается через кол-во шагов сервопривода
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_TurnByDutyCycle(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_TurnByDutyCycleParams) (*RI_SDK_Exec_ServoDrive_TurnByDutyCycleReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_TurnByDutyCycle(req.Desrciptor, req.Steps)
	return &RI_SDK_Exec_ServoDrive_TurnByDutyCycleReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_TurnByPulse - Абсолютный поворот. Угол задается через значение импульса
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_TurnByPulse(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_TurnByPulseParams) (*RI_SDK_Exec_ServoDrive_TurnByPulseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_TurnByPulse(req.Desrciptor, req.Pulse)
	return &RI_SDK_Exec_ServoDrive_TurnByPulseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_GetCurrentAngle - Возвращает текущий угол сервопривода
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_GetCurrentAngle(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_GetCurrentAngleParams) (*RI_SDK_Exec_ServoDrive_GetCurrentAngleReturn, error) {
	angle, errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_GetCurrentAngle(req.Desrciptor)
	return &RI_SDK_Exec_ServoDrive_GetCurrentAngleReturn{
		Angle:     angle,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_GetState - Получение состояния сервопривода
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_GetState(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_GetStateParams) (*RI_SDK_Exec_ServoDrive_GetStateReturn, error) {
	state, errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_GetState(req.Desrciptor)
	return &RI_SDK_Exec_ServoDrive_GetStateReturn{
		State:     state,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange - Устанавливает сервопривод в середину рабочего диапазона
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRangeParams) (*RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRangeReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRange(req.Descriptor_)
	return &RI_SDK_Exec_ServoDrive_SetPositionToMidWorkingRangeReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_MinStepRotate - Выполняет поворот сервопривода на минимальный шаг
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_MinStepRotate(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_MinStepRotateParams) (*RI_SDK_Exec_ServoDrive_MinStepRotateReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_MinStepRotate(req.Descriptor_, req.Direction, req.Speed, req.Async)
	return &RI_SDK_Exec_ServoDrive_MinStepRotateReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_Turn - Выполняет поворот сервопривода на заданный угол с заданной угловой скоростью. Если угол положительный, поворот осуществляется по часовой стрелке, если отрицательный, то против часовой.
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_Turn(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_TurnParams) (*RI_SDK_Exec_ServoDrive_TurnReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_Turn(req.Descriptor_, req.Angle, req.Speed, req.Async)
	return &RI_SDK_Exec_ServoDrive_TurnReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_Stop - Прекращает подачу сигнала к сервоприводу
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_Stop(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_StopParams) (*RI_SDK_Exec_ServoDrive_StopReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_Stop(req.Descriptor_)
	return &RI_SDK_Exec_ServoDrive_StopReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_Rotate - Выполняет вращение сервопривода с заданной угловой скоростью до
// тех пор, пока не будет достигнут максимальный или минимальный
// угол сервопривода.
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_Rotate(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_RotateParams) (*RI_SDK_Exec_ServoDrive_RotateReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_Rotate(req.Descriptor_, req.Direction, req.Speed, req.Async)
	return &RI_SDK_Exec_ServoDrive_RotateReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed - Выполняет вращение сервопривода с заданным
// процентом от максимальной скорости до тех пор, пока не будет достигнут
// максимальный или минимальный угол сервопривода
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeedParams) (*RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeedReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed(req.Descriptor_, req.Direction, req.Speed, req.Async)
	return &RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeedReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_exec_ServoDrive_TurnWithRelativeSpeed - Выполняет поворот сервопривода на заданный угол с заданным процентом
// от максимальной скорости. Если угол положительный, поворот
// осуществляется по часовой стрелке, если отрицательный, то против часовой.
func (m *GRPCServer) RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed(
	ctx context.Context,
	req *RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeedParams) (*RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeedReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed(req.Descriptor_, req.Angle, req.Speed, req.Async)
	return &RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeedReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
