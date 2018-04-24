package mesq

import (
    "time"
    "sync"
    "fmt"
    "unicode"
)

var (
    Heap map[string]THeapElement
)

type THeapElement struct {
    Mes QMesType
    sync.Mutex
}

type QMesType struct {
    T time.Time
    Body []byte
}

type INMes struct {
    Command string `json:"command"`
    CitizenID int64 `json:"citizen_id"`
    Channels []int64 `json:"channels"`
}

func init () {
    println("Call init")
    defer println("Finish init")
    Heap = make(map[string]THeapElement)
}

func QInit() {
    println("Call QInit")
    defer println("Finish QInit")
    // TODO: создаем хранилище для объектов
    // TODO: стартует ФП отправки сообщений
}

// AddMessageToQ - TODO: функция записи объекта в хранилище
func AddMessageToQ (body []byte) error {
    var (
        in INMes
        err error
    )
    // разбираем запрос через easyjson
    if err = in.UnmarshalJSON(body); err != nil {
        return err
    }
    
    return nil
}

// MakeKey - создаем ключ из сообщения
func MakeKey (cmd string, citizenID, channel int64) string {
    return fmt.Sprintf("%c-%d-%d", unicode.ToLower(rune(cmd[0])), citizenID, channel)
}