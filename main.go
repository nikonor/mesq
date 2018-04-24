package mesq

import (
	"fmt"
	"sync"
	"time"
)

const (
	SUBSCRIBE       = 200040001
	UNSUBSCRIBE int = 200040002
)

var (
	// период ФП
	BGPeriod time.Duration
	// канал для завершения ФП
	ExitChan chan interface{}
	// Хранилище пакетов
	EventHeap TEventHeap
	// Хранилище событий
	MesHeap TMesHeap
)

type TEventHeap struct {
	sync.Mutex
	Heap map[string]EventType
}

type TMesHeap struct {
	sync.Mutex
	Heap map[string]TMessage
}

// TMessage - пакет событий
type TMessage struct {
	TMark    time.Time `json:"-"`
	Datetime string    `json:"datetime"`
	ID       string    `json:"id"`
	SystemID string    `json:"systemId"`
	Token    string    `json:"token"`
}

func init() {
	println("Call init")
	defer println("Finish init")
	// иницилизиуем хранилище для пакетов сообытий
	MesHeap = TMesHeap{}
	MesHeap.Heap = make(map[string]TMessage)
	// иницилизиуем хранилище для сообытий
	EventHeap = TEventHeap{}
	EventHeap.Heap = make(map[string]EventType)
}

// QInit - инициализация ФП
func QInit(timeOut time.Duration) {
	println("Call QInit")
	defer println("Finish QInit")

	BGPeriod = timeOut
	ExitChan = make(chan interface{})

	// стартует ФП отправки сообщений
	ticker := time.NewTicker(time.Second)
	go func(exitChan chan interface{}) {
		for {
			select {
			case <-ticker.C:
				println("Got ticker")
			case <-exitChan:
				println("Got Quit")
				ticker.Stop()
				return
			}
		}
	}(ExitChan)

}

func Quit() {
	close(ExitChan)
}

//
// func sendMes() {
// 	for k, v := range Heap.Dict {
// 		fmt.Printf("\t%s => %s\n", k, v.T.Format("15:04:05-0700"))
// 	}
//
// }
//
// func showQ() {
// 	fmt.Printf("Cur state on Heap\n")
// 	for k, v := range Heap.Dict {
//
// 		fmt.Printf("\t%s => %s\n", k, v.T.Format("15:04:05-0700"))
// 	}
// 	println()
// }
//
// AddMessageToQ - TODO: функция записи объекта в хранилище
func AddMessageToQ(body []byte) error {
	var (
		in  IncomMessageType
		err error
	)
	// разбираем запрос через easyjson
	if err = in.UnmarshalJSON(body); err != nil {
		return err
	}

	// добавляем его в хранилище пакетов
	mes := TMessage{
		TMark:    time.Now(),
		Datetime: in.Datetime,
		ID:       in.ID,
		SystemID: in.SystemID,
		Token:    in.Token,
	}
	MesHeap.Heap[in.ID] = mes

	// бежим по событиям и добавляем их в хранилище

	for _, event := range in.Events.Event {
		//    // создаем ключ и антиключ для хранения события
		directKey, backKey := MakeKey(
			event.TypeID,
			event.Filters.Filter[0].Persons.Person[0].SSOID,
			event.StreamID,
		)
		//    // проверяем антиключ
		EventHeap.Lock()
		if _, ok := EventHeap.Heap[backKey]; ok {
			// удаляем событие
			delete(EventHeap.Heap, backKey)
		} else {
			// добавляем событие
			event.MesID = in.ID
			EventHeap.Heap[directKey] = event
		}
		EventHeap.Unlock()
	}
	// fmt.Printf("Получили:\n\ttype=%d\n\tSSOID=%s\n\tStreamID=%d\n\n",
	// 	in.Events.Event[0].TypeID,
	// 	in.Events.Event[0].Filters.Filter[0].Persons.Person[0].SSOID,
	// 	in.Events.Event[0].StreamID,
	// )

	// directKey, backKey := MakeKey(in., in.CitizenID, in.Channel)
	// Heap.Lock()
	// if _, ok := Heap.Dict[backKey]; ok {
	// 	// нашли противоположную комаду - удаляем
	// 	delete(Heap.Dict, backKey)
	// } else {
	// 	// добавляем запись в хранилище
	// 	Heap.Dict[directKey] = QMesType{T: time.Now(), Body: body}
	// }
	// Heap.Unlock()

	return nil
}

func sORu(typeID int) (rune, rune) {
	switch typeID {
	case SUBSCRIBE:
		return 's', 'u'
	case UNSUBSCRIBE:
		return 'u', 's'
	default:
		return '-', '-'
	}
}

//
// MakeKey - создаем ключ из сообщения
func MakeKey(cmd int, ssoID string, channel int) (string, string) {
	r, rr := sORu(cmd)
	return fmt.Sprintf("%c#%s#%d", r, ssoID, channel), fmt.Sprintf("%c#%s#%d", rr, ssoID, channel)
}
