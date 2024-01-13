export class Jwt {
  token: string = '';
  expiration: Date = new Date();
  refreshToken: RefreshToken = new RefreshToken();

  static From(cpy: Jwt): Jwt {
    if (!cpy) {
      return new Jwt();
    }
    const instance = new Jwt();
    Object.assign(instance, cpy);
    instance.expiration = new Date(cpy.expiration);
    cpy.refreshToken = RefreshToken.From(cpy.refreshToken);
    return instance;
  }
}

export class RefreshToken {
  token: string = '';
  expiration: Date = new Date();

  static From(cpy: RefreshToken): RefreshToken {
    if (!cpy) {
      return new RefreshToken();
    }
    const instance = new RefreshToken();
    Object.assign(instance, cpy);
    instance.expiration = new Date(cpy.expiration);
    return instance;
  }
}