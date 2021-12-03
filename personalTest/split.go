package main

import (
	"fmt"
	"strings"
)

const msg = `Received: from axis-accc8e94fc46 (unknown [192.168.1.150])
by Amber (MyServerApp) with SMTP
for <amberl@chekt.com>; Thu,  2 Dec 2021 17:01:25 +0900 (KST)
To: =?utf-8?B?YW1iZXJsQGNoZWt0LmNvbQ==?= <amberl@chekt.com>
From: =?utf-8?B?YW1iZXJsQGNoZWt0LmNvbQ==?= <amberl@chekt.com>
Subject: =?utf-8?B?bW90aW9uIGRldGVjdGVk?=
Date: Thu, 02 Dec 2021 17:02:10 +0000MIME-version: 1.0
Content-Type: multipart/mixed; boundary=curlsink-boundary

--curlsink-boundary
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: BASE64

YW1iZXI=`

func main() {
	test := strings.Split(msg, "boundary=")
	// fmt.Println("test[0]", test[0])
	fmt.Println("test[1]", test[1])
	test1 := test[1]
	test2 := strings.Split(test1, "\n")
	fmt.Println("test2>>>", test2[0])
}

/*
test[0] Received: from axis-accc8e94fc46 (unknown [192.168.1.150])
by Amber (MyServerApp) with SMTP
for <amberl@chekt.com>; Thu,  2 Dec 2021 17:01:25 +0900 (KST)
To: =?utf-8?B?YW1iZXJsQGNoZWt0LmNvbQ==?= <amberl@chekt.com>
From: =?utf-8?B?YW1iZXJsQGNoZWt0LmNvbQ==?= <amberl@chekt.com>
Subject: =?utf-8?B?bW90aW9uIGRldGVjdGVk?=
Date: Thu, 02 Dec 2021 17:02:10 +0000MIME-version: 1.0
Content-Type: multipart/mixed;
*/

/*
test[1] =curlsink-boundary

--curlsink-boundary
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: BASE64
*/
