import {APIBase} from './base/APIBase.tsx';
import {APIConfig} from './base/conf.tsx';
import {AxiosRequestConfig} from "axios";

export class API extends APIBase {

    constructor(base: AxiosRequestConfig) {
        super(base);
    }

    public postAudio(data): Promise<string> {
        return this.postFile<string>('/audio', data)
            .then((response) => {
                const {data} = response;
                return data?.data;
            })
            .catch((error) => {
                throw error;
            });
    }
}
const api = new API(APIConfig)

export {api as APIAudio};

