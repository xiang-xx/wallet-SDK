package sdk

import (
	"bytes"
	"fmt"
)

type Recipient struct {
	Address string
	Amount  string
}

func NewRecipient(address string, amount string) *Recipient {
	return &Recipient{
		Address: address,
		Amount:  amount,
	}
}

// 支持转账多地址的对象 ( golang 的 sdk api 方法不支持数组类型，因此需要这种对象 )
type MultiRecipient struct {
	recipients []*Recipient
}

func NewMultiRecipient() *MultiRecipient {
	return &MultiRecipient{
		recipients: []*Recipient{},
	}
}

// 添加转账对象
func (m *MultiRecipient) Add(recipient *Recipient) {
	m.recipients = append(m.recipients, recipient)
}

// 添加转账对象
func (m *MultiRecipient) AddAddress(address string, amount string) {
	m.recipients = append(m.recipients, NewRecipient(address, amount))
}

// 移除指定的转账对象
// @param address 要移除的对象的地址
// @return 如果多地址中存在指定对象，移除后会返回它
func (m *MultiRecipient) Remove(address string) *Recipient {
	index := m.IndexOfAddress(address)
	if index == -1 {
		return nil
	}

	var r = m.recipients[index]
	m.recipients = append(m.recipients[:index], m.recipients[index+1:]...)
	return r
}

// 获取指定下标的元素
// @return 如果下标越界，返回 nil
func (m *MultiRecipient) ObjectAt(index int) *Recipient {
	if index < 0 || index >= len(m.recipients) {
		return nil
	}
	return m.recipients[index]
}

// 查询指定元素的下标
// @return 如果没有找到，返回 -1
func (m *MultiRecipient) IndexOf(object *Recipient) int {
	return m.IndexOfAddress(object.Address)
}

// 查询指定元素的下标
// @return 如果没有找到，返回 -1
func (m *MultiRecipient) IndexOfAddress(address string) int {
	for i, v := range m.recipients {
		if address == v.Address {
			return i
		}
	}
	return -1
}

// 输出多地址信息，方便 debug
func (m *MultiRecipient) DebugInfo() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Contains %d recipients : \n", len(m.recipients)))
	for _, v := range m.recipients {
		buffer.WriteString(fmt.Sprintf("\tto: %s, amount: %s,\n", v.Address, v.Amount))
	}
	buffer.WriteString(")")
	return buffer.String()
}
