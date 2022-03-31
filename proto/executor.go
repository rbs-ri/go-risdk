package proto

import "context"

// RI_SDK_executor_Extend - расширяет базовый компонент до компонента группы исполнителей
func (m *GRPCClient) RI_SDK_Executor_Extend(basicDescriptor int64) (desrciptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Executor_Extend(context.Background(), &RI_SDK_Executor_ExtendParams{
		Basic: basicDescriptor,
	})
	if err != nil {
		return
	}
	return resp.Desrciptor, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_executor_State - возвращает состояние исполнителя
func (m *GRPCClient) RI_SDK_Executor_State(desrciptor int64) (state int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Executor_State(context.Background(), &RI_SDK_Executor_StateParams{
		Desrciptor: desrciptor,
	})
	if err != nil {
		return
	}
	return resp.State, resp.ErrorText, resp.ErrorCode, nil

}

// RI_SDK_Connector_I2C_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Executor_Extend(
	ctx context.Context,
	req *RI_SDK_Executor_ExtendParams) (*RI_SDK_Executor_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Executor_Extend(req.Basic)
	return &RI_SDK_Executor_ExtendReturn{
		Desrciptor: descriptor,
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}, err
}

// RI_SDK_executor_State - возвращает состояние исполнителя
func (m *GRPCServer) RI_SDK_Executor_State(
	ctx context.Context,
	req *RI_SDK_Executor_StateParams) (*RI_SDK_Executor_StateReturn, error) {
	state, errorText, errorCode, err := m.Impl.RI_SDK_Executor_State(req.Desrciptor)
	return &RI_SDK_Executor_StateReturn{
		State:     state,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
