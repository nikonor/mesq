package mesq

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddMessageToQ(t *testing.T) {
	// QInit(3 * time.Second)

	cases := []struct {
		in    string
		cmd   string
		count int
	}{
		{`{
	"datetime": "2018-04-24T12:02:13+03:00",
	"events": {
		"event": [
			{
				"datetime": "2018-04-24T12:02:11+03:00",
				"description": "Произошла подписка",
				"filters": {
					"filter": [
						{
							"persons": {
								"person": [
									{
										"SSOID": "SSOID123"
									}
								]
							}
						}
					]
				},
				"id": "ID11111-55555-66666-4444ID",
				"message": {
					"parameters": {
						"parameter": [
							{
								"name": "streamId",
								"value": "4"
							},
							{
								"name": "channelFlags",
								"value": "2"
							},
							{
								"name": "opt_mpgu_vc",
								"value": "1231231211"
							},
							{
								"name": "msisdn",
								"value": "79876543210"
							},
							{
								"name": "email",
								"value": "mail@mail.ru"
							},
							{
								"name": "sendStartTimeEmail",
								"value": "00:00:00"
							},
							{
								"name": "sendStopTimeEmail",
								"value": "23:59:59"
							},
							{
								"name": "sendExcludeDaysEmail",
								"value": "0"
							},
							{
								"name": "dateTime",
								"value": "2018-04-24T12:02:11+03:001524560531"
							}
						]
					}
				},
				"streamId": 4,
				"typeId": 200040001
			}
		]
	},
	"id": "ID11111-777777-888888-4444ID",
	"systemId": "EMP",
	"token": "123"
}`, "s", 1},
		{`{
	"datetime": "2018-04-24T12:02:50+03:00",
	"events": {
		"event": [
			{
				"datetime": "2018-04-24T12:02:47+03:00",
				"description": "Произошла отписка",
				"filters": {
					"filter": [
						{
							"persons": {
								"person": [
									{
										"SSOID": "SSOID123"
									}
								]
							}
						}
					]
				},
				"id": "ID11111-9999999-000000-4444ID",
				"message": {
					"parameters": {
						"parameter": [
							{
								"name": "streamId",
								"value": "4"
							},
							{
								"name": "channelFlags",
								"value": "2"
							},
							{
								"name": "msisdn",
								"value": "79876543210"
							},
							{
								"name": "email",
								"value": "mail@mail.ru"
							},
							{
								"name": "opt_mpgu_vc",
								"value": "1231231211"
							},
							{
								"name": "dateTime",
								"value": "2018-04-24T12:02:47+03:001524560567"
							}
						]
					}
				},
				"streamId": 4,
				"typeId": 200040002
			}
		]
	},
	"id": "ID11111-2222-33333-4444ID",
	"systemId": "EMP",
	"token": "123"
}`, "u", 1},
		// {`{"command":"subscribe","citizen_id":222,"channel":1}`, "S", 2},
		// {`{"command":"subscribe","citizen_id":222,"channel":2}`, "S", 1},
		// {`{"command":"subscribe","citizen_id":222,"channel":3}`, "S", 2},
		// {`{"command":"unsubscribe","citizen_id":222,"channel":1}`, "S", 1},
	}

	for _, c := range cases {
		if err := AddMessageToQ([]byte(c.in)); err != nil {
			t.Error(err)
		}
		fmt.Printf("LEN Mes=%d\nLEN Events=%d\n", len(MesHeap.Heap), len(EventHeap.Heap))
		// if len(Heap.Dict) != c.count {
		// 	t.Error("ошибка подсчета")
		// }
	}

	// showQ()
}

func TestMakeKey(t *testing.T) {
	cases := []struct {
		cmd       int
		ssoID     string
		channel   int
		outDirect string
		outBack   string
	}{
		{SUBSCRIBE, "1", 1, "s#1#1", "u#1#1"},
		{SUBSCRIBE, "1", 1, "s#1#1", "u#1#1"},
		{UNSUBSCRIBE, "1", 1, "u#1#1", "s#1#1"},
		{UNSUBSCRIBE, "1", 1, "u#1#1", "s#1#1"},
	}

	for _, c := range cases {
		s1, s2 := MakeKey(c.cmd, c.ssoID, c.channel)
		if s1 != c.outDirect || s2 != c.outBack {
			t.Errorf("%s != %s || %s != %s\n", s1, c.outDirect, s2, c.outBack)
		}
	}
}

func TestQuit(t *testing.T) {
	println("Start TestQuit:", time.Now().Format("15:04:05"))
	QInit(3 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(4 * time.Second)
		Quit()
		wg.Done()
	}()
	wg.Wait()
	println("Finish TestQuit:", time.Now().Format("15:04:05"))
}
