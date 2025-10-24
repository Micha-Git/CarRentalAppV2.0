import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class ListAvailableCarsRequest extends jspb.Message {
  getStartdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartdate(value?: google_protobuf_timestamp_pb.Timestamp): ListAvailableCarsRequest;
  hasStartdate(): boolean;
  clearStartdate(): ListAvailableCarsRequest;

  getEnddate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEnddate(value?: google_protobuf_timestamp_pb.Timestamp): ListAvailableCarsRequest;
  hasEnddate(): boolean;
  clearEnddate(): ListAvailableCarsRequest;

  getLocation(): string;
  setLocation(value: string): ListAvailableCarsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAvailableCarsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAvailableCarsRequest): ListAvailableCarsRequest.AsObject;
  static serializeBinaryToWriter(message: ListAvailableCarsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAvailableCarsRequest;
  static deserializeBinaryFromReader(message: ListAvailableCarsRequest, reader: jspb.BinaryReader): ListAvailableCarsRequest;
}

export namespace ListAvailableCarsRequest {
  export type AsObject = {
    startdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    enddate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    location: string,
  }
}

export class ListAvailableCarsResponse extends jspb.Message {
  getCarsList(): Array<RentableCar>;
  setCarsList(value: Array<RentableCar>): ListAvailableCarsResponse;
  clearCarsList(): ListAvailableCarsResponse;
  addCars(value?: RentableCar, index?: number): RentableCar;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): ListAvailableCarsResponse;
  hasError(): boolean;
  clearError(): ListAvailableCarsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAvailableCarsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAvailableCarsResponse): ListAvailableCarsResponse.AsObject;
  static serializeBinaryToWriter(message: ListAvailableCarsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAvailableCarsResponse;
  static deserializeBinaryFromReader(message: ListAvailableCarsResponse, reader: jspb.BinaryReader): ListAvailableCarsResponse;
}

export namespace ListAvailableCarsResponse {
  export type AsObject = {
    carsList: Array<RentableCar.AsObject>,
    error?: ErrorDetail.AsObject,
  }
}

export class ListCarRentalsRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): ListCarRentalsRequest;
  hasVin(): boolean;
  clearVin(): ListCarRentalsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCarRentalsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCarRentalsRequest): ListCarRentalsRequest.AsObject;
  static serializeBinaryToWriter(message: ListCarRentalsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCarRentalsRequest;
  static deserializeBinaryFromReader(message: ListCarRentalsRequest, reader: jspb.BinaryReader): ListCarRentalsRequest;
}

export namespace ListCarRentalsRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
  }
}

export class ListCarRentalsResponse extends jspb.Message {
  getRentalsList(): Array<Rental>;
  setRentalsList(value: Array<Rental>): ListCarRentalsResponse;
  clearRentalsList(): ListCarRentalsResponse;
  addRentals(value?: Rental, index?: number): Rental;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): ListCarRentalsResponse;
  hasError(): boolean;
  clearError(): ListCarRentalsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCarRentalsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCarRentalsResponse): ListCarRentalsResponse.AsObject;
  static serializeBinaryToWriter(message: ListCarRentalsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCarRentalsResponse;
  static deserializeBinaryFromReader(message: ListCarRentalsResponse, reader: jspb.BinaryReader): ListCarRentalsResponse;
}

export namespace ListCarRentalsResponse {
  export type AsObject = {
    rentalsList: Array<Rental.AsObject>,
    error?: ErrorDetail.AsObject,
  }
}

export class ListCustomerRentalsRequest extends jspb.Message {
  getCustomerid(): string;
  setCustomerid(value: string): ListCustomerRentalsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCustomerRentalsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCustomerRentalsRequest): ListCustomerRentalsRequest.AsObject;
  static serializeBinaryToWriter(message: ListCustomerRentalsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCustomerRentalsRequest;
  static deserializeBinaryFromReader(message: ListCustomerRentalsRequest, reader: jspb.BinaryReader): ListCustomerRentalsRequest;
}

export namespace ListCustomerRentalsRequest {
  export type AsObject = {
    customerid: string,
  }
}

export class ListCustomerRentalsResponse extends jspb.Message {
  getRentalsList(): Array<Rental>;
  setRentalsList(value: Array<Rental>): ListCustomerRentalsResponse;
  clearRentalsList(): ListCustomerRentalsResponse;
  addRentals(value?: Rental, index?: number): Rental;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): ListCustomerRentalsResponse;
  hasError(): boolean;
  clearError(): ListCustomerRentalsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCustomerRentalsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCustomerRentalsResponse): ListCustomerRentalsResponse.AsObject;
  static serializeBinaryToWriter(message: ListCustomerRentalsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCustomerRentalsResponse;
  static deserializeBinaryFromReader(message: ListCustomerRentalsResponse, reader: jspb.BinaryReader): ListCustomerRentalsResponse;
}

