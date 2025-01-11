/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import { Expr } from "../google/type/expr";

export const protobufPackage = "api.v1";

/** 请求参数 */
export interface GetUserRequest {
  /**
   * 字段解释描述文本
   *
   * syntax. CEL is a C-like expression language. The syntax and semantics of CEL
   */
  age: number;
  /** 购买日期 */
  purchaseDate:
    | Date
    | undefined;
  /** 交付日期 */
  deliveryDate:
    | Date
    | undefined;
  /** 物品名 */
  name: string;
  /** 自定义表达式 */
  customExpr: Expr | undefined;
}

/** 用户信息 */
export interface User {
  /** 用户id */
  id: string;
  /** 用户名 */
  name: string;
  son: string;
}

function createBaseGetUserRequest(): GetUserRequest {
  return { age: 0, purchaseDate: undefined, deliveryDate: undefined, name: "", customExpr: undefined };
}

export const GetUserRequest = {
  encode(message: GetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.age !== 0) {
      writer.uint32(24).int32(message.age);
    }
    if (message.purchaseDate !== undefined) {
      Timestamp.encode(toTimestamp(message.purchaseDate), writer.uint32(34).fork()).ldelim();
    }
    if (message.deliveryDate !== undefined) {
      Timestamp.encode(toTimestamp(message.deliveryDate), writer.uint32(42).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(82).string(message.name);
    }
    if (message.customExpr !== undefined) {
      Expr.encode(message.customExpr, writer.uint32(90).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          if (tag !== 24) {
            break;
          }

          message.age = reader.int32();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.purchaseDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.deliveryDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.name = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.customExpr = Expr.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserRequest {
    return {
      age: isSet(object.age) ? globalThis.Number(object.age) : 0,
      purchaseDate: isSet(object.purchaseDate) ? fromJsonTimestamp(object.purchaseDate) : undefined,
      deliveryDate: isSet(object.deliveryDate) ? fromJsonTimestamp(object.deliveryDate) : undefined,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      customExpr: isSet(object.customExpr) ? Expr.fromJSON(object.customExpr) : undefined,
    };
  },

  toJSON(message: GetUserRequest): unknown {
    const obj: any = {};
    if (message.age !== 0) {
      obj.age = Math.round(message.age);
    }
    if (message.purchaseDate !== undefined) {
      obj.purchaseDate = message.purchaseDate.toISOString();
    }
    if (message.deliveryDate !== undefined) {
      obj.deliveryDate = message.deliveryDate.toISOString();
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.customExpr !== undefined) {
      obj.customExpr = Expr.toJSON(message.customExpr);
    }
    return obj;
  },

  create(base?: DeepPartial<GetUserRequest>): GetUserRequest {
    return GetUserRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetUserRequest>): GetUserRequest {
    const message = createBaseGetUserRequest();
    message.age = object.age ?? 0;
    message.purchaseDate = object.purchaseDate ?? undefined;
    message.deliveryDate = object.deliveryDate ?? undefined;
    message.name = object.name ?? "";
    message.customExpr = (object.customExpr !== undefined && object.customExpr !== null)
      ? Expr.fromPartial(object.customExpr)
      : undefined;
    return message;
  },
};

function createBaseUser(): User {
  return { id: "", name: "", son: "" };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.son !== "") {
      writer.uint32(26).string(message.son);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.son = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      son: isSet(object.son) ? globalThis.String(object.son) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.son !== "") {
      obj.son = message.son;
    }
    return obj;
  },

  create(base?: DeepPartial<User>): User {
    return User.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<User>): User {
    const message = createBaseUser();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.son = object.son ?? "";
    return message;
  },
};

/**
 * 提供问候语服务。
 * 1. 服务级别的注解 option (google.api.default_host) = "https://huige.api.com";
 */
export type HelloServiceDefinition = typeof HelloServiceDefinition;
export const HelloServiceDefinition = {
  name: "HelloService",
  fullName: "api.v1.HelloService",
  methods: {
    /** 获取用户信息 */
    getUser: {
      name: "GetUser",
      requestType: GetUserRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          508010: [
            new Uint8Array([
              40,
              144,
              234,
              48,
              1,
              154,
              234,
              48,
              15,
              104,
              103,
              46,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
              115,
              46,
              115,
              101,
              116,
              160,
              234,
              48,
              1,
              168,
              234,
              48,
              1,
              176,
              234,
              48,
              10,
              184,
              234,
              48,
              232,
              7,
            ]),
          ],
          578365826: [
            new Uint8Array([
              44,
              58,
              1,
              42,
              90,
              19,
              18,
              17,
              47,
              118,
              49,
              47,
              123,
              97,
              103,
              101,
              61,
              117,
              115,
              101,
              114,
              115,
              47,
              42,
              125,
              34,
              18,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              117,
              115,
              101,
              114,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof globalThis.Date) {
    return o;
  } else if (typeof o === "string") {
    return new globalThis.Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
