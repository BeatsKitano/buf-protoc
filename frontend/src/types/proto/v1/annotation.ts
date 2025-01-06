/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "api.v1";

export enum AuthMethod {
  AUTH_METHOD_UNSPECIFIED = 0,
  /** IAM - IAM uses the standard IAM authorization check on the organizational resources. */
  IAM = 1,
  /** CUSTOM - Custom authorization method. */
  CUSTOM = 2,
  UNRECOGNIZED = -1,
}

export function authMethodFromJSON(object: any): AuthMethod {
  switch (object) {
    case 0:
    case "AUTH_METHOD_UNSPECIFIED":
      return AuthMethod.AUTH_METHOD_UNSPECIFIED;
    case 1:
    case "IAM":
      return AuthMethod.IAM;
    case 2:
    case "CUSTOM":
      return AuthMethod.CUSTOM;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AuthMethod.UNRECOGNIZED;
  }
}

export function authMethodToJSON(object: AuthMethod): string {
  switch (object) {
    case AuthMethod.AUTH_METHOD_UNSPECIFIED:
      return "AUTH_METHOD_UNSPECIFIED";
    case AuthMethod.IAM:
      return "IAM";
    case AuthMethod.CUSTOM:
      return "CUSTOM";
    case AuthMethod.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ServiceExtend {
  serviceSignature: string;
  /**
   * Whether to allow access without credentials
   * 是否允许无凭证访问
   */
  allowWithoutCredential: boolean;
  /**
   * Permission required to access the method
   * 访问该方法所需的权限
   */
  permission: string;
  /**
   * Authorization method
   * 授权方法
   */
  authMethod: AuthMethod;
  /**
   * Whether to audit the method
   * 是否审计该方法
   */
  audit: boolean;
  /**
   * Rate limit per minute
   * 每分钟的速率限制
   */
  rateLimit: number;
  /**
   * Timeout in milliseconds
   * 超时控制（毫秒）
   */
  timeout: number;
}

export interface MethodExtend {
  /**
   * Method signature
   * 方法签名
   */
  methodSignature: string;
  /**
   * Whether to allow access without credentials
   * 是否允许无凭证访问
   */
  allowWithoutCredential: boolean;
  /**
   * Permission required to access the method
   * 访问该方法所需的权限
   */
  permission: string;
  /**
   * Authorization method
   * 授权方法
   */
  authMethod: AuthMethod;
  /**
   * Whether to audit the method
   * 是否审计该方法
   */
  audit: boolean;
  /**
   * Rate limit per minute
   * 每分钟的速率限制
   */
  rateLimit: number;
  /**
   * Timeout in milliseconds
   * 超时控制（毫秒）
   */
  timeout: number;
}

function createBaseServiceExtend(): ServiceExtend {
  return {
    serviceSignature: "",
    allowWithoutCredential: false,
    permission: "",
    authMethod: 0,
    audit: false,
    rateLimit: 0,
    timeout: 0,
  };
}

export const ServiceExtend = {
  encode(message: ServiceExtend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.serviceSignature !== "") {
      writer.uint32(10).string(message.serviceSignature);
    }
    if (message.allowWithoutCredential === true) {
      writer.uint32(16).bool(message.allowWithoutCredential);
    }
    if (message.permission !== "") {
      writer.uint32(26).string(message.permission);
    }
    if (message.authMethod !== 0) {
      writer.uint32(32).int32(message.authMethod);
    }
    if (message.audit === true) {
      writer.uint32(40).bool(message.audit);
    }
    if (message.rateLimit !== 0) {
      writer.uint32(80).int32(message.rateLimit);
    }
    if (message.timeout !== 0) {
      writer.uint32(88).int32(message.timeout);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ServiceExtend {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceExtend();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.serviceSignature = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.allowWithoutCredential = reader.bool();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.permission = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.authMethod = reader.int32() as any;
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.audit = reader.bool();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.rateLimit = reader.int32();
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.timeout = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ServiceExtend {
    return {
      serviceSignature: isSet(object.serviceSignature) ? globalThis.String(object.serviceSignature) : "",
      allowWithoutCredential: isSet(object.allowWithoutCredential)
        ? globalThis.Boolean(object.allowWithoutCredential)
        : false,
      permission: isSet(object.permission) ? globalThis.String(object.permission) : "",
      authMethod: isSet(object.authMethod) ? authMethodFromJSON(object.authMethod) : 0,
      audit: isSet(object.audit) ? globalThis.Boolean(object.audit) : false,
      rateLimit: isSet(object.rateLimit) ? globalThis.Number(object.rateLimit) : 0,
      timeout: isSet(object.timeout) ? globalThis.Number(object.timeout) : 0,
    };
  },

  toJSON(message: ServiceExtend): unknown {
    const obj: any = {};
    if (message.serviceSignature !== "") {
      obj.serviceSignature = message.serviceSignature;
    }
    if (message.allowWithoutCredential === true) {
      obj.allowWithoutCredential = message.allowWithoutCredential;
    }
    if (message.permission !== "") {
      obj.permission = message.permission;
    }
    if (message.authMethod !== 0) {
      obj.authMethod = authMethodToJSON(message.authMethod);
    }
    if (message.audit === true) {
      obj.audit = message.audit;
    }
    if (message.rateLimit !== 0) {
      obj.rateLimit = Math.round(message.rateLimit);
    }
    if (message.timeout !== 0) {
      obj.timeout = Math.round(message.timeout);
    }
    return obj;
  },

  create(base?: DeepPartial<ServiceExtend>): ServiceExtend {
    return ServiceExtend.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ServiceExtend>): ServiceExtend {
    const message = createBaseServiceExtend();
    message.serviceSignature = object.serviceSignature ?? "";
    message.allowWithoutCredential = object.allowWithoutCredential ?? false;
    message.permission = object.permission ?? "";
    message.authMethod = object.authMethod ?? 0;
    message.audit = object.audit ?? false;
    message.rateLimit = object.rateLimit ?? 0;
    message.timeout = object.timeout ?? 0;
    return message;
  },
};

function createBaseMethodExtend(): MethodExtend {
  return {
    methodSignature: "",
    allowWithoutCredential: false,
    permission: "",
    authMethod: 0,
    audit: false,
    rateLimit: 0,
    timeout: 0,
  };
}

export const MethodExtend = {
  encode(message: MethodExtend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.methodSignature !== "") {
      writer.uint32(10).string(message.methodSignature);
    }
    if (message.allowWithoutCredential === true) {
      writer.uint32(16).bool(message.allowWithoutCredential);
    }
    if (message.permission !== "") {
      writer.uint32(26).string(message.permission);
    }
    if (message.authMethod !== 0) {
      writer.uint32(32).int32(message.authMethod);
    }
    if (message.audit === true) {
      writer.uint32(40).bool(message.audit);
    }
    if (message.rateLimit !== 0) {
      writer.uint32(80).int32(message.rateLimit);
    }
    if (message.timeout !== 0) {
      writer.uint32(88).int32(message.timeout);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MethodExtend {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMethodExtend();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.methodSignature = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.allowWithoutCredential = reader.bool();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.permission = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.authMethod = reader.int32() as any;
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.audit = reader.bool();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.rateLimit = reader.int32();
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.timeout = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MethodExtend {
    return {
      methodSignature: isSet(object.methodSignature) ? globalThis.String(object.methodSignature) : "",
      allowWithoutCredential: isSet(object.allowWithoutCredential)
        ? globalThis.Boolean(object.allowWithoutCredential)
        : false,
      permission: isSet(object.permission) ? globalThis.String(object.permission) : "",
      authMethod: isSet(object.authMethod) ? authMethodFromJSON(object.authMethod) : 0,
      audit: isSet(object.audit) ? globalThis.Boolean(object.audit) : false,
      rateLimit: isSet(object.rateLimit) ? globalThis.Number(object.rateLimit) : 0,
      timeout: isSet(object.timeout) ? globalThis.Number(object.timeout) : 0,
    };
  },

  toJSON(message: MethodExtend): unknown {
    const obj: any = {};
    if (message.methodSignature !== "") {
      obj.methodSignature = message.methodSignature;
    }
    if (message.allowWithoutCredential === true) {
      obj.allowWithoutCredential = message.allowWithoutCredential;
    }
    if (message.permission !== "") {
      obj.permission = message.permission;
    }
    if (message.authMethod !== 0) {
      obj.authMethod = authMethodToJSON(message.authMethod);
    }
    if (message.audit === true) {
      obj.audit = message.audit;
    }
    if (message.rateLimit !== 0) {
      obj.rateLimit = Math.round(message.rateLimit);
    }
    if (message.timeout !== 0) {
      obj.timeout = Math.round(message.timeout);
    }
    return obj;
  },

  create(base?: DeepPartial<MethodExtend>): MethodExtend {
    return MethodExtend.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MethodExtend>): MethodExtend {
    const message = createBaseMethodExtend();
    message.methodSignature = object.methodSignature ?? "";
    message.allowWithoutCredential = object.allowWithoutCredential ?? false;
    message.permission = object.permission ?? "";
    message.authMethod = object.authMethod ?? 0;
    message.audit = object.audit ?? false;
    message.rateLimit = object.rateLimit ?? 0;
    message.timeout = object.timeout ?? 0;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