export namespace ListCustomerRentalsResponse {
  export type AsObject = {
    rentalsList: Array<Rental.AsObject>,
    error?: ErrorDetail.AsObject,
  }
}

export class AddCarToRentalRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): AddCarToRentalRequest;
  hasVin(): boolean;
  clearVin(): AddCarToRentalRequest;

  getLocation(): string;
  setLocation(value: string): AddCarToRentalRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddCarToRentalRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddCarToRentalRequest): AddCarToRentalRequest.AsObject;
  static serializeBinaryToWriter(message: AddCarToRentalRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddCarToRentalRequest;
  static deserializeBinaryFromReader(message: AddCarToRentalRequest, reader: jspb.BinaryReader): AddCarToRentalRequest;
}

export namespace AddCarToRentalRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
    location: string,
  }
}

export class AddCarToRentalResponse extends jspb.Message {
  getCar(): RentableCar | undefined;
  setCar(value?: RentableCar): AddCarToRentalResponse;
  hasCar(): boolean;
  clearCar(): AddCarToRentalResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): AddCarToRentalResponse;
  hasError(): boolean;
  clearError(): AddCarToRentalResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddCarToRentalResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddCarToRentalResponse): AddCarToRentalResponse.AsObject;
  static serializeBinaryToWriter(message: AddCarToRentalResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddCarToRentalResponse;
  static deserializeBinaryFromReader(message: AddCarToRentalResponse, reader: jspb.BinaryReader): AddCarToRentalResponse;
}

export namespace AddCarToRentalResponse {
  export type AsObject = {
    car?: RentableCar.AsObject,
    error?: ErrorDetail.AsObject,
  }
}

export class RemoveRentableCarRequest extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): RemoveRentableCarRequest;
  hasVin(): boolean;
  clearVin(): RemoveRentableCarRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveRentableCarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveRentableCarRequest): RemoveRentableCarRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveRentableCarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveRentableCarRequest;
  static deserializeBinaryFromReader(message: RemoveRentableCarRequest, reader: jspb.BinaryReader): RemoveRentableCarRequest;
}

export namespace RemoveRentableCarRequest {
  export type AsObject = {
    vin?: Vin.AsObject,
  }
}

export class RemoveRentableCarResponse extends jspb.Message {
  getResult(): boolean;
  setResult(value: boolean): RemoveRentableCarResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): RemoveRentableCarResponse;
  hasError(): boolean;
  clearError(): RemoveRentableCarResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveRentableCarResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveRentableCarResponse): RemoveRentableCarResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveRentableCarResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveRentableCarResponse;
  static deserializeBinaryFromReader(message: RemoveRentableCarResponse, reader: jspb.BinaryReader): RemoveRentableCarResponse;
}

export namespace RemoveRentableCarResponse {
  export type AsObject = {
    result: boolean,
    error?: ErrorDetail.AsObject,
  }
}

export class RentCarRequest extends jspb.Message {
  getCustomerid(): string;
  setCustomerid(value: string): RentCarRequest;

  getStartdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartdate(value?: google_protobuf_timestamp_pb.Timestamp): RentCarRequest;
  hasStartdate(): boolean;
  clearStartdate(): RentCarRequest;

  getEnddate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEnddate(value?: google_protobuf_timestamp_pb.Timestamp): RentCarRequest;
  hasEnddate(): boolean;
  clearEnddate(): RentCarRequest;

  getVin(): Vin | undefined;
  setVin(value?: Vin): RentCarRequest;
  hasVin(): boolean;
  clearVin(): RentCarRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RentCarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RentCarRequest): RentCarRequest.AsObject;
  static serializeBinaryToWriter(message: RentCarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RentCarRequest;
  static deserializeBinaryFromReader(message: RentCarRequest, reader: jspb.BinaryReader): RentCarRequest;
}

export namespace RentCarRequest {
  export type AsObject = {
    customerid: string,
    startdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    enddate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    vin?: Vin.AsObject,
  }
}

export class RentCarResponse extends jspb.Message {
  getRental(): Rental | undefined;
  setRental(value?: Rental): RentCarResponse;
  hasRental(): boolean;
  clearRental(): RentCarResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): RentCarResponse;
  hasError(): boolean;
  clearError(): RentCarResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RentCarResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RentCarResponse): RentCarResponse.AsObject;
  static serializeBinaryToWriter(message: RentCarResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RentCarResponse;
  static deserializeBinaryFromReader(message: RentCarResponse, reader: jspb.BinaryReader): RentCarResponse;
}

