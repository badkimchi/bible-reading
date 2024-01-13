// import { AxiosResponse } from 'axios';
// import { request } from '../setupAxios';
// import { ChangePasswordInput, LoginDTO, RegisterDTO, ResetPasswordInput } from '../dtos/AuthInput';
// import { LoginInfo } from '../../models/account';
//
// export const register = (body: RegisterDTO): Promise<AxiosResponse<LoginInfo>> => request.post('/account/register', body);
//
// export const login = (body: LoginDTO): Promise<AxiosResponse<LoginInfo>> => request.post('/account/login', body);
//
// export const logout = (): Promise<AxiosResponse> => request.post('/account/logout');
//
// export const forgotPassword = (email: string): Promise<AxiosResponse<boolean>> =>
//   request.post('/account/forgot-password', { email });
//
// export const changePassword = (body: ChangePasswordInput): Promise<AxiosResponse> =>
//   request.put('/account/change-password', body);
//
// export const resetPassword = (body: ResetPasswordInput): Promise<AxiosResponse<LoginInfo>> =>
//   request.post('/account/reset-password', body);
