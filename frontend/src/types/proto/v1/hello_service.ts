/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import { Expr } from "../google/type/expr";
import { State, stateFromJSON, stateToJSON } from "./common";

export const protobufPackage = "api.v1";

/** 请求参数 */
export interface GetUserRequest {
  /**
   * 字段解释描述文本
   *
   * syntax. CEL is a C-like expression language. The syntax and semantics of CEL
   *
   * are documented at https://github.com/google/cel-spec.
   * {
   *   "age": 25,
   *   "purchase_date": {
   *     "seconds": 1680054400, // 2023-03-27T00:00:00Z
   *     "nanos": 0
   *   },
   *   "delivery_date": {
   *     "seconds": 1680140800, // 2023-03-28T00:00:00Z
   *     "nanos": 0
   *   },
   *   "name": "productA",
   *   "custom_expr": {
   *     // 自定义表达式的JSON表示，根据具体需求填充
   *     "expr": "value > 10",
   *     "type": "BOOL"
   *   }
   * }
   *
   * 段落1:
   *
   *     标题1
   *     这就是描述1
   *     document.summary.size() < 100
   *
   * 段落2:
   *
   *     标题2
   *     这就是描述2
   *     document.owner == request.auth.claims.email
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

export interface GetDatabaseRequest {
  /**
   * The name of the database to retrieve.
   * Format: instances/{instance}/databases/{database}
   */
  name: string;
}

export interface Database {
  /**
   * The name of the database.
   * Format: instances/{instance}/databases/{database}
   * {database} is the database name in the instance.
   */
  name: string;
  /** The existence of a database on latest sync. */
  syncState: State;
  /** The latest synchronization time. */
  successfulSyncTime:
    | Date
    | undefined;
  /**
   * The project for a database.
   * Format: projects/{project}
   */
  project: string;
  /** The version of database schema. */
  schemaVersion: string;
  /**
   * The environment resource.
   * Format: environments/prod where prod is the environment resource ID.
   */
  environment: string;
  /**
   * The effective environment based on environment tag above and environment
   * tag on the instance. Inheritance follows
   * https://cloud.google.com/resource-manager/docs/tags/tags-overview.
   */
  effectiveEnvironment: string;
  /** Labels will be used for deployment and policy control. */
  labels: { [key: string]: string };
  /** The database is available for DML prior backup. */
  backupAvailable: boolean;
}

