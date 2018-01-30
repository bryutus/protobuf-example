PROTO_CMD=protoc
SRC_DIR=/Users/segawar/dev/src/github.com/bryutus/protobuf-tutorial
DST_DIR=/Users/segawar/dev/src/github.com/bryutus/protobuf-tutorial/tutorial

protocol:
	$(PROTO_CMD) -I=$(SRC_DIR) --go_out=$(DST_DIR) $(SRC_DIR)/addressbook.proto
