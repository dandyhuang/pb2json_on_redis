package biz

import (
	"fmt"
	pb2json_api "github.com/dandyhuang/cmd_tools/api/ads"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/op/go-logging"
	"sync"
)

var globalProtoMap map[string]*desc.FileDescriptor
var IsCached = true
var lk sync.RWMutex

func init() {
	globalProtoMap = make(map[string]*desc.FileDescriptor)
}
var log = logging.MustGetLogger("example")

func getProto(path string) *desc.FileDescriptor {
	lk.Lock()
	defer lk.Unlock()

	if IsCached {
		fd, ok := globalProtoMap[path]
		if ok {
			log.Debugf("getProto path:%v cached", path)
			return fd
		}
	}
	p := protoparse.Parser{
	}
	fds, err := p.ParseFiles(path)
	if err != nil {
		log.Errorf("getProto ParseFiles error:%v", err)
		return nil
	}
	// log.Debugf("JsonToPb fd %v, err %v", fds[0], err)
	fd := fds[0]

	if IsCached {
		globalProtoMap[path] = fd
	}

	return fd
}

// JsonToPb 传入proto文件的path, proto中对应的message.name，js的原始数据
// 返回生成的proto.Marshal的[]byte
// example:
// path := "$PROTOPATH/helloworld.proto"
// messageName "helloworld.HelloRequest"
// JsonToPb(path,"helloworld.HelloRequest", []byte(`{"name":"yzh"}`))
func JsonToPb(protoPath, messageName string, jsonStr []byte) ([]byte, error) {
	log.Debugf("JsonToPb protoPath %v", protoPath)

	fd := getProto(protoPath)

	msg := fd.FindMessage(messageName)
	fmt.Println(messageName, msg, "jsonStr:", jsonStr)
	dymsg := dynamic.NewMessage(msg)
	err := dymsg.UnmarshalJSON(jsonStr)
	if err != nil {
		log.Errorf("JsonToPb UnmarshalJSON error:%v", err)
		return nil, nil
	}
	log.Debugf("JsonToPb UnmarshalJSON dymsg %v", dymsg)

	any, err := ptypes.MarshalAny(dymsg)
	if err != nil {
		log.Errorf("JsonToPb MarshalAny error:%v", err)
		return nil, nil
	}
	log.Debugf("JsonToPb marshal any %v", any.Value)

	return any.Value, nil
}


// PbToJson 传入proto的byte数据，返回它对应的json数据
// example:
// path := "$PROTOPATH/helloworld.proto"
// messageName "helloworld.HelloRequest"
// jsonByte, err := PbToJson(path, messageName, pbByte)
func PbToJson(protoPath, messageName string, protoData []byte) ([]byte, error) {
	log.Debugf("PbToJson protoPath %v", protoPath)
	fd := getProto(protoPath)
	msg := fd.FindMessage(messageName)
	dymsg := dynamic.NewMessage(msg)

	err := proto.Unmarshal(protoData, dymsg)
	log.Debugf("PbToJson Unmarshal err:%v", err)

	jsonByte, err := dymsg.MarshalJSON()
	return jsonByte, err
}

func EncodeItemMessage(value []byte) ([]byte, error) {
	data:=&pb2json_api.ItemFeature {
		ItemFeature: value,
	}
	return proto.Marshal(data)
}