export interface Database_LabelsEntry {
  key: string;
  value: string;
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

function createBaseGetDatabaseRequest(): GetDatabaseRequest {
  return { name: "" };
}

export const GetDatabaseRequest = {
  encode(message: GetDatabaseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDatabaseRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetDatabaseRequest();
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

  fromJSON(object: any): GetDatabaseRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: GetDatabaseRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<GetDatabaseRequest>): GetDatabaseRequest {
    return GetDatabaseRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetDatabaseRequest>): GetDatabaseRequest {
    const message = createBaseGetDatabaseRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseDatabase(): Database {
  return {
    name: "",
    syncState: 0,
    successfulSyncTime: undefined,
    project: "",
    schemaVersion: "",
    environment: "",
    effectiveEnvironment: "",
    labels: {},
    backupAvailable: false,
  };
}

export const Database = {
  encode(message: Database, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.syncState !== 0) {
      writer.uint32(24).int32(message.syncState);
    }
    if (message.successfulSyncTime !== undefined) {
      Timestamp.encode(toTimestamp(message.successfulSyncTime), writer.uint32(34).fork()).ldelim();
    }
    if (message.project !== "") {
      writer.uint32(42).string(message.project);
    }
    if (message.schemaVersion !== "") {
      writer.uint32(50).string(message.schemaVersion);
    }
    if (message.environment !== "") {
      writer.uint32(58).string(message.environment);
    }
    if (message.effectiveEnvironment !== "") {
      writer.uint32(66).string(message.effectiveEnvironment);
    }
    Object.entries(message.labels).forEach(([key, value]) => {
      Database_LabelsEntry.encode({ key: key as any, value }, writer.uint32(74).fork()).ldelim();
    });
    if (message.backupAvailable === true) {
      writer.uint32(88).bool(message.backupAvailable);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Database {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDatabase();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.syncState = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.successfulSyncTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.project = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.schemaVersion = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.environment = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.effectiveEnvironment = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          const entry9 = Database_LabelsEntry.decode(reader, reader.uint32());
          if (entry9.value !== undefined) {
            message.labels[entry9.key] = entry9.value;
          }
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.backupAvailable = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Database {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      syncState: isSet(object.syncState) ? stateFromJSON(object.syncState) : 0,
      successfulSyncTime: isSet(object.successfulSyncTime) ? fromJsonTimestamp(object.successfulSyncTime) : undefined,
      project: isSet(object.project) ? globalThis.String(object.project) : "",
      schemaVersion: isSet(object.schemaVersion) ? globalThis.String(object.schemaVersion) : "",
      environment: isSet(object.environment) ? globalThis.String(object.environment) : "",
      effectiveEnvironment: isSet(object.effectiveEnvironment) ? globalThis.String(object.effectiveEnvironment) : "",
      labels: isObject(object.labels)
        ? Object.entries(object.labels).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      backupAvailable: isSet(object.backupAvailable) ? globalThis.Boolean(object.backupAvailable) : false,
    };
  },

  toJSON(message: Database): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.syncState !== 0) {
      obj.syncState = stateToJSON(message.syncState);
    }
    if (message.successfulSyncTime !== undefined) {
      obj.successfulSyncTime = message.successfulSyncTime.toISOString();
    }
    if (message.project !== "") {
      obj.project = message.project;
    }
    if (message.schemaVersion !== "") {
      obj.schemaVersion = message.schemaVersion;
    }
    if (message.environment !== "") {
      obj.environment = message.environment;
    }
    if (message.effectiveEnvironment !== "") {
      obj.effectiveEnvironment = message.effectiveEnvironment;
    }
    if (message.labels) {
      const entries = Object.entries(message.labels);
      if (entries.length > 0) {
        obj.labels = {};
        entries.forEach(([k, v]) => {
          obj.labels[k] = v;
        });
      }
    }
    if (message.backupAvailable === true) {
      obj.backupAvailable = message.backupAvailable;
    }
    return obj;
  },

  create(base?: DeepPartial<Database>): Database {
    return Database.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Database>): Database {
    const message = createBaseDatabase();
    message.name = object.name ?? "";
    message.syncState = object.syncState ?? 0;
    message.successfulSyncTime = object.successfulSyncTime ?? undefined;
    message.project = object.project ?? "";
    message.schemaVersion = object.schemaVersion ?? "";
    message.environment = object.environment ?? "";
    message.effectiveEnvironment = object.effectiveEnvironment ?? "";
    message.labels = Object.entries(object.labels ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = globalThis.String(value);
      }
      return acc;
    }, {});
    message.backupAvailable = object.backupAvailable ?? false;
    return message;
  },
};

function createBaseDatabase_LabelsEntry(): Database_LabelsEntry {
  return { key: "", value: "" };
}

export const Database_LabelsEntry = {
  encode(message: Database_LabelsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Database_LabelsEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDatabase_LabelsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Database_LabelsEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: Database_LabelsEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create(base?: DeepPartial<Database_LabelsEntry>): Database_LabelsEntry {
    return Database_LabelsEntry.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Database_LabelsEntry>): Database_LabelsEntry {
    const message = createBaseDatabase_LabelsEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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
 * 2. 服务级别的注解 option (google.api.oauth_scopes) = "https://www.huige.com/auth/user";
 * 3. 服务级别的注解
 * 4. 服务级别的注解
 * 5. 服务级别的注解
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
    /**
     * 使用<br>换行
     * instances/* /databases/*的注解<br>
     * Gets the database with the given name.<br>
     * The name must be in the format: instances/{instance}/databases/{database}.
     */
    getDatabase: {
      name: "GetDatabase",
      requestType: GetDatabaseRequest,
      requestStream: false,
      responseType: Database,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              36,
              18,
              34,
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
              105,
              110,
              115,
              116,
              97,
              110,
              99,
              101,
              115,
              47,
              42,
              47,
              100,
              97,
              116,
              97,
              98,
              97,
              115,
              101,
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
