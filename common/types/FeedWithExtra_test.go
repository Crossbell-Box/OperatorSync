package types

import (
	"encoding/json"
	"testing"
)

func TestFeedWithExtra(t *testing.T) {

	jsonFeed := `{
    "version": "https://jsonfeed.org/version/1.1",
    "title": "Twitter @柠檬🍋",
    "home_page_url": "https://twitter.com/lc499",
    "description": "是柠檬呀（测试账号）\nlc499@crossbell - Made with love by RSSHub(https://github.com/DIYgod/RSSHub)",
    "icon": "https://pbs.twimg.com/profile_images/1534462048540602369/Va32doHM.jpg",
    "language": "zh-cn",
    "items": [
        {
            "id": "https://twitter.com/lc499/status/1612993512630542339",
            "url": "https://twitter.com/lc499/status/1612993512630542339",
            "title": "引用测试 2301111003",
            "content_html": "引用测试 2301111003<div class=\"rsshub-quote\"><br><br>柠檬🍋: 发送带有媒体资源文件的推文测试（请注意媒体顺序） 2211101058<br><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7Bt2VQAESCTX?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BuDUAAE1Nmc?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BuTUoAIg14Y?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BumVEAIkJfl?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"></div>",
            "date_published": "2023-01-11T02:03:16.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ],
            "_extra": {
                "links": [
                    {
                        "url": "https://twitter.com/lc499/status/1590539285400539136",
                        "content_html": "<div class=\"rsshub-quote\"><br><br>柠檬🍋:&ensp;发送带有媒体资源文件的推文测试（请注意媒体顺序） 2211101058<br><br><img  style=\"\"  src=\"https://pbs.twimg.com/media/FhK7Bt2VQAESCTX?format=png&name=orig\"><br><img  style=\"\"  src=\"https://pbs.twimg.com/media/FhK7BuDUAAE1Nmc?format=png&name=orig\"><br><img  style=\"\"  src=\"https://pbs.twimg.com/media/FhK7BuTUoAIg14Y?format=png&name=orig\"><br><img  style=\"\"  src=\"https://pbs.twimg.com/media/FhK7BumVEAIkJfl?format=png&name=orig\"></div>",
                        "type": "quote"
                    }
                ]
            }
        },
        {
            "id": "https://twitter.com/lc499/status/1605543399204933633",
            "url": "https://twitter.com/lc499/status/1605543399204933633",
            "title": "Requesting $CSB funds from the Faucet on the #Crossbell blockchain. Address: 0xcD501C2b25CfE41e0d8EB96DEC602Eb2Bc59c1Ed",
            "content_html": "Requesting $CSB funds from the Faucet on the #Crossbell blockchain. Address: 0xcD501C2b25CfE41e0d8EB96DEC602Eb2Bc59c1Ed <a href=\"https://faucet.crossbell.io/\" target=\"_blank\" rel=\"noopener noreferrer\">https://faucet.crossbell.io/</a>",
            "date_published": "2022-12-21T12:39:10.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1590539285400539136",
            "url": "https://twitter.com/lc499/status/1590539285400539136",
            "title": "发送带有媒体资源文件的推文测试（请注意媒体顺序） 2211101058",
            "content_html": "发送带有媒体资源文件的推文测试（请注意媒体顺序） 2211101058<br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7Bt2VQAESCTX?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BuDUAAE1Nmc?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BuTUoAIg14Y?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\"><br><img style=\"\" src=\"https://pbs.twimg.com/media/FhK7BumVEAIkJfl?format=png&amp;name=orig\" referrerpolicy=\"no-referrer\">",
            "date_published": "2022-11-10T02:58:11.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1590216889690906624",
            "url": "https://twitter.com/lc499/status/1590216889690906624",
            "title": "消息发送测试 2211091337",
            "content_html": "消息发送测试 2211091337",
            "date_published": "2022-11-09T05:37:05.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1590215760009625600",
            "url": "https://twitter.com/lc499/status/1590215760009625600",
            "title": "消息发送测试 2211091332",
            "content_html": "消息发送测试 2211091332",
            "date_published": "2022-11-09T05:32:36.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1589988813145014272",
            "url": "https://twitter.com/lc499/status/1589988813145014272",
            "title": "消息发送测试 2211082230",
            "content_html": "消息发送测试 2211082230",
            "date_published": "2022-11-08T14:30:48.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1588519262473576448",
            "url": "https://twitter.com/lc499/status/1588519262473576448",
            "title": "消息发送测试 2211042111",
            "content_html": "消息发送测试 2211042111",
            "date_published": "2022-11-04T13:11:20.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1588519165429968896",
            "url": "https://twitter.com/lc499/status/1588519165429968896",
            "title": "消息发送测试 2211042110",
            "content_html": "消息发送测试 2211042110",
            "date_published": "2022-11-04T13:10:56.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1588095605116530688",
            "url": "https://twitter.com/lc499/status/1588095605116530688",
            "title": "消息发送测试 2211031707",
            "content_html": "消息发送测试 2211031707",
            "date_published": "2022-11-03T09:07:52.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1586995075883548673",
            "url": "https://twitter.com/lc499/status/1586995075883548673",
            "title": "消息发送测试 2210311614",
            "content_html": "消息发送测试 2210311614",
            "date_published": "2022-10-31T08:14:45.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1586949395085348864",
            "url": "https://twitter.com/lc499/status/1586949395085348864",
            "title": "消息发送测试 2210311313",
            "content_html": "消息发送测试 2210311313",
            "date_published": "2022-10-31T05:13:14.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1579720927117717505",
            "url": "https://twitter.com/lc499/status/1579720927117717505",
            "title": "图片发送测试 2210111430",
            "content_html": "图片发送测试 2210111430<br><img style=\"\" src=\"https://pbs.twimg.com/media/FexLzY-UUAAUFs_?format=jpg&amp;name=orig\" referrerpolicy=\"no-referrer\">",
            "date_published": "2022-10-11T06:29:53.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1579707110295998464",
            "url": "https://twitter.com/lc499/status/1579707110295998464",
            "title": "消息发送测试 2210111335",
            "content_html": "消息发送测试 2210111335",
            "date_published": "2022-10-11T05:34:59.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1579700925060771840",
            "url": "https://twitter.com/lc499/status/1579700925060771840",
            "title": "消息发送测试 2210111310",
            "content_html": "消息发送测试 2210111310",
            "date_published": "2022-10-11T05:10:24.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1579424505734250497",
            "url": "https://twitter.com/lc499/status/1579424505734250497",
            "title": "Brave 浏览器 消息发送测试 2210101852",
            "content_html": "Brave 浏览器 消息发送测试 2210101852",
            "date_published": "2022-10-10T10:52:01.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1575319220270661632",
            "url": "https://twitter.com/lc499/status/1575319220270661632",
            "title": "转发测试 2209291059",
            "content_html": "转发测试 2209291059<div class=\"rsshub-quote\"><br><br>柠檬🍋: 回复测试 2209291057<br></div>",
            "date_published": "2022-09-29T02:59:04.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ],
            "_extra": {
                "links": [
                    {
                        "url": "https://twitter.com/lc499/status/1575318771329224704",
                        "content_html": "<div class=\"rsshub-quote\"><br><br>柠檬🍋:&ensp;回复测试 2209291057<br></div>",
                        "type": "quote"
                    }
                ]
            }
        },
        {
            "id": "https://twitter.com/lc499/status/1575318771329224704",
            "url": "https://twitter.com/lc499/status/1575318771329224704",
            "title": "Re 回复测试 2209291057",
            "content_html": "Re 回复测试 2209291057",
            "date_published": "2022-09-29T02:57:17.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ],
            "_extra": {
                "links": [
                    {
                        "url": "https://twitter.com/lc499/status/1575025392032698368",
                        "type": "reply"
                    }
                ]
            }
        },
        {
            "id": "https://twitter.com/lc499/status/1575025392032698368",
            "url": "https://twitter.com/lc499/status/1575025392032698368",
            "title": "图片发送测试 2209281531",
            "content_html": "图片发送测试 2209281531<br><img style=\"\" src=\"https://pbs.twimg.com/media/FdudPD7UYAAMvte?format=jpg&amp;name=orig\" referrerpolicy=\"no-referrer\">",
            "date_published": "2022-09-28T07:31:30.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1574626648682426368",
            "url": "https://twitter.com/lc499/status/1574626648682426368",
            "title": "内容发送测试 2209271307",
            "content_html": "内容发送测试 2209271307",
            "date_published": "2022-09-27T05:07:02.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ]
        },
        {
            "id": "https://twitter.com/lc499/status/1574367617191550976",
            "url": "https://twitter.com/lc499/status/1574367617191550976",
            "title": "Re 连续创建内容测试 2209261957 - 3",
            "content_html": "Re 连续创建内容测试 2209261957 - 3",
            "date_published": "2022-09-26T11:57:44.000Z",
            "authors": [
                {
                    "name": "柠檬🍋"
                }
            ],
            "_extra": {
                "links": [
                    {
                        "url": "https://twitter.com/lc499/status/1574367614125473792",
                        "type": "reply"
                    }
                ]
            }
        }
    ]
}`

	var feed FeedWithExtra

	err := json.Unmarshal([]byte(jsonFeed), &feed)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(feed)

}