export namespace RentCarResponse {
  export type AsObject = {
    rental?: Rental.AsObject,
    error?: ErrorDetail.AsObject,
  }
}

export class CancelRentalRequest extends jspb.Message {
  getCustomerid(): string;
  setCustomerid(value: string): CancelRentalRequest;

  getRentalid(): string;
  setRentalid(value: string): CancelRentalRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CancelRentalRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CancelRentalRequest): CancelRentalRequest.AsObject;
  static serializeBinaryToWriter(message: CancelRentalRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CancelRentalRequest;
  static deserializeBinaryFromReader(message: CancelRentalRequest, reader: jspb.BinaryReader): CancelRentalRequest;
}

export namespace CancelRentalRequest {
  export type AsObject = {
    customerid: string,
    rentalid: string,
  }
}

export class CancelRentalResponse extends jspb.Message {
  getResult(): boolean;
  setResult(value: boolean): CancelRentalResponse;

  getError(): ErrorDetail | undefined;
  setError(value?: ErrorDetail): CancelRentalResponse;
  hasError(): boolean;
  clearError(): CancelRentalResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CancelRentalResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CancelRentalResponse): CancelRentalResponse.AsObject;
  static serializeBinaryToWriter(message: CancelRentalResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CancelRentalResponse;
  static deserializeBinaryFromReader(message: CancelRentalResponse, reader: jspb.BinaryReader): CancelRentalResponse;
}

export namespace CancelRentalResponse {
  export type AsObject = {
    result: boolean,
    error?: ErrorDetail.AsObject,
  }
}

export class Customer extends jspb.Message {
  getId(): string;
  setId(value: string): Customer;

  getName(): string;
  setName(value: string): Customer;

  getRentalsList(): Array<Rental>;
  setRentalsList(value: Array<Rental>): Customer;
  clearRentalsList(): Customer;
  addRentals(value?: Rental, index?: number): Rental;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Customer.AsObject;
  static toObject(includeInstance: boolean, msg: Customer): Customer.AsObject;
  static serializeBinaryToWriter(message: Customer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Customer;
  static deserializeBinaryFromReader(message: Customer, reader: jspb.BinaryReader): Customer;
}

export namespace Customer {
  export type AsObject = {
    id: string,
    name: string,
    rentalsList: Array<Rental.AsObject>,
  }
}

export class Rental extends jspb.Message {
  getId(): string;
  setId(value: string): Rental;

  getStartdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartdate(value?: google_protobuf_timestamp_pb.Timestamp): Rental;
  hasStartdate(): boolean;
  clearStartdate(): Rental;

  getEnddate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEnddate(value?: google_protobuf_timestamp_pb.Timestamp): Rental;
  hasEnddate(): boolean;
  clearEnddate(): Rental;

  getCar(): RentableCar | undefined;
  setCar(value?: RentableCar): Rental;
  hasCar(): boolean;
  clearCar(): Rental;

  getPrice(): number;
  setPrice(value: number): Rental;

  getCustomerid(): string;
  setCustomerid(value: string): Rental;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Rental.AsObject;
  static toObject(includeInstance: boolean, msg: Rental): Rental.AsObject;
  static serializeBinaryToWriter(message: Rental, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Rental;
  static deserializeBinaryFromReader(message: Rental, reader: jspb.BinaryReader): Rental;
}

export namespace Rental {
  export type AsObject = {
    id: string,
    startdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    enddate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    car?: RentableCar.AsObject,
    price: number,
    customerid: string,
  }
}

export class RentableCar extends jspb.Message {
  getVin(): Vin | undefined;
  setVin(value?: Vin): RentableCar;
  hasVin(): boolean;
  clearVin(): RentableCar;

  getBrand(): string;
  setBrand(value: string): RentableCar;

  getModel(): string;
  setModel(value: string): RentableCar;

  getLocation(): string;
  setLocation(value: string): RentableCar;

  getPriceperday(): number;
  setPriceperday(value: number): RentableCar;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RentableCar.AsObject;
  static toObject(includeInstance: boolean, msg: RentableCar): RentableCar.AsObject;
  static serializeBinaryToWriter(message: RentableCar, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RentableCar;
  static deserializeBinaryFromReader(message: RentableCar, reader: jspb.BinaryReader): RentableCar;
}

export namespace RentableCar {
  export type AsObject = {
    vin?: Vin.AsObject,
    brand: string,
    model: string,
    location: string,
    priceperday: number,
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

