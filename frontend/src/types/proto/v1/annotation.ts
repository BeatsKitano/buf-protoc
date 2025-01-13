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
   * Rate limit per second
   * 每秒的速率限制
   */
  rateLimitPerSecond: number;
  /**
   * Timeout in milliseconds
   * 超时控制（毫秒）
   */
  timeout: number;
  /**
   * Concurrent
   * 并发控制
   */
  concurrent: number;
}

function createBaseMethodExtend(): MethodExtend {
  return {
    methodSignature: "",
    allowWithoutCredential: false,
    permission: "",
    authMethod: 0,
    audit: false,
    rateLimitPerSecond: 0,
    timeout: 0,
    concurrent: 0,
  };
}

export const MethodExtend = {
  encode(message: MethodExtend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.methodSignature !== "") {
      writer.uint32(800010).string(message.methodSignature);
    }
    if (message.allowWithoutCredential === true) {
      writer.uint32(800016).bool(message.allowWithoutCredential);
    }
    if (message.permission !== "") {
      writer.uint32(800026).string(message.permission);
    }
    if (message.authMethod !== 0) {
      writer.uint32(800032).int32(message.authMethod);
    }
    if (message.audit === true) {
      writer.uint32(800040).bool(message.audit);
    }
    if (message.rateLimitPerSecond !== 0) {
      writer.uint32(800048).int32(message.rateLimitPerSecond);
    }
    if (message.timeout !== 0) {
      writer.uint32(800056).int32(message.timeout);
    }
    if (message.concurrent !== 0) {
      writer.uint32(800064).int32(message.concurrent);
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
        case 100001:
          if (tag !== 800010) {
            break;
          }

          message.methodSignature = reader.string();
          continue;
        case 100002:
          if (tag !== 800016) {
            break;
          }

          message.allowWithoutCredential = reader.bool();
          continue;
        case 100003:
          if (tag !== 800026) {
            break;
          }

          message.permission = reader.string();
          continue;
        case 100004:
          if (tag !== 800032) {
            break;
          }

          message.authMethod = reader.int32() as any;
          continue;
        case 100005:
          if (tag !== 800040) {
            break;
          }

          message.audit = reader.bool();
          continue;
        case 100006:
          if (tag !== 800048) {
            break;
          }

          message.rateLimitPerSecond = reader.int32();
          continue;
        case 100007:
          if (tag !== 800056) {
            break;
          }

          message.timeout = reader.int32();
          continue;
        case 100008:
          if (tag !== 800064) {
            break;
          }

          message.concurrent = reader.int32();
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
      rateLimitPerSecond: isSet(object.rateLimitPerSecond) ? globalThis.Number(object.rateLimitPerSecond) : 0,
      timeout: isSet(object.timeout) ? globalThis.Number(object.timeout) : 0,
      concurrent: isSet(object.concurrent) ? globalThis.Number(object.concurrent) : 0,
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
    if (message.rateLimitPerSecond !== 0) {
      obj.rateLimitPerSecond = Math.round(message.rateLimitPerSecond);
    }
    if (message.timeout !== 0) {
      obj.timeout = Math.round(message.timeout);
    }
    if (message.concurrent !== 0) {
      obj.concurrent = Math.round(message.concurrent);
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
    message.rateLimitPerSecond = object.rateLimitPerSecond ?? 0;
    message.timeout = object.timeout ?? 0;
    message.concurrent = object.concurrent ?? 0;
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
