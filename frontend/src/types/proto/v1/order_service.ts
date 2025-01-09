/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import { OperateResult } from "./common";

export const protobufPackage = "api.v1";

/** 订单状态 */
export enum OrderStatus {
  /** ORDER_STATUS_UNSPECIFIED - 订单状态 */
  ORDER_STATUS_UNSPECIFIED = 0,
  /** ORDER_STATUS_CREATED - 订单创建 */
  ORDER_STATUS_CREATED = 1,
  /** ORDER_STATUS_PAID - 订单已支付 */
  ORDER_STATUS_PAID = 2,
  /** ORDER_STATUS_SHIPPED - 订单已发货 */
  ORDER_STATUS_SHIPPED = 3,
  /** ORDER_STATUS_DELIVERED - 订单已送达 */
  ORDER_STATUS_DELIVERED = 4,
  /** ORDER_STATUS_CANCELLED - 订单已取消 */
  ORDER_STATUS_CANCELLED = 5,
  UNRECOGNIZED = -1,
}

export function orderStatusFromJSON(object: any): OrderStatus {
  switch (object) {
    case 0:
    case "ORDER_STATUS_UNSPECIFIED":
      return OrderStatus.ORDER_STATUS_UNSPECIFIED;
    case 1:
    case "ORDER_STATUS_CREATED":
      return OrderStatus.ORDER_STATUS_CREATED;
    case 2:
    case "ORDER_STATUS_PAID":
      return OrderStatus.ORDER_STATUS_PAID;
    case 3:
    case "ORDER_STATUS_SHIPPED":
      return OrderStatus.ORDER_STATUS_SHIPPED;
    case 4:
    case "ORDER_STATUS_DELIVERED":
      return OrderStatus.ORDER_STATUS_DELIVERED;
    case 5:
    case "ORDER_STATUS_CANCELLED":
      return OrderStatus.ORDER_STATUS_CANCELLED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OrderStatus.UNRECOGNIZED;
  }
}

