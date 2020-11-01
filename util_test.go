package test1pkg

import "testing"

func TestMd5(t *testing.T) {
	hashs := map[string]string{
		"a": "0cc175b9c0f1b6a831c399e269772661",
	}
	for k, v := range hashs {
		if Md5(k) != v {
			//fmt.Println(Md5(k))
			t.Errorf("%s md5 not is %s", k, v)
		}
	}
}
