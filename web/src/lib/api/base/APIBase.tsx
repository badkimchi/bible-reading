import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig} from "axios";
import {loginInfoStore} from "../../stores/loginInfoStore.ts";

export interface ResponseObject<T = any> {
    message: string;
    data: T;
}

export interface BaseResponse<T = any> extends AxiosResponse {
    data: ResponseObject<T>;
}

export class APIBase {
    private api: AxiosInstance;
    private config: AxiosRequestConfig;

    public constructor(config: AxiosRequestConfig) {
        this.api = axios.create(config);
        this.config = config;

        //Middleware run before request is sent.
        this.api.interceptors.request.use((param: InternalAxiosRequestConfig) => {
            // jwt
            param.headers['Authorization'] = "Bearer " + loginInfoStore.getState().current?.jwt.token;
            return param
        });

        //Middleware run before response is returned.
        this.api.interceptors.response.use((param: AxiosResponse) => {
            return param
        }, this.globalErrorHandler);

    }

    public globalErrorHandler(error: any): Promise<Error> {
        if (!error.response) {
            return Promise.reject(error.message);
        }
        if (error.response?.status === 401 &&
            typeof error.response?.data?.data === 'string' &&
            (error.response?.data?.data?.indexOf("token is expired") > -1 ||
                error.response?.data?.data?.indexOf("token is unauthorized") > -1)) {
            error.response.data = {message: error.response?.data?.data}
            window.location.href = "/signin";
        }

        return Promise.reject(error.response?.data?.message);
    }

    public getUri(config?: AxiosRequestConfig): string {
        return this.api.getUri(config);
    }

    public request<T, R = AxiosResponse<T>>(config: AxiosRequestConfig): Promise<R> {
        return this.api.request(config);
    }

    public get<T, R = BaseResponse<T>>(url: string, config?: AxiosRequestConfig): Promise<R> {
        return this.api.get(url, config);
    }

    public delete<T, R = BaseResponse<T>>(url: string, config?: AxiosRequestConfig): Promise<R> {
        return this.api.delete(url, config);
    }

    public post<T, R = BaseResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R> {
        return this.api.post(url, data, config);
    }

    public put<T, R = BaseResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R> {
        return this.api.put(url, data, config);
    }

    public postFile<T, R = BaseResponse<T>>(url: string, data?: any, eventListener?: any): Promise<R> {
        this.config.onUploadProgress = eventListener;
        return this.api.post(url, data, this.config);
    }

    public patch<T, R = BaseResponse<T>>(url: string, data?: string, config?: AxiosRequestConfig): Promise<R> {
        return this.api.patch(url, data, config);
    }
}