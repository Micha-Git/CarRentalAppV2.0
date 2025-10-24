import * as jspb from 'google-protobuf'



export class AddCarToFleetRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): AddCarToFleetRequest;
  hasVin(): boolean;
  clearVin(): AddCarToFleetRequest;

  getFleetid(): string;
  setFleetid(value: string): AddCarToFleetRequest;

  getLocation(): string;
  setLocation(value: string): AddCarToFleetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddCarToFleetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddCarToFleetRequest): AddCarToFleetRequest.AsObject;
  static serializeBinaryToWriter(message: AddCarToFleetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddCarToFleetRequest;
  static deserializeBinaryFromReader(message: AddCarToFleetRequest, reader: jspb.BinaryReader): AddCarToFleetRequest;
}

export namespace AddCarToFleetRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
    fleetid: string,
    location: string,
  }
}

export class AddCarToFleetResponse extends jspb.Message {
  getCar(): Car | undefined;
  setCar(value?: Car): AddCarToFleetResponse;
  hasCar(): boolean;
  clearCar(): AddCarToFleetResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): AddCarToFleetResponse;
  hasError(): boolean;
  clearError(): AddCarToFleetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddCarToFleetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddCarToFleetResponse): AddCarToFleetResponse.AsObject;
  static serializeBinaryToWriter(message: AddCarToFleetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddCarToFleetResponse;
  static deserializeBinaryFromReader(message: AddCarToFleetResponse, reader: jspb.BinaryReader): AddCarToFleetResponse;
}

export namespace AddCarToFleetResponse {
  export type AsObject = {
    car?: Car.AsObject,
    error?: ErrorDetail.AsObject,
  }
}

export class ListCarsInFleetRequest extends jspb.Message {
  getFleetid(): string;
  setFleetid(value: string): ListCarsInFleetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCarsInFleetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCarsInFleetRequest): ListCarsInFleetRequest.AsObject;
  static serializeBinaryToWriter(message: ListCarsInFleetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCarsInFleetRequest;
  static deserializeBinaryFromReader(message: ListCarsInFleetRequest, reader: jspb.BinaryReader): ListCarsInFleetRequest;
}

export namespace ListCarsInFleetRequest {
  export type AsObject = {
    fleetid: string,
  }
}

export class ListCarsInFleetResponse extends jspb.Message {
  getCarsList(): Array<Car>;
  setCarsList(value: Array<Car>): ListCarsInFleetResponse;
  clearCarsList(): ListCarsInFleetResponse;
  addCars(value?: Car, index?: number): Car;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): ListCarsInFleetResponse;
  hasError(): boolean;
  clearError(): ListCarsInFleetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCarsInFleetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCarsInFleetResponse): ListCarsInFleetResponse.AsObject;
  static serializeBinaryToWriter(message: ListCarsInFleetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCarsInFleetResponse;
  static deserializeBinaryFromReader(message: ListCarsInFleetResponse, reader: jspb.BinaryReader): ListCarsInFleetResponse;
}

export namespace ListCarsInFleetResponse {
  export type AsObject = {
    carsList: Array<Car.AsObject>,
    error?: ErrorDetail.AsObject,
  }
}

export class RemoveCarFromFleetRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): RemoveCarFromFleetRequest;
  hasVin(): boolean;
  clearVin(): RemoveCarFromFleetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveCarFromFleetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveCarFromFleetRequest): RemoveCarFromFleetRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveCarFromFleetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveCarFromFleetRequest;
  static deserializeBinaryFromReader(message: RemoveCarFromFleetRequest, reader: jspb.BinaryReader): RemoveCarFromFleetRequest;
}

export namespace RemoveCarFromFleetRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
  }
}

export class RemoveCarFromFleetResponse extends jspb.Message {
  getResult(): boolean;
  setResult(value: boolean): RemoveCarFromFleetResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): RemoveCarFromFleetResponse;
  hasError(): boolean;
  clearError(): RemoveCarFromFleetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveCarFromFleetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveCarFromFleetResponse): RemoveCarFromFleetResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveCarFromFleetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveCarFromFleetResponse;
  static deserializeBinaryFromReader(message: RemoveCarFromFleetResponse, reader: jspb.BinaryReader): RemoveCarFromFleetResponse;
}

