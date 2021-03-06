# Copyright 2018-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

PROTO_SRC := $(wildcard *.proto)
PROTO_GEN_SRC := $(PROTO_SRC:.proto=.pb.go)
PROTO_GEN_SRC_GRPC := $(addprefix grpc/,$(PROTO_GEN_SRC))
PROTO_GEN_SRC_TTRPC := $(addprefix ttrpc/,$(PROTO_GEN_SRC))

$(PROTO_GEN_SRC_TTRPC): $(PROTO_SRC)
	protoc \
		--gogottrpc_out=plugins=ttrpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:ttrpc \
		$^

$(PROTO_GEN_SRC_GRPC): $(PROTO_SRC)
	protoc \
		--gogo_out=plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:grpc \
		$^

proto: $(PROTO_GEN_SRC_GRPC) $(PROTO_GEN_SRC_TTRPC)

clean:
	- rm -f $(PROTO_GEN_SRC_GRPC) $(PROTO_GEN_SRC_TTRPC)

.PHONY: clean proto
