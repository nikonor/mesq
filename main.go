package mesq

import (
	"fmt"
	"sync"
	"time"
	"unicode"
)

var (
	Heap THeap
)

type THeap struct {
	Dict map[string]QMesType
	sync.Mutex
}

type QMesType struct {
	T    time.Time
	Body []byte
}

type INMes struct {
	Command   string `json:"command"`
	CitizenID int64  `json:"citizen_id"`
	Channel   int64  `json:"channel"`
}

func init() {
	println("Call init")
	defer println("Finish init")
	Heap = THeap{}
	Heap.Dict = make(map[string]QMesType)
}

func QInit() {
	println("Call QInit")
	defer println("Finish QInit")
	// TODO: создаем хранилище для объектов
	// TODO: стартует ФП отправки сообщений
}

func showQ() {
	fmt.Printf("Cur state on Heap\n")
	for k, v := range Heap.Dict {
		fmt.Printf("\t%s => %s\n", k, v.T.Format("15:04:05-0700"))
	}
	println()
}

// AddMessageToQ - TODO: функция записи объекта в хранилище
func AddMessageToQ(body []byte) error {
	var (
		in  INMes
		err error
	)
	// разбираем запрос через easyjson
	if err = in.UnmarshalJSON(body); err != nil {
		return err
	}

	directKey, backKey := MakeKey(in.Command, in.CitizenID, in.Channel)
	Heap.Lock()
	if _, ok := Heap.Dict[backKey]; ok {
		// нашли противоположную комаду - удаляем
		delete(Heap.Dict, backKey)
	} else {
		// добавляем запись в хранилище
		Heap.Dict[directKey] = QMesType{T: time.Now(), Body: body}
	}
	Heap.Unlock()

	return nil
}

// MakeKey - создаем ключ из сообщения
func MakeKey(cmd string, citizenID, channel int64) (string, string) {
	r := unicode.ToLower(rune(cmd[0]))
	rr := rune('s')
	if r == 's' {
		rr = 'u'
	}
	return fmt.Sprintf("%c-%d-%d", r, citizenID, channel), fmt.Sprintf("%c-%d-%d", rr, citizenID, channel)
}