export namespace RemoveCarFromFleetResponse {
  export type AsObject = {
    result: boolean,
    error?: ErrorDetail.AsObject,
  }
}

export class ViewCarInformationRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): ViewCarInformationRequest;
  hasVin(): boolean;
  clearVin(): ViewCarInformationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ViewCarInformationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ViewCarInformationRequest): ViewCarInformationRequest.AsObject;
  static serializeBinaryToWriter(message: ViewCarInformationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ViewCarInformationRequest;
  static deserializeBinaryFromReader(message: ViewCarInformationRequest, reader: jspb.BinaryReader): ViewCarInformationRequest;
}

export namespace ViewCarInformationRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
  }
}

export class ViewCarInformationResponse extends jspb.Message {
  getCar(): Car | undefined;
  setCar(value?: Car): ViewCarInformationResponse;
  hasCar(): boolean;
  clearCar(): ViewCarInformationResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): ViewCarInformationResponse;
  hasError(): boolean;
  clearError(): ViewCarInformationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ViewCarInformationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ViewCarInformationResponse): ViewCarInformationResponse.AsObject;
  static serializeBinaryToWriter(message: ViewCarInformationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ViewCarInformationResponse;
  static deserializeBinaryFromReader(message: ViewCarInformationResponse, reader: jspb.BinaryReader): ViewCarInformationResponse;
}

export namespace ViewCarInformationResponse {
  export type AsObject = {
    car?: Car.AsObject,
    error?: ErrorDetail.AsObject,
  }
}

export class Vin extends jspb.Message {
  getVin(): string;
  setVin(value: string): Vin;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Vin.AsObject;
  static toObject(includeInstance: boolean, msg: Vin): Vin.AsObject;
  static serializeBinaryToWriter(message: Vin, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Vin;
  static deserializeBinaryFromReader(message: Vin, reader: jspb.BinaryReader): Vin;
}

export namespace Vin {
  export type AsObject = {
    vin: string,
  }
}

export class Car extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): Car;
  hasVin(): boolean;
  clearVin(): Car;

  getBrand(): string;
  setBrand(value: string): Car;

  getModel(): string;
  setModel(value: string): Car;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Car.AsObject;
  static toObject(includeInstance: boolean, msg: Car): Car.AsObject;
  static serializeBinaryToWriter(message: Car, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Car;
  static deserializeBinaryFromReader(message: Car, reader: jspb.BinaryReader): Car;
}

export namespace Car {
  export type AsObject = {
    vin?: Vin.AsObject,
    brand: string,
    model: string,
  }
}

export class Fleet extends jspb.Message {
  getFleetid(): string;
  setFleetid(value: string): Fleet;

  getCarsList(): Array<Car>;
  setCarsList(value: Array<Car>): Fleet;
  clearCarsList(): Fleet;
  addCars(value?: Car, index?: number): Car;

  getLocation(): string;
  setLocation(value: string): Fleet;

  getFleetmanager(): string;
  setFleetmanager(value: string): Fleet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Fleet.AsObject;
  static toObject(includeInstance: boolean, msg: Fleet): Fleet.AsObject;
  static serializeBinaryToWriter(message: Fleet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Fleet;
  static deserializeBinaryFromReader(message: Fleet, reader: jspb.BinaryReader): Fleet;
}

export namespace Fleet {
  export type AsObject = {
    fleetid: string,
    carsList: Array<Car.AsObject>,
    location: string,
    fleetmanager: string,
  }
}

export class ErrorDetail extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): ErrorDetail;

  getDetails(): string;
  setDetails(value: string): ErrorDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ErrorDetail.AsObject;
  static toObject(includeInstance: boolean, msg: ErrorDetail): ErrorDetail.AsObject;
  static serializeBinaryToWriter(message: ErrorDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ErrorDetail;
  static deserializeBinaryFromReader(message: ErrorDetail, reader: jspb.BinaryReader): ErrorDetail;
}

export namespace ErrorDetail {
  export type AsObject = {
    message: string,
    details: string,
  }
}

