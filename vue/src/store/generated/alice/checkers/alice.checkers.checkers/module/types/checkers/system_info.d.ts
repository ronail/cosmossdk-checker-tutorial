import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "alice.checkers.checkers";
export interface SystemInfo {
    nextId: number;
}
export declare const SystemInfo: {
    encode(message: SystemInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SystemInfo;
    fromJSON(object: any): SystemInfo;
    toJSON(message: SystemInfo): unknown;
    fromPartial(object: DeepPartial<SystemInfo>): SystemInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
