/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'alice.checkers.checkers';
const baseSystemInfo = { nextId: 0 };
export const SystemInfo = {
    encode(message, writer = Writer.create()) {
        if (message.nextId !== 0) {
            writer.uint32(8).uint64(message.nextId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSystemInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.nextId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSystemInfo };
        if (object.nextId !== undefined && object.nextId !== null) {
            message.nextId = Number(object.nextId);
        }
        else {
            message.nextId = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.nextId !== undefined && (obj.nextId = message.nextId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSystemInfo };
        if (object.nextId !== undefined && object.nextId !== null) {
            message.nextId = object.nextId;
        }
        else {
            message.nextId = 0;
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
