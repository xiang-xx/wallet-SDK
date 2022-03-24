package sdk

import "testing"

func TestRecipient(t *testing.T) {
	m := &MultiRecipient{}

	m.AddAddress("XXXca", "124")
	m.Add(&Recipient{"BBBB3", "124"})
	m.Add(&Recipient{"BBCCC", "1000"})

	if m.Count() != 3 {
		t.Fatal("元素个数错误")
	}
	if m.IndexOfAddress("BBCCC") != 2 {
		t.Fatal("下标获取错误")
	}
	if m.ObjectAt(1).Address != "BBBB3" {
		t.Fatal("获取元素错误")
	}
	if m.Remove("BBBB3").Value != "124" {
		t.Fatal("删除元素错误")
	}
	if m.Count() != 2 {
		t.Fatal("元素个数错误")
	}
	println(m.DebugInfo())
}
