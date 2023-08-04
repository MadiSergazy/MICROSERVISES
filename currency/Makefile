.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency


install_grpccurl:
	curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz" | sudo tar -xz -C /usr/local/bin

grpc_list:
	grpcurl --plaintext localhost:9092 list 

grpc_list_currency:
	grpcurl --plaintext localhost:9092 list 


grpc_describe_currency:
	grpcurl --plaintext localhost:9092 describe Currency.GetRate

grpc_describe_req_res:

grpc_describe_currency:
	grpcurl --plaintext localhost:9092 describe .RateRequest
	grpcurl --plaintext localhost:9092 describe .RateResponse


grpc_client_test:
	grpcurl --plaintext -d '{"base": "GRB", "destination": "USD"}' localhost:9092 Currency.GetRate
