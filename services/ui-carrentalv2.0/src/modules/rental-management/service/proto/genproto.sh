# Plugin to generate proto messages and the service client stub from .proto definitions
# https://github.com/grpc/grpc-web

DIR="./"
OUT_DIR="./"

protoc -I=$DIR api_specification_am_rental_management.proto \
  --js_out=import_style=commonjs,binary:$OUT_DIR \
  --grpc-web_out=import_style=typescript,mode=grpcweb:$OUT_DIR