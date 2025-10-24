# Explanation: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

OUT_RELATIVE_DIR="../controller/pb"

protoc --go_out=$OUT_RELATIVE_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_RELATIVE_DIR --go-grpc_opt=paths=source_relative \
    ./api_specification_am_rental_management.proto