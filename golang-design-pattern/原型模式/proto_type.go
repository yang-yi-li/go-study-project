/**
原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。

原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。

 */

package proto_type

type CloneAble interface {
	Clone() CloneAble
}

type ProtoTypeManager struct {
	protoTypes map[string]CloneAble
}

func NewProtoTypeManager() *ProtoTypeManager {
	return &ProtoTypeManager{
		protoTypes: make(map[string]CloneAble),
	}
}

func (p *ProtoTypeManager) Get(name string) CloneAble {
	return p.protoTypes[name]
}

func (p *ProtoTypeManager) Set(name string, prototype CloneAble) {
	p.protoTypes[name] = prototype
}
