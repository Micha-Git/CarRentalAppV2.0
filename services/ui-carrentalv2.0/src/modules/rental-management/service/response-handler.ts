import * as grpcWeb from 'grpc-web';
import {environment} from "../../../environments/environment";
import {ErrorDetail} from "./proto/api_specification_am_rental_management_pb";

export function handleServiceResponse<T>(
    action: string,
    err: grpcWeb.RpcError,
    response: any,
    errorDetail: ErrorDetail | undefined,
    promiseResolveCandidate: T | undefined,
    resolve: (value: T | PromiseLike<T>) => void,
    reject: (reason?: any) => void
): void {
    if (err) {
        reject(`Error ${action}: Error while communicating with service: ${err.message}`);
    } else if (response === null) {
        reject(`Error ${action}: ${environment.rentalManagementServiceHost} or 
               ${environment.fleetManagementServiceHost} responded with empty body`)
    } else if (errorDetail !== undefined) {
        reject(`Error ${action}: ${errorDetail.getDetails()}`);
    } else if (promiseResolveCandidate === undefined) {
        reject(`Error ${action}: ${environment.rentalManagementServiceHost} or 
                ${environment.fleetManagementServiceHost} did not respond as expected`);
    } else {
        resolve(promiseResolveCandidate)
    }
}