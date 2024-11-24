package main

import (
	"errors"
	"log/slog"
	"os"
	"sync"
	"time"
)

type TemperatureSensor struct {
	kitchenTemp float32
	bedroomTemp float32
	cabinetTemp float32
	mu          sync.RWMutex
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	sensor := TemperatureSensor{kitchenTemp: 25.0, bedroomTemp: 24.1, cabinetTemp: 21.2}
	go getTemp(&sensor, "kitchen")
	go getTemp(&sensor, "bedroom")
	go getTemp(&sensor, "cabinet")

	go setTemp(&sensor, "kichen", 26)
	go setTemp(&sensor, "bedroom", 27)
	go setTemp(&sensor, "cabinet", 28)

	time.Sleep(2 * time.Second)

}

func getTemp(t *TemperatureSensor, room string) (float32, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	switch room {
	case "kitchen":
		slog.Info("Getting kitchen temperature")
		return t.kitchenTemp, nil
	case "bedroom":
		slog.Info("Getting bedroom temperature")
		return t.bedroomTemp, nil
	case "cabinet":
		slog.Info("Getting cabinet temperature")
		return t.cabinetTemp, nil
	default:
		return .0, errors.New("room not found")
	}
}

func setTemp(t *TemperatureSensor, room string, temp float32) (float32, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	switch room {
	case "kitchen":
		slog.Info("Setting kitchen temperature", "temperature", temp)
		t.kitchenTemp = temp
	case "bedroom":
		slog.Info("Setting bedroom temperature", "temperature", temp)
		t.bedroomTemp = temp
	case "cabinet":
		slog.Info("Setting cabinet temperature", "temperature", temp)
		t.cabinetTemp = temp
	default:
		return .0, errors.New("room not found")
	}
	return temp, nil
}

// 3. Мониторинг температуры (Чтение и запись)
// Задача: Написать программу для мониторинга температуры в разных местах. Несколько горутин могут параллельно читать значения температуры,
// но только одна горутина может обновить температуру в каждом месте.

// Описание:

// Создайте структуру TemperatureSensor, которая хранит температуру для разных датчиков (например, для разных комнат или территорий).
// Используйте мьютекс на чтение и запись для синхронизации доступа к данным температуры.
// Несколько горутин могут одновременно читать температуру.
// Только одна горутина может обновить температуру в одном датчике.
// Напишите несколько горутин для чтения температуры.
// Напишите несколько горутин для обновления значений температуры (например, изменение температуры из-за изменения условий).
// Программа должна корректно работать с параллельными операциями.
// Подсказка: Используйте sync.RWMutex для обеспечения безопасности данных и методы RLock(), RUnlock(), Lock(), Unlock() для чтения и записи.
