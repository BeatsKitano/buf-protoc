/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "api.v1";

export enum State {
  STATE_UNSPECIFIED = 0,
  ACTIVE = 1,
  DELETED = 2,
  UNRECOGNIZED = -1,
}

export function stateFromJSON(object: any): State {
  switch (object) {
    case 0:
    case "STATE_UNSPECIFIED":
      return State.STATE_UNSPECIFIED;
    case 1:
    case "ACTIVE":
      return State.ACTIVE;
    case 2:
    case "DELETED":
      return State.DELETED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return State.UNRECOGNIZED;
  }
}

export function stateToJSON(object: State): string {
  switch (object) {
    case State.STATE_UNSPECIFIED:
      return "STATE_UNSPECIFIED";
    case State.ACTIVE:
      return "ACTIVE";
    case State.DELETED:
      return "DELETED";
    case State.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface OperateResult {
  code: number;
  message: string;
  details: string[];
}

function createBaseOperateResult(): OperateResult {
  return { code: 0, message: "", details: [] };
}

export const OperateResult = {
  encode(message: OperateResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    for (const v of message.details) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OperateResult {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOperateResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.message = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.details.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OperateResult {
    return {
      code: isSet(object.code) ? globalThis.Number(object.code) : 0,
      message: isSet(object.message) ? globalThis.String(object.message) : "",
      details: globalThis.Array.isArray(object?.details) ? object.details.map((e: any) => globalThis.String(e)) : [],
    };
  },

  toJSON(message: OperateResult): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = Math.round(message.code);
    }
    if (message.message !== "") {
      obj.message = message.message;
    }
    if (message.details?.length) {
      obj.details = message.details;
    }
    return obj;
  },

  create(base?: DeepPartial<OperateResult>): OperateResult {
    return OperateResult.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<OperateResult>): OperateResult {
    const message = createBaseOperateResult();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.details = object.details?.map((e) => e) || [];
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