export function orderStatusToJSON(object: OrderStatus): string {
  switch (object) {
    case OrderStatus.ORDER_STATUS_UNSPECIFIED:
      return "ORDER_STATUS_UNSPECIFIED";
    case OrderStatus.ORDER_STATUS_CREATED:
      return "ORDER_STATUS_CREATED";
    case OrderStatus.ORDER_STATUS_PAID:
      return "ORDER_STATUS_PAID";
    case OrderStatus.ORDER_STATUS_SHIPPED:
      return "ORDER_STATUS_SHIPPED";
    case OrderStatus.ORDER_STATUS_DELIVERED:
      return "ORDER_STATUS_DELIVERED";
    case OrderStatus.ORDER_STATUS_CANCELLED:
      return "ORDER_STATUS_CANCELLED";
    case OrderStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Order {
  /** 订单id */
  id: string;
  /** 订单no */
  no: string;
  /** 订单创建时间 */
  createTime:
    | Date
    | undefined;
  /** 订单更新时间 */
  updateTime:
    | Date
    | undefined;
  /** 订单状态 */
  status: OrderStatus;
  /** 订单金额 */
  amount: number;
}

/** 创建订单参数 */
export interface CreateOrderRequest {
  /** 订单金额 */
  amount: number;
  /** 明细 */
  details: string[];
}

/** 获取订单参数 */
export interface GetOrderRequest {
  name: string;
}

/** 删除订单参数 */
export interface DeleteOrderRequest {
  name: string;
}

function createBaseOrder(): Order {
  return { id: "", no: "", createTime: undefined, updateTime: undefined, status: 0, amount: 0 };
}

export const Order = {
  encode(message: Order, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.no !== "") {
      writer.uint32(18).string(message.no);
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(26).fork()).ldelim();
    }
    if (message.updateTime !== undefined) {
      Timestamp.encode(toTimestamp(message.updateTime), writer.uint32(34).fork()).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(40).int32(message.status);
    }
    if (message.amount !== 0) {
      writer.uint32(49).double(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder();
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

          message.no = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.updateTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 6:
          if (tag !== 49) {
            break;
          }

          message.amount = reader.double();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Order {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      no: isSet(object.no) ? globalThis.String(object.no) : "",
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      updateTime: isSet(object.updateTime) ? fromJsonTimestamp(object.updateTime) : undefined,
      status: isSet(object.status) ? orderStatusFromJSON(object.status) : 0,
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
    };
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.no !== "") {
      obj.no = message.no;
    }
    if (message.createTime !== undefined) {
      obj.createTime = message.createTime.toISOString();
    }
    if (message.updateTime !== undefined) {
      obj.updateTime = message.updateTime.toISOString();
    }
    if (message.status !== 0) {
      obj.status = orderStatusToJSON(message.status);
    }
    if (message.amount !== 0) {
      obj.amount = message.amount;
    }
    return obj;
  },

  create(base?: DeepPartial<Order>): Order {
    return Order.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Order>): Order {
    const message = createBaseOrder();
    message.id = object.id ?? "";
    message.no = object.no ?? "";
    message.createTime = object.createTime ?? undefined;
    message.updateTime = object.updateTime ?? undefined;
    message.status = object.status ?? 0;
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseCreateOrderRequest(): CreateOrderRequest {
  return { amount: 0, details: [] };
}

export const CreateOrderRequest = {
  encode(message: CreateOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== 0) {
      writer.uint32(9).double(message.amount);
    }
    for (const v of message.details) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 9) {
            break;
          }

          message.amount = reader.double();
          continue;
        case 2:
          if (tag !== 18) {
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

  fromJSON(object: any): CreateOrderRequest {
    return {
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
      details: globalThis.Array.isArray(object?.details) ? object.details.map((e: any) => globalThis.String(e)) : [],
    };
  },

  toJSON(message: CreateOrderRequest): unknown {
    const obj: any = {};
    if (message.amount !== 0) {
      obj.amount = message.amount;
    }
    if (message.details?.length) {
      obj.details = message.details;
    }
    return obj;
  },

  create(base?: DeepPartial<CreateOrderRequest>): CreateOrderRequest {
    return CreateOrderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateOrderRequest>): CreateOrderRequest {
    const message = createBaseCreateOrderRequest();
    message.amount = object.amount ?? 0;
    message.details = object.details?.map((e) => e) || [];
    return message;
  },
};

function createBaseGetOrderRequest(): GetOrderRequest {
  return { name: "" };
}

export const GetOrderRequest = {
  encode(message: GetOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(82).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 10:
          if (tag !== 82) {
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

  fromJSON(object: any): GetOrderRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: GetOrderRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<GetOrderRequest>): GetOrderRequest {
    return GetOrderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetOrderRequest>): GetOrderRequest {
    const message = createBaseGetOrderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseDeleteOrderRequest(): DeleteOrderRequest {
  return { name: "" };
}

export const DeleteOrderRequest = {
  encode(message: DeleteOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteOrderRequest();
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

  fromJSON(object: any): DeleteOrderRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: DeleteOrderRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<DeleteOrderRequest>): DeleteOrderRequest {
    return DeleteOrderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DeleteOrderRequest>): DeleteOrderRequest {
    const message = createBaseDeleteOrderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

/** 订单服务 */
export type OrderServiceDefinition = typeof OrderServiceDefinition;
export const OrderServiceDefinition = {
  name: "OrderService",
  fullName: "api.v1.OrderService",
  methods: {
    /** 创建订单 */
    createOrder: {
      name: "CreateOrder",
      requestType: CreateOrderRequest,
      requestStream: false,
      responseType: Order,
      responseStream: false,
      options: {
        _unknownFields: {
          508010: [
            new Uint8Array([
              37,
              144,
              234,
              48,
              1,
              154,
              234,
              48,
              12,
              111,
              114,
              100,
              101,
              114,
              46,
              99,
              114,
              101,
              97,
              116,
              101,
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
          578365826: [new Uint8Array([12, 34, 10, 47, 118, 49, 47, 111, 114, 100, 101, 114, 115])],
        },
      },
    },
    /** 获取订单 */
    getOrder: {
      name: "GetOrder",
      requestType: GetOrderRequest,
      requestStream: false,
      responseType: Order,
      responseStream: false,
      options: {
        _unknownFields: {
          508010: [
            new Uint8Array([
              34,
              144,
              234,
              48,
              1,
              154,
              234,
              48,
              9,
              111,
              114,
              100,
              101,
              114,
              46,
              103,
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
              111,
              114,
              100,
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
    /** 删除订单 */
    deleteOrder: {
      name: "DeleteOrder",
      requestType: DeleteOrderRequest,
      requestStream: false,
      responseType: OperateResult,
      responseStream: false,
      options: {
        _unknownFields: {
          508010: [
            new Uint8Array([
              52,
              144,
              234,
              48,
              1,
              154,
              234,
              48,
              27,
              111,
              114,
              100,
              101,
              114,
              46,
              113,
              117,
              101,
              114,
              121,
              32,
              38,
              38,
              32,
              111,
              114,
              100,
              101,
              114,
              46,
              100,
              101,
              108,
              101,
              116,
              101,
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
              21,
              42,
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
              111,
              114,
              100,
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
