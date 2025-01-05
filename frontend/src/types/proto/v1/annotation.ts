/* eslint-disable */

export const protobufPackage = "bytebase.v1";

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
