package redislibs

import (
	"fmt"
	"testing"
)

func Test_ListNodes(t *testing.T) {

	// master_addr, err := GetMasterAddrByName("127.0.0.1", "5000", "mymaster")
	// if err == nil {
	// 	t.Log("测试通过了, 地址:", master_addr) //记录一些你期望记录的信息

	// }
	// SlaveOf("127.0.0.1", "6003", "127.0.0.1", "6001")

	// SlavesBySentinel("127.0.0.1", "5000", "mymaster")

	// if slaves, err := Slaves(master_addr.Host, master_addr.Port); err == nil {
	// 	t.Log("测试通过了, slave地址:", slaves)
	// }

	// sentinel_addrs := []*Address{BuildAddress("127.0.0.1", "5001"), BuildAddress("127.0.0.1", "5002"), BuildAddress("127.0.0.1", "5000")}
	// if err := RemoveSlaveFromSentinel("127.0.0.1", "6002", "master_service", sentinel_addrs...); err == nil {
	// 	t.Log("remove slave from sentinel pass")
	// }
	// if err := RemoveSlaveFromSentinel("127.0.0.1", "6002", "master_service", sentinel_addrs...); err == nil {
	// 	t.Log("remove slave from sentinel pass")
	// }
	res, _ := GetMasterAddrByName("127.0.0.1", "5000", "master_service")
	fmt.Println(res)
	// r, _ = GetMasterAddrByName("127.0.0.1", "5000", "master_service2")
	// fmt.Println(r)

	// GetSlavesInSentinel("127.0.0.1", "5000", "master_service")
	// GetSlavesInSentinel("127.0.0.1", "5000", "master_service2")

	RoleStatus("127.0.0.1", "6001")
	if role, status, err := RoleStatus("127.0.0.1", "6002"); err == nil {
		fmt.Println(role, status)
	}

	r, err := RedisNodeInfo("127.0.0.1", "6002")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println(r)

	masters, err := FetchMastersInSentinel("127.0.0.1", "5000")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, m := range masters {
		fmt.Println(m)
	}
	// addr, _ := NewAddress(":5000")
	// fmt.Println(addr)

}