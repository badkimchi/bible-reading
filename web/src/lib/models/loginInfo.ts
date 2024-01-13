import {Jwt} from "./jwt.ts";

export class LoginInfo {
    id: string = '';
    username: string = '';
    email: string = '';
    image: string = '';
    jwt: Jwt = new Jwt();

    static From(cpy: LoginInfo): LoginInfo {
        const instance = new LoginInfo();
        Object.assign(instance, cpy);
        instance.jwt = Jwt.From(cpy.jwt);
        return instance
    }
}
