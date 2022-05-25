// 建造者模式 声明参数复杂的对象
package builder

import "sync"

type Message struct {
	Header *Header 
	Body   *Body
}
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

// 建造者
type builder struct {
	once *sync.Once
	msg  *Message
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{}, Body: &Body{}},
	}
}

// 设置Message结构体的Header结构体的SrcAddr信息
func (b *builder) SetSrcAddr(value string) *builder {
	b.msg.Header.SrcAddr = value
	return b
}

// 设置Message结构体的Header结构体的SrcPort信息
func (b *builder) SetSrcPort(value uint64) *builder {
	b.msg.Header.SrcPort = value
	return b
}

// 设置Message结构体的Header结构体的DestAddr信息
func (b *builder) SetDestAddr(value string) *builder {
	b.msg.Header.DestAddr = value
	return b
}

// 设置Message结构体的Header结构体的DestPort信息
func (b *builder) SetDestPort(value uint64) *builder {
	b.msg.Header.DestPort = value
	return b
}

// 设置Message结构体的Header结构体的Items信息
func (b *builder) SetHeaderItems(key, value string) *builder {
	// 保证map只初始化一次
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

// 设置Message结构体的Body结构体的Items信息
func (b *builder) SetBodyItems(value string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, value)
	return b
}

// 获取Message结构体
func (b *builder) GetMessage() *Message {
	return b.msg
}
