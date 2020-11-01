
package test1pkg

import (
"crypto/md5"
"fmt"
)

func Md5(text string) string {
	fmt.Println("v0.0.2")
	fmt.Println("v0.0.3")
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))

}
