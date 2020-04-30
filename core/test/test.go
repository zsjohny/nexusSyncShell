package main

import "nexusSync/core"

func main() {

	//fmt.Println(strings.HasPrefix("11123a", "1113"))
	config := core.Config{"POST",
		"xiao:xiao",
		"https://",
		2,
		"/lsj_test11",
		"E:\\learn\\15.左耳听风-陈皓", "WEB", "WEB", "job7Yvet2wB6HkaG"}

	servcie := new(core.NexusSyncService)
	servcie.StartUpload(&config)
	/*config = core.Config{"GET",
	"xiao:xiao",
	"https://",
	2,
	"/lsj_test11",
	"c:\\test",
	"xiao", "xiao", "xiao"}*/

	//servcie.StartDownload(&config)
	//url := "https://"
	//fmt.Println(path.Base(url))
	/*jsonStr:=`{
	  "items": [
	    {
	      "id": "eGlhbzoyOTVkMzU0YjcxM2EwOTI0ODY5Y2UyNWNkNzlkODM2ZA",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/lichuanxiao",
	      "name": "lichuanxiao/defaults.js",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/lichuanxiao/defaults.js",
	          "path": "lichuanxiao/defaults.js",
	          "id": "eGlhbzo5NjE4ODg5ZjVmODA2Mjg1YWJiNjY5NDAzNGU4YTExNQ",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "ab1943c4519eba0e3d1ffdc29f2e2f92f42ab166",
	            "sha512": "8d4a6c1db7f6d74a825d300235f7b05957081a59d6050cdd3c4c447d37f04692ff7bd38b5145031c6453cf16f14d87cfbfee955486947c629e9c2ea11e8ae68c",
	            "sha256": "8725f0b323ff61d6a0756be9bf79946f7064e4f369996dc1fa1ff2e57bdca086",
	            "md5": "7d119e28d7cb5778fa9e1300264c709d"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzpjNjkzMjA4YzRjZGY4NGI0Yzk4YTg1ZWM5ODVlMjZhOQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/chenjingxiong",
	      "name": "chenjingxiong/.jpg",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/chenjingxiong/.jpg",
	          "path": "chenjingxiong/.jpg",
	          "id": "eGlhbzo0ZDU3ZjcxYzFkNzk4ZDJkMTM0MDVkNjQ2MTE4ZWE0NQ",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "49ce21d3bb855afd3cdc6cc47100efd3a3fff7e0",
	            "sha512": "5c27023ac1474cf69096efbdb9ef44f277dfa71a91feddd5d928028b29d85c6df6a838dd1e8b053907c16243e64b225e87a495a14691c2a9ddd0c97e9686cad7",
	            "sha256": "469516720f81efd6b2ae8cb8cde14031e2bf9c8b1173aa32d9c187f14b5d6c63",
	            "md5": "46d99ee2610f91315978eb7dcfaa2e62"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzpkNmNkNTM2OTlmYTRiMDVlYjIzNzI3MWYzZDgyZTg2MQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/a_test",
	      "name": "a_test/mls_lca.log",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/a_test/mls_lca.log",
	          "path": "a_test/mls_lca.log",
	          "id": "eGlhbzoyMjVhZjY3YTc5NjJmNjU3ZDg4YzI5NzlkMTk4NWNlMw",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "eac008464fa833bad521b2234bc7ab510669c244",
	            "sha512": "cbbb17e6f06752a625feeba079c28f5779d41a31cee41c1910f26128b6bc1a58446b6a54000b70da833f813e22a0a41b817a48b1004cc08166a4245b47349f83",
	            "sha256": "4d2b8f91e9c73354376acc0a724997723c492b29a712dc5e01b4b62c7d07d7c2",
	            "md5": "31c8e497a70e7377c99fff096c121f9a"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzozNmE2NDBhMWIzMTBmZWQ2N2ExMDg4MGNiZmViNmJhNg",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/a_test",
	      "name": "a_test/a.jpg",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/a_test/a.jpg",
	          "path": "a_test/a.jpg",
	          "id": "eGlhbzo0Y2ZmNGZkMTQ0YWMzMWY2Mzg2OGY1ZDliZWY2MGViOQ",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "49ce21d3bb855afd3cdc6cc47100efd3a3fff7e0",
	            "sha512": "5c27023ac1474cf69096efbdb9ef44f277dfa71a91feddd5d928028b29d85c6df6a838dd1e8b053907c16243e64b225e87a495a14691c2a9ddd0c97e9686cad7",
	            "sha256": "469516720f81efd6b2ae8cb8cde14031e2bf9c8b1173aa32d9c187f14b5d6c63",
	            "md5": "46d99ee2610f91315978eb7dcfaa2e62"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzo3NjIyYzA5YTA5NTFhYjcxZDFiMTI5NTc0NGM0NjIwZQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/a_test",
	      "name": "a_test/.jpg",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/a_test/.jpg",
	          "path": "a_test/.jpg",
	          "id": "eGlhbzoyMmEyMDRlZTJhYjJiNmI4MzRjYzRjZWVkNTY3M2YxMg",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "49ce21d3bb855afd3cdc6cc47100efd3a3fff7e0",
	            "sha512": "5c27023ac1474cf69096efbdb9ef44f277dfa71a91feddd5d928028b29d85c6df6a838dd1e8b053907c16243e64b225e87a495a14691c2a9ddd0c97e9686cad7",
	            "sha256": "469516720f81efd6b2ae8cb8cde14031e2bf9c8b1173aa32d9c187f14b5d6c63",
	            "md5": "46d99ee2610f91315978eb7dcfaa2e62"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzozYzdhNDIxNzA3ZjIwYTc0MTRmNTU2MGNlMGIwZjE3Yw",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/a_test",
	      "name": "a_test/tony_test_scp.jpg",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/a_test/tony_test_scp.jpg",
	          "path": "a_test/tony_test_scp.jpg",
	          "id": "eGlhbzo0ZDU3ZjcxYzFkNzk4ZDJkMGJkZGIxMTZkZDBkZDk3MA",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "49ce21d3bb855afd3cdc6cc47100efd3a3fff7e0",
	            "sha512": "5c27023ac1474cf69096efbdb9ef44f277dfa71a91feddd5d928028b29d85c6df6a838dd1e8b053907c16243e64b225e87a495a14691c2a9ddd0c97e9686cad7",
	            "sha256": "469516720f81efd6b2ae8cb8cde14031e2bf9c8b1173aa32d9c187f14b5d6c63",
	            "md5": "46d99ee2610f91315978eb7dcfaa2e62"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzo1ZjQ5MWI2NmRmYzM0NjllMjdlZWViNzUzMjI5NDFkZQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/a_test",
	      "name": "a_test/record.txt",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/a_test/record.txt",
	          "path": "a_test/record.txt",
	          "id": "eGlhbzpmYjBhM2JiMmI0NGFhMzUwN2JhY2Q0YjY0MGQ1Mjc2Mg",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "f8b8a399a1186bdeea92baf12f13a1de71eddb04",
	            "sha512": "ad0b0d3da6a9179360d73a450802810965a5ccaeaae3e5b61c650efa8877b03174c53927df20b62e8fb591ce5437a0ac2230158891273d552ee957f4b371d69b",
	            "sha256": "6a00c9f39116dbb5878696dbd1f738444c6e9dd239ce9cf8868a03cab5e21ab7",
	            "md5": "cd4b3ea49b2fdc71a80816e20cff81ab"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzphODc2OTBiNTliNGQ0Y2Q3MTM5YjhlYWMwODBjMTY0OA",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/lsj_test",
	      "name": "lsj_test/mls_lca.log",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/lsj_test/mls_lca.log",
	          "path": "lsj_test/mls_lca.log",
	          "id": "eGlhbzo5NjE4ODg5ZjVmODA2Mjg1YjVkNDNlM2IzYjdhYzgzOQ",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "eac008464fa833bad521b2234bc7ab510669c244",
	            "sha512": "cbbb17e6f06752a625feeba079c28f5779d41a31cee41c1910f26128b6bc1a58446b6a54000b70da833f813e22a0a41b817a48b1004cc08166a4245b47349f83",
	            "sha256": "4d2b8f91e9c73354376acc0a724997723c492b29a712dc5e01b4b62c7d07d7c2",
	            "md5": "31c8e497a70e7377c99fff096c121f9a"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzowYjY1MzJhYTY0Yzc2YjhjMmNmOWRlMjQxZWJlZWQ0OQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/lsj_test",
	      "name": "lsj_test/1.log",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/lsj_test/1.log",
	          "path": "lsj_test/1.log",
	          "id": "eGlhbzozN2JhZGU2NDE1YjQzZDFiMmE5NWE5NGMzY2VkZmVlOQ",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "463469f3129bee419928e406ebfdcbfdda5fb54e",
	            "sha512": "0a838b27d7ad6c8e9b45e424ae327232d4989bd58d3a34b591ed871a8e30467434b066fa9e4350a6fb2b2fe1616b20c4c2a6ffe1ecfed20653025030ea9ca1ca",
	            "sha256": "acf426d021e5cfdabc6f6e61e98749940f86caafea3ddf20910821b09e2f0040",
	            "md5": "992d90fc29feefd8870a5a661b9f63d9"
	          }
	        }
	      ]
	    },
	    {
	      "id": "eGlhbzowZTgyNzc5MDcxODY1MGI1OWZhYmVjNzVjYjQxZmFmNQ",
	      "repository": "xiao",
	      "format": "raw",
	      "group": "/",
	      "name": ".pdf",
	      "version": null,
	      "assets": [
	        {
	          "downloadUrl": "https://nexus..com/repository/xiao/.pdf",
	          "path": ".pdf",
	          "id": "eGlhbzpmYjM1MTEwMDg4ZDc4ZGQwYzY5YWM5ODU2NWJjYjhmZA",
	          "repository": "xiao",
	          "format": "raw",
	          "checksum": {
	            "sha1": "436e25f2c66fcb8318f097b60a9abc5ec0e5b416",
	            "sha512": "61c9642aced874124ca465f700fad09f204652a13dd48b1ee1ea84d811f2d84435d31e27a0e8df30820d6b5add8bc4d503f5bcc75f56ce997a1548a08baa84e9",
	            "sha256": "f6f234a66846b20833c6c75bb7652d1bb91e2bf484c5d75b5cccc05791709a45",
	            "md5": "6c3b04a5d20db077055abfefbfea7bd8"
	          }
	        }
	      ]
	    }
	  ],
	  "continuationToken": "0e827790718650b59fabec75cb41faf5"
	}`

		var userJSON core.Body

		if err:=json.Unmarshal([]byte(jsonStr),&userJSON);err==nil{

			fmt.Println(userJSON)   //打印结果：{Tom 123456 [Li Fei]}

		}*/
	//path := "c:\\test\\"
	//
	//fmt.Println(util.GetAllFiles(path))
	//
	/*config := core.Config{"POST",
	"xiao:xiao",
	"https://nexus..com/service/rest/v1/components?repository=xiao",
	2,
	"/lsj_test",
	path}

	servcie := new(core.NexusSyncService)
	servcie.StartUpload(&config)*/

}
