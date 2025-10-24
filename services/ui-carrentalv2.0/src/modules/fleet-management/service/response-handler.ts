import { ErrorDetail } from './proto/api_specification_am_fleet_management_pb';
import * as grpcWeb from 'grpc-web';



export function handleServiceResponse<T>(
    action: string,
    err: grpcWeb.RpcError,
    response: any,
    errorDetail: ErrorDetail | undefined,
    promiseResolveCandidate: T | undefined,
    resolve: (value: T | PromiseLike<T>) => void,
    reject: (reason?: any) => void): void {

    if (err) {
        reject(`Error ${action}: Error while communicating with service: ${""}`);
    } else if (response === null) {
        reject(`Error ${action}: ${"http://localhost:82"} responded with empty body`)
    } else if (errorDetail !== undefined) {
        reject(`Error ${action}: ${errorDetail.getDetails()}`);
    } else if (promiseResolveCandidate === undefined) {
        reject(`Error ${action}: ${"http://localhost:82"} did not respond as expected`);
    } else {
        resolve(promiseResolveCandidate)
    }
}