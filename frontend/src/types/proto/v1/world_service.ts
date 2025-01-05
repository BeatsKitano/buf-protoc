/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "api.v1";

export interface Req2 {
  /** ip地址 */
  ip: Uint8Array;
  /** 掩码 */
  mask: number;
  /** 年龄 */
  age: number;
  /** 日期 */
  deliveryDate:
    | Date
    | undefined;
  /** 姓名 */
  name: string;
}

function createBaseReq2(): Req2 {
  return { ip: new Uint8Array(0), mask: 0, age: 0, deliveryDate: undefined, name: "" };
}

export const Req2 = {
  encode(message: Req2, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Req2 {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReq2();
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

  fromJSON(object: any): Req2 {
    return {
      ip: isSet(object.ip) ? bytesFromBase64(object.ip) : new Uint8Array(0),
      mask: isSet(object.mask) ? globalThis.Number(object.mask) : 0,
      age: isSet(object.age) ? globalThis.Number(object.age) : 0,
      deliveryDate: isSet(object.deliveryDate) ? fromJsonTimestamp(object.deliveryDate) : undefined,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
    };
  },

  toJSON(message: Req2): unknown {
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

  create(base?: DeepPartial<Req2>): Req2 {
    return Req2.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Req2>): Req2 {
    const message = createBaseReq2();
    message.ip = object.ip ?? new Uint8Array(0);
    message.mask = object.mask ?? 0;
    message.age = object.age ?? 0;
    message.deliveryDate = object.deliveryDate ?? undefined;
    message.name = object.name ?? "";
    return message;
  },
};

export type WorldServiceDefinition = typeof WorldServiceDefinition;
export const WorldServiceDefinition = {
  name: "WorldService",
  fullName: "api.v1.WorldService",
  methods: {
    echo: {
      name: "Echo",
      requestType: Req2,
      requestStream: false,
      responseType: Req2,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              18,
              19,
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
              50,
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
