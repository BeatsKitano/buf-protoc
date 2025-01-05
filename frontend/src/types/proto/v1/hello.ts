/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "api.v1";

/** 请求参数 */
export interface Req {
  /** 1. 必须是ipv4或ipv6 */
  ip: Uint8Array;
  /**
   * 掩码
   * 掩码范围
   */
  mask: number;
  /**
   * 年龄
   * 3. 年龄限制
   */
  age: number;
  /**
   * 日期
   * 4. 日期限制
   */
  deliveryDate:
    | Date
    | undefined;
  /**
   * 姓名
   * The name of the branch.
   * Format: projects/{project}/branches/{branch}
   * {branch} should be the id of a sheet.
   */
  name: string;
}

export interface User {
  /**
   * The name of the user.
   * Format: users/{user}. {user} is a system-generated unique ID.
   */
  name: string;
}

function createBaseReq(): Req {
  return { ip: new Uint8Array(0), mask: 0, age: 0, deliveryDate: undefined, name: "" };
}

export const Req = {
  encode(message: Req, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ip.length !== 0) {
      writer.uint32(10).bytes(message.ip);
    }
    if (message.mask !== 0) {
      writer.uint32(16).uint32(message.mask);
    }
    if (message.age !== 0) {
      writer.uint32(24).int32(message.age);
    }
    if (message.deliveryDate !== undefined) {
      Timestamp.encode(toTimestamp(message.deliveryDate), writer.uint32(34).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Req {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ip = reader.bytes();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.mask = reader.uint32();
          continue;
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

          message.deliveryDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Req {
    return {
      ip: isSet(object.ip) ? bytesFromBase64(object.ip) : new Uint8Array(0),
      mask: isSet(object.mask) ? globalThis.Number(object.mask) : 0,
      age: isSet(object.age) ? globalThis.Number(object.age) : 0,
      deliveryDate: isSet(object.deliveryDate) ? fromJsonTimestamp(object.deliveryDate) : undefined,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
    };
  },

  toJSON(message: Req): unknown {
    const obj: any = {};
    if (message.ip.length !== 0) {
      obj.ip = base64FromBytes(message.ip);
    }
    if (message.mask !== 0) {
      obj.mask = Math.round(message.mask);
    }
    if (message.age !== 0) {
      obj.age = Math.round(message.age);
    }
    if (message.deliveryDate !== undefined) {
      obj.deliveryDate = message.deliveryDate.toISOString();
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<Req>): Req {
    return Req.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Req>): Req {
    const message = createBaseReq();
    message.ip = object.ip ?? new Uint8Array(0);
    message.mask = object.mask ?? 0;
    message.age = object.age ?? 0;
    message.deliveryDate = object.deliveryDate ?? undefined;
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseUser(): User {
  return { name: "" };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
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

          message.name = reader.string();
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
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<User>): User {
    return User.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<User>): User {
    const message = createBaseUser();
    message.name = object.name ?? "";
    return message;
  },
};

export type HelloServiceDefinition = typeof HelloServiceDefinition;
export const HelloServiceDefinition = {
  name: "HelloService",
  fullName: "api.v1.HelloService",
  methods: {
    getUser: {
      name: "GetUser",
      requestType: Req,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              20,
              18,
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(globalThis.String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
